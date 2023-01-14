package logic

import (
	"context"
	"github.com/reation/micro_service_stock/config"
	"github.com/reation/micro_service_stock/model"

	"github.com/reation/micro_service_stock/protoc"
	"github.com/reation/micro_service_stock/stock_op/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsStockByGoodsIDListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsStockByGoodsIDListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsStockByGoodsIDListLogic {
	return &GetGoodsStockByGoodsIDListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsStockByGoodsIDListLogic) GetGoodsStockByGoodsIDList(in *protoc.GetGoodsStockRequest) (*protoc.GetGoodsStockResponse, error) {
	// todo: add your logic here and delete this line
	states, goodsStockList, _ := l.GetGoodsStockList(in.GoodsID)
	if states != config.GOODS_STOCK_STATE_NORMAL {
		return &protoc.GetGoodsStockResponse{States: states, GoodsStock: nil}, nil
	}
	var resp = make([]*protoc.StockGoodsData, len(*goodsStockList))
	for _, v := range *goodsStockList {
		resp[v.GoodsId] = &protoc.StockGoodsData{
			GoodsId:  v.GoodsId,
			GoodsNum: v.GoodsNum,
		}
	}

	return &protoc.GetGoodsStockResponse{States: states, GoodsStock: resp}, nil
}

func (l *GetGoodsStockByGoodsIDListLogic) GetGoodsStockList(goodsIDList []*protoc.StockOpGoodsID) (int64, *[]model.GoodsStock, error) {
	var goodsIDS = make([]int64, len(goodsIDList))
	for k, v := range goodsIDList {
		goodsIDS[k] = v.GoodsId
	}
	goodsStockList, err := l.svcCtx.StockModel.GoodsStockModel.GetGoodsStockListByGoodsIDList(l.ctx, goodsIDS)
	if err == model.ErrNotFound {
		return config.GOODS_STOCK_STATE_EMPTY, nil, nil
	}
	if err != nil {
		return config.GOODS_STOCK_STATE_ERROR, nil, nil
	}

	return config.GOODS_STOCK_STATE_NORMAL, goodsStockList, nil
}
