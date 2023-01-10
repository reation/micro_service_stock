// Code generated by goctl. DO NOT EDIT.
// Source: stock_op.proto

package stockop

import (
	"context"

	"micro_service_stock/protoc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	OrderReturnRequest       = protoc.OrderReturnRequest
	OrderReturnResponse      = protoc.OrderReturnResponse
	StockOpGoodInfo          = protoc.StockOpGoodInfo
	SupplierIncreaseRequest  = protoc.SupplierIncreaseRequest
	SupplierIncreaseResponse = protoc.SupplierIncreaseResponse
	SupplierReduceRequest    = protoc.SupplierReduceRequest
	SupplierReduceResponse   = protoc.SupplierReduceResponse

	StockOp interface {
		OrderReturn(ctx context.Context, in *OrderReturnRequest, opts ...grpc.CallOption) (*OrderReturnResponse, error)
		SupplierIncrease(ctx context.Context, in *SupplierIncreaseRequest, opts ...grpc.CallOption) (*SupplierIncreaseResponse, error)
		SupplierReduce(ctx context.Context, in *SupplierReduceRequest, opts ...grpc.CallOption) (*SupplierReduceResponse, error)
	}

	defaultStockOp struct {
		cli zrpc.Client
	}
)

func NewStockOp(cli zrpc.Client) StockOp {
	return &defaultStockOp{
		cli: cli,
	}
}

func (m *defaultStockOp) OrderReturn(ctx context.Context, in *OrderReturnRequest, opts ...grpc.CallOption) (*OrderReturnResponse, error) {
	client := protoc.NewStockOpClient(m.cli.Conn())
	return client.OrderReturn(ctx, in, opts...)
}

func (m *defaultStockOp) SupplierIncrease(ctx context.Context, in *SupplierIncreaseRequest, opts ...grpc.CallOption) (*SupplierIncreaseResponse, error) {
	client := protoc.NewStockOpClient(m.cli.Conn())
	return client.SupplierIncrease(ctx, in, opts...)
}

func (m *defaultStockOp) SupplierReduce(ctx context.Context, in *SupplierReduceRequest, opts ...grpc.CallOption) (*SupplierReduceResponse, error) {
	client := protoc.NewStockOpClient(m.cli.Conn())
	return client.SupplierReduce(ctx, in, opts...)
}