package topk_service;

import (
	"context"
	"log"
	"io"
	"google.golang.org/grpc"
	"container/heap"
	pb "simple/protos/topk"
	pq "simple/service/topk/pq"
) ;

type TopkService struct {
	pb.UnimplementedTopkServer
	mapping map[string]*pq.TopKItem
	pq *pq.PriorityQueue
	k int
}

func increaseCounter(s *TopkService, key string) {
	if s.mapping[key] == nil {
		s.mapping[key] = &pq.TopKItem{
			Value: key,
			Priority: 1,
		}
		log.Printf("[IncreaseCounter] Addess for %v: %p", key, s.mapping[key])
		heap.Push(s.pq, s.mapping[key])
	}else{
		s.mapping[key].Priority += 1
		log.Printf("[IncreaseCounter] Addess for %v: %p", key, s.mapping[key])
		heap.Fix(s.pq, s.mapping[key].Index)
	}
	if s.pq.Len() > s.k {
		item := heap.Pop(s.pq).(*pq.TopKItem)
		log.Printf("[IncreaseCounter] Size exceed, pop %s with priority %d", item.Value, item.Priority)
	}
}

func listTopK(s *TopkService) ( res map[string]uint32 ) {
	var heapItems []*pq.TopKItem

	for s.pq.Len() > 0 {
		item := heap.Pop(s.pq).(*pq.TopKItem)
		log.Printf("[ListTopK] Pop Addess for %v: %p", item.Value, item)

		res[item.Value] = uint32(item.Priority)
		heapItems = append(heapItems, item)
	}

	log.Printf("[ListTopK] After pop all pq: %v", s.pq.Len())

	for _, heapItem := range heapItems {
		log.Printf("[ListTopK] Push Addess for %v: %p", heapItem.Value, heapItem)
		heap.Push(s.pq, heapItem)
	}

	log.Printf("[ListTopK] Re import pq: %v", s.pq.Len())

	return res
}

func (s *TopkService) IncreaseCounter(ctx context.Context, in *pb.IncreaseCounterReq) (*pb.IncreaseCounterRes, error) {
	key := in.GetKey()
	log.Printf("[IncreaseCounter] Received Key: %v", key)

	increaseCounter(s, key)

	log.Printf("[IncreaseCounter] Key counter: %d", s.mapping[key].Priority)
	return &pb.IncreaseCounterRes{ IsSuccess:true, Message: "success"}, nil
}

func (s *TopkService) ListTopK(ctx context.Context, in *pb.ListTopKReq) (*pb.ListTopKRes, error) {
	k := in.GetK()
	log.Printf("[ListTopK] Received Key: %v", k)

	res := listTopK(s)

	return &pb.ListTopKRes{ IsSuccess:true, Message: res}, nil
}

func (s *TopkService) StreamListTopK(in *pb.ListTopKReq, stream pb.Topk_StreamListTopKServer ) error {
	k := in.GetK()
	log.Printf("[ListTopK] Received Key: %v", k)
	for _, item := range *s.pq {
		msg := pb.ListTopKRes {
			IsSuccess:true, 
			Message: map[string]uint32{item.Value:uint32(item.Priority)},
		}

		if err := stream.Send(&msg); err != nil {
			return err
		}
	}
	return nil
}

func (s *TopkService) StreamIncreaseCounter(stream pb.Topk_StreamIncreaseCounterServer ) error {

	for {
		in, err := stream.Recv()

		if err == io.EOF {
		  	return nil
		}
		if err != nil {
			return err
		}

		key := in.GetKey()

		log.Printf("[StreamIncreaseCounter] Received Key: %v", key)

		increaseCounter(s, key)
		
		for _, item := range *s.pq {

			msg := pb.ListTopKRes {
				IsSuccess:true, 
				Message: map[string]uint32{item.Value:uint32(item.Priority)},
			}
	
			if err := stream.Send(&msg); err != nil {
				return err
			}
		}
		
	}
}

func Register(s grpc.ServiceRegistrar) {
	
	service := TopkService{}
	service.mapping = make(map[string]*pq.TopKItem)
	service.k = 5

	pq := make(pq.PriorityQueue, 0)

	heap.Init(&pq)
	service.pq = &pq

	pb.RegisterTopkServer(s, &service);
}	