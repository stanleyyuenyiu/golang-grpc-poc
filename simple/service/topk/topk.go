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
		log.Printf("[IncreaseCounter] Addess for %v: %p %v", key, s.mapping[key], 1)
		heap.Push(s.pq, s.mapping[key])
	}else{
		item := s.mapping[key]
		item.Priority +=1

		if item.Index < 0 {
			heap.Push(s.pq, item)
		} else {
			heap.Fix(s.pq, item.Index)
		}
		log.Printf("[IncreaseCounter] Addess for %v: %p %v  index: %v", key, item, item.Priority, item.Index)
	}
	if s.pq.Len() > s.k {
		item := heap.Pop(s.pq).(*pq.TopKItem)
		log.Printf("[IncreaseCounter] Size exceed, pop %s with priority %d", item.Value, item.Priority)
	}
}

func listTopk(q *pq.PriorityQueue, k int)  (msg map[string]uint32) {
	msg = make(map[string]uint32, 0)
	var qClone pq.PriorityQueue

	c := *q
	i := 0
	for i < q.Len() {
		//Item in PriorityQueue stored as address, hence we need to clone it, instead of reference it, as heap will modifiy the item property 
		item := *c[i]
		qClone = append(qClone , &pq.TopKItem{
			Value: item.Value,
			Priority: item.Priority,
		})
		i+=1
	}
	
	//Re heapify the cloned q
	heap.Init(&qClone)

	log.Printf("[listTopk] len for org %v, new %v", q.Len(), qClone.Len())

	for qClone.Len() > 0 && k > 0 {
		item := heap.Pop(&qClone).(*pq.TopKItem)
		msg[item.Value] = uint32(item.Priority)
		k -= 1
	}

	log.Printf("[listTopk] len for org %v, new %v", q.Len(), qClone.Len())
	
	return msg
}


func (s *TopkService) IncreaseCounter(ctx context.Context, in *pb.IncreaseCounterReq) (*pb.IncreaseCounterRes, error) {
	key := in.GetKey()
	log.Printf("[IncreaseCounter] Received Key: %v", key)

	increaseCounter(s, key)

	log.Printf("[IncreaseCounter] Key counter: %d", s.mapping[key].Priority)
	return &pb.IncreaseCounterRes{ IsSuccess:true, Message: "success"}, nil
}

func (s *TopkService) ListTopK(ctx context.Context, in *pb.ListTopKReq) (*pb.ListTopKRes, error) {
	k := int(in.GetK())
	log.Printf("[ListTopK] Received Key: %v", k)

	res :=  listTopk(s.pq, k)

	return &pb.ListTopKRes{ IsSuccess:true, Message: res}, nil
}

func (s *TopkService) StreamListTopK(in *pb.ListTopKReq, stream pb.Topk_StreamListTopKServer ) error {
	k := int(in.GetK())

	log.Printf("[StreamListTopK] Received Key: %v", k)

	msg :=  listTopk(s.pq, k)
	if err := stream.Send(&pb.ListTopKRes {
		IsSuccess:true, 
		Message: msg,
	}); err != nil {
		return err
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
		log.Printf("[IncreaseCounter] Key counter: %d", s.mapping[key].Priority)
		
		msg :=  listTopk(s.pq, s.k)
		if err := stream.Send(&pb.ListTopKRes {
			IsSuccess:true, 
			Message: msg,
		}); err != nil {
			return err
		}
	}
}


func Register(s grpc.ServiceRegistrar) {
	service := TopkService{}
	service.mapping = make(map[string]*pq.TopKItem)
	service.k = 5

	pq := make(pq.PriorityQueue, 0)

	//PriorityQueue needed to modify pq property, hence we use address 
	heap.Init(&pq)
	service.pq = &pq

	pb.RegisterTopkServer(s, &service);
}	