package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsStockModel = (*customGoodsStockModel)(nil)

type (
	// GoodsStockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsStockModel.
	GoodsStockModel interface {
		goodsStockModel
	}

	customGoodsStockModel struct {
		*defaultGoodsStockModel
	}
)

// NewGoodsStockModel returns a model for the database table.
func NewGoodsStockModel(conn sqlx.SqlConn) GoodsStockModel {
	return &customGoodsStockModel{
		defaultGoodsStockModel: newGoodsStockModel(conn),
	}
}
