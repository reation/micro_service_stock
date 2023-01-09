package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsStockLogModel = (*customGoodsStockLogModel)(nil)

type (
	// GoodsStockLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsStockLogModel.
	GoodsStockLogModel interface {
		goodsStockLogModel
	}

	customGoodsStockLogModel struct {
		*defaultGoodsStockLogModel
	}
)

// NewGoodsStockLogModel returns a model for the database table.
func NewGoodsStockLogModel(conn sqlx.SqlConn) GoodsStockLogModel {
	return &customGoodsStockLogModel{
		defaultGoodsStockLogModel: newGoodsStockLogModel(conn),
	}
}
