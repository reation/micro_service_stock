package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro_service_stock/config"
	"time"

	"micro_service_stock/protoc"
	"micro_service_stock/stock_op/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderReturnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnLogic {
	return &OrderReturnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnLogic) OrderReturn(in *protoc.OrderReturnRequest) (*protoc.OrderReturnResponse, error) {
	// todo: add your logic here and delete this line
	err := l.returnGoodsStockNum(in.GetGoodsInfo(), in.GetOrderID())
	if err != nil {
		return &protoc.OrderReturnResponse{States: config.ORDER_INCREASE_STATES_ERROR}, nil
	}
	return &protoc.OrderReturnResponse{States: config.ORDER_INCREASE_STATES_NORMAL}, nil
}

func (l *OrderReturnLogic) returnGoodsStockNum(goodsInfoList []*protoc.StockOpGoodInfo, orderID int64) error {
	err := l.svcCtx.Conn.Transact(func(session sqlx.Session) error {
		var err error
		for _, v := range goodsInfoList {
			stockSql := fmt.Sprintf("update %s set `goods_num` = `goods_num` + %d where `id` = %d",
				"goods_stock", v.GetNum(), v.GetGoodsId())
			_, err = session.Exec(stockSql)
			if err != nil {
				return err
			}

			stockLogSql := fmt.Sprintf("insert into %s (%s) values (%d, %d, %d, %d, %d, \"%s\", \"%s\")",
				"goods_stock_log", "`goods_id`, `operation_type`, `operation_num`, `type`, `operation_id`, `update_time`, `create_time`",
				v.GetGoodsId(), config.STOCK_OPERATION_TYPE_INCREASE, v.GetNum(),
				config.TYPE_ORDER_INCREASE, orderID, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
			_, err = session.Exec(stockLogSql)
			if err != nil {
				return err
			}

		}

		return nil
	})

	return err
}
