// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"micro_service_stock/tool"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	goodsStockFieldNames          = builder.RawFieldNames(&GoodsStock{})
	goodsStockRows                = strings.Join(goodsStockFieldNames, ",")
	goodsStockRowsExpectAutoSet   = strings.Join(stringx.Remove(goodsStockFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")
	goodsStockRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsStockFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), "=?,") + "=?"
)

type (
	goodsStockModel interface {
		Insert(ctx context.Context, data *GoodsStock) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*GoodsStock, error)
		Update(ctx context.Context, data *GoodsStock) error
		Delete(ctx context.Context, id int64) error
		OrderReduceGoodsStockByGoods(ctx context.Context, id, goodsNum int64) error
		GetGoodsStockInfoByGoodsID(ctx context.Context, goodsID int64) (*GoodsStock, error)
		GetGoodsStockInfoByGoodsIDList(ctx context.Context, goodsID []int64) (map[int64]GoodsStock, error)
	}

	defaultGoodsStockModel struct {
		conn  sqlx.SqlConn
		table string
		Table string
	}

	GoodsStock struct {
		Id            int64     `db:"id"`
		GoodsId       int64     `db:"goods_id"` // 商品ID
		GoodsNum      int64     `db:"goods_num"`
		GoodsAlertNum int64     `db:"goods_alert_num"` // 商品报警数据量
		UpdateTime    time.Time `db:"update_time"`     // 更新时间
		CreateTime    time.Time `db:"create_time"`     // 创建时间
	}
)

func newGoodsStockModel(conn sqlx.SqlConn) *defaultGoodsStockModel {
	return &defaultGoodsStockModel{
		conn:  conn,
		table: "`goods_stock`",
		Table: "`goods_stock`",
	}
}

func (m *defaultGoodsStockModel) GetGoodsStockInfoByGoodsIDList(ctx context.Context, goodsID []int64) (map[int64]GoodsStock, error) {
	ids := strings.Join(tool.IntToStringArr(goodsID), ",")
	query := fmt.Sprintf("select %s from %s where `goods_id` in (%s)", goodsStockRows, m.table, ids)
	var stockInfo []GoodsStock
	var resp = make(map[int64]GoodsStock)
	err := m.conn.QueryRowsCtx(ctx, &stockInfo, query)
	if err == sqlc.ErrNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	for _, v := range stockInfo {
		fmt.Println("-----------mysql-----------------")
		fmt.Println(v)
		fmt.Println("-----------mysql-----------------")
		resp[v.GoodsId] = v
	}
	fmt.Println("-----------mysqlq-----------------")
	fmt.Println(resp[2])
	fmt.Println(resp[3])
	fmt.Println("-----------mysqlq-----------------")
	return resp, nil
}

func (m *defaultGoodsStockModel) GetGoodsStockInfoByGoodsID(ctx context.Context, goodsID int64) (*GoodsStock, error) {
	query := fmt.Sprintf("select %s from %s where `goods_id` = ? limit 1", goodsStockRows, m.table)
	var resp GoodsStock
	err := m.conn.QueryRowCtx(ctx, &resp, query, goodsID)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGoodsStockModel) OrderReduceGoodsStockByGoods(ctx context.Context, id, goodsNum int64) error {
	query := fmt.Sprintf("update %s set `goods_num` = `goods_num` - ? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, goodsNum, id)
	return err
}

func (m *defaultGoodsStockModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGoodsStockModel) FindOne(ctx context.Context, id int64) (*GoodsStock, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", goodsStockRows, m.table)
	var resp GoodsStock
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGoodsStockModel) Insert(ctx context.Context, data *GoodsStock) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, goodsStockRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.GoodsId, data.GoodsNum, data.GoodsAlertNum)
	return ret, err
}

func (m *defaultGoodsStockModel) Update(ctx context.Context, data *GoodsStock) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, goodsStockRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.GoodsId, data.GoodsNum, data.GoodsAlertNum, data.Id)
	return err
}

func (m *defaultGoodsStockModel) tableName() string {
	return m.table
}