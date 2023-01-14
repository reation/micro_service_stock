package logic

import (
	"context"
	"fmt"
	"github.com/reation/micro_service_stock/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"github.com/reation/micro_service_stock/protoc"
	"github.com/reation/micro_service_stock/stock_op/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SupplierIncreaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSupplierIncreaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SupplierIncreaseLogic {
	return &SupplierIncreaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SupplierIncreaseLogic) SupplierIncrease(in *protoc.SupplierIncreaseRequest) (*protoc.SupplierIncreaseResponse, error) {
	// todo: add your logic here and delete this line
	err := l.supplierIncreaseStock(in.GetGoodsInfo(), in.GetSupplierID())
	if err != nil {
		return &protoc.SupplierIncreaseResponse{States: config.SUPPLIER_INCREASE_STATES_ERROR}, nil
	}

	return &protoc.SupplierIncreaseResponse{States: config.SUPPLIER_INCREASE_STATES_NORMAL}, nil
}

func (l *SupplierIncreaseLogic) supplierIncreaseStock(goodsInfoList []*protoc.StockOpGoodInfo, supplierID int64) error {
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
				config.TYPE_SUPPLIER_INCREASE, supplierID, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
			_, err = session.Exec(stockLogSql)
			if err != nil {
				return err
			}

		}

		return nil
	})

	return err
}
