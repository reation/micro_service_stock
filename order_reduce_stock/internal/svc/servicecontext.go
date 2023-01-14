package svc

import (
	"fmt"
	"github.com/reation/micro_service_stock/model"
	"github.com/reation/micro_service_stock/order_reduce_stock/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type StockDB struct {
	GoodsStockModel model.GoodsStockModel
	StockLogModel   model.GoodsStockLogModel
}

type ServiceContext struct {
	Config     config.Config
	StockModel StockDB
	Conn       sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSourceUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		c.Mysql.StockTable.User,
		c.Mysql.StockTable.Passwd,
		c.Mysql.StockTable.Addr,
		c.Mysql.StockTable.Port,
		c.Mysql.StockTable.DBName,
	)
	conn := sqlx.NewMysql(dataSourceUrl)
	StockConn := model.NewGoodsStockModel(conn)
	StockLogConn := model.NewGoodsStockLogModel(conn)
	return &ServiceContext{
		Config: c,
		StockModel: StockDB{
			GoodsStockModel: StockConn,
			StockLogModel:   StockLogConn,
		},
		Conn: conn,
	}
}
