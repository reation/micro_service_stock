package logic

import (
	"context"

	"github.com/reation/micro_service_stock/protoc"
	"github.com/reation/micro_service_stock/stock_op/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SupplierReduceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSupplierReduceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SupplierReduceLogic {
	return &SupplierReduceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SupplierReduceLogic) SupplierReduce(in *protoc.SupplierReduceRequest) (*protoc.SupplierReduceResponse, error) {
	// todo: add your logic here and delete this line

	return &protoc.SupplierReduceResponse{}, nil
}
