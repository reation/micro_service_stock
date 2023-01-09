package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"micro_service_stock/protoc"
	"time"
)

const (
	StockOrderReduceAddress = "192.168.1.21:9090"
)

type User struct {
	id   int64
	name string
}

func main() {
	orderReduce()
}

func orderReduce() {

	conn, err := grpc.Dial(StockOrderReduceAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewOrderReduceStockOperationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var data = make([]*protoc.OrderGoodsID, 2)

	data[0] = &protoc.OrderGoodsID{
		GoodsId:  3,
		GoodsNum: 5,
	}
	data[1] = &protoc.OrderGoodsID{
		GoodsId:  2,
		GoodsNum: 3,
	}

	r, err := c.OrderReduceStock(ctx, &protoc.OrderReduceStockRequest{GoodsID: data, OrderID: 8729991, IsAll: 0})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	for _, v := range r.GetGoodsStates() {
		log.Println(v)
	}

}
