package logic

import (
	"context"
	"fmt"
	"github.com/reation/micro_service_stock/config"
	"github.com/reation/micro_service_stock/order_reduce_stock/internal/svc"
	"github.com/reation/micro_service_stock/protoc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type OrderReduceStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReduceStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReduceStockLogic {
	return &OrderReduceStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReduceStockLogic) OrderReduceStock(in *protoc.OrderReduceStockRequest) (*protoc.OrderReduceStockResponse, error) {
	// todo: add your logic here and delete this line
	if in.IsAll == config.IS_ALL_YES {
		resp := l.isAllYes(in.GoodsID, in.OrderID)
		return &protoc.OrderReduceStockResponse{States: config.ORDER_REDUCE_STATES_NORMAL, GoodsStates: resp}, nil
	}

	resp, states := l.isAllNo(in.GoodsID, in.OrderID)

	return &protoc.OrderReduceStockResponse{States: states, GoodsStates: resp}, nil
}

func (l *OrderReduceStockLogic) isAllNo(goodsInfoList []*protoc.OrderGoodsID, orderID int64) (res []*protoc.OrderReturn, states int64) {
	var resp = make([]*protoc.OrderReturn, len(goodsInfoList))
	var goodsIDList = make([]int64, len(goodsInfoList))
	states = config.ORDER_REDUCE_STATES_NORMAL
	for k, v := range goodsInfoList {
		goodsIDList[k] = v.GetGoodsId()
	}
	stockList, err := l.svcCtx.StockModel.GoodsStockModel.GetGoodsStockInfoByGoodsIDList(l.ctx, goodsIDList)
	if err != nil {
		return resp, config.ORDER_REDUCE_STATES_ERROR
	}
	if len(stockList) != len(goodsIDList) {
		return resp, config.ORDER_REDUCE_STATES_ERROR
	}
	err = l.svcCtx.Conn.Transact(func(session sqlx.Session) error {
		var err error
		num := 0
		transact := true

		for _, val := range goodsInfoList {
			resp[num] = &protoc.OrderReturn{
				GoodsId: val.GetGoodsId(),
				Demand:  val.GetGoodsNum(),
				Stock:   stockList[val.GetGoodsId()].GoodsNum,
				Deal:    0,
			}
			if val.GetGoodsNum() > stockList[val.GetGoodsId()].GoodsNum {
				states = config.ORDER_REDUCE_STATES_LESS
				resp[num].Deal = 0
				transact = false
				num++
				continue
			}
			if transact == true {
				stockSql := fmt.Sprintf("update %s set `goods_num` = `goods_num` - %d where `id` = %d and `goods_num` = %d ",
					"goods_stock", val.GetGoodsNum(), stockList[val.GetGoodsId()].GoodsId, stockList[val.GetGoodsId()].GoodsNum)
				_, err = session.Exec(stockSql)
				if err != nil {
					transact = false
					continue
				}
				stockLogSql := fmt.Sprintf("insert into %s (%s) values (%d, %d, %d, %d, %d, \"%s\", \"%s\")",
					"goods_stock_log", "`goods_id`, `operation_type`, `operation_num`, `type`, `operation_id`, `update_time`, `create_time`",
					stockList[val.GetGoodsId()].GoodsId, config.STOCK_OPERATION_TYPE_REDUCE, val.GetGoodsNum(),
					config.TYPE_ORDER_REDUCE, orderID, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
				_, err = session.Exec(stockLogSql)
				if err != nil {
					transact = false
					continue
				}
				resp[num].Deal = val.GetGoodsNum()
				num++
			}

		}
		if transact == true {
			return nil
		}

		return err
	})
	if err != nil && states == config.ORDER_REDUCE_STATES_NORMAL {
		states = config.ORDER_REDUCE_STATES_ERROR
	}

	return resp, states
}

func (l *OrderReduceStockLogic) isAllYes(goodsInfoList []*protoc.OrderGoodsID, orderID int64) []*protoc.OrderReturn {
	var goodsStatesList []*protoc.OrderReturn

	for i, v := range goodsInfoList {

		goodsStockInfo, err := l.svcCtx.StockModel.GoodsStockModel.GetGoodsStockInfoByGoodsID(l.ctx, v.GoodsId)
		if err != nil {
			goodsStatesList[i] = &protoc.OrderReturn{
				GoodsId: v.GoodsId,
				Demand:  v.GoodsNum,
				Stock:   0,
				Deal:    0,
			}
			continue
		}
		dealNum := v.GoodsNum
		if goodsStockInfo.GoodsNum < v.GoodsNum {
			dealNum = goodsStockInfo.GoodsNum
		}

		err = l.svcCtx.StockModel.GoodsStockModel.OrderReduceGoodsStockByGoods(l.ctx, v.GoodsId, dealNum)
		if err != nil {
			goodsStatesList[i] = &protoc.OrderReturn{
				GoodsId: v.GoodsId,
				Demand:  v.GoodsNum,
				Stock:   0,
				Deal:    0,
			}
			continue
		}
		goodsStatesList[i] = &protoc.OrderReturn{
			GoodsId: v.GoodsId,
			Demand:  v.GoodsNum,
			Stock:   goodsStockInfo.GoodsNum,
			Deal:    dealNum,
		}
		l.svcCtx.StockModel.StockLogModel.AddOrderReduceLog(l.ctx, v.GoodsId, dealNum, orderID)
	}

	return goodsStatesList
}
