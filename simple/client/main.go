package main;

import (
	"context"
	"flag"
	"log"
	"time"
	"io"
	"strconv"
	"math/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "simple/protos/topk"
) ;

const (
	defaultKey = ""
	defaultFunc = "1"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	key = flag.String("key", defaultKey, "Key to increment topK")
	f = flag.String("f", defaultFunc, "func")
)


func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTopkClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Second)
	defer cancel()

	if *f == "1" {
		increase(ctx, c)
	} else if *f == "2" {
		get(ctx, c)
	} else if *f == "3"  {
		steamGet(ctx, c)
	} else if *f == "4"  {
		steamIncrease(ctx, c)
	} else {
		log.Print("No Such key")
	}
	
}

func increase(ctx context.Context, c pb.TopkClient) {
	if *key != "" {
		r, err := c.IncreaseCounter(ctx, &pb.IncreaseCounterReq{Key: *key })
		if err != nil {
			log.Fatalf("could not increaseCounter: %v", err)
		}
		log.Printf(" %s", r.GetMessage())
	} else {
		i := 0
		for i < 5 {
			r, err := c.IncreaseCounter(ctx, &pb.IncreaseCounterReq{Key: "key"+strconv.Itoa(i) })
			if err != nil {
				log.Fatalf("could not increaseCounter: %v", err)
			}
			i += 1
			log.Printf(" %s", r.GetMessage())
		}
	}
}

func get(ctx context.Context, c pb.TopkClient) {
	r, err := c.ListTopK(ctx, &pb.ListTopKReq{K:5})
	if err != nil {
		log.Fatalf("could not ListTopK: %v", err)
	}
	for k,v := range r.GetMessage() {
		log.Printf("%v %v", k, v)
	}
}

func steamGet(ctx context.Context, c pb.TopkClient) {
	r, err := c.StreamListTopK(ctx, &pb.ListTopKReq{K:5})
	if err != nil {
		log.Fatalf("could not StreamListTopK: %v", err)
	}
	for {
		res, err := r.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.StreamListTopK(_) = _, %v", c, err)
		}
		log.Println(res)
	}
}

func steamIncrease(ctx context.Context, c pb.TopkClient) {
	r, err := c.StreamIncreaseCounter(ctx)
	if err != nil {
		log.Fatalf("could not steamIncrease: %v", err)
	}

	// init memory less channel to avoid race condtition
	ch := make(chan struct{})
	// start goroutines for concurrent stream read
	go func() {
		// executed when program end
		defer close(ch)

		for {
			res, err := r.Recv()

			if err == io.EOF {
			  //close channel when read end
			  return
			}

			if err != nil {
			  log.Fatalf("Failed to recv steamIncrease : %v", err)
			}
			log.Printf("Recv: %s" , res)
		  }
	}()
	
	i := 0
	for i < 5 {
		key := "key" + strconv.Itoa(rand.Intn(50))

		req := pb.IncreaseCounterReq{Key: key}

		log.Printf("Send: %v" , key)

		if err := r.Send( &req ); err != nil {
			log.Fatalf("Failed to send steamIncrease: %v", err)
		}
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		i+=1
	}

	r.CloseSend()

	//assign channel = end of the channel
	<-ch
}
