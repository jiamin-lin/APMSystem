package dogapm

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/redis/go-redis/v9"
)

// 基础设施
type infra struct {
	// 基础设施名称
	Db  *sql.DB
	Rdb *redis.Client
}

// NewInfra 创建基础设施
var Infra = &infra{}

// InfraOption 基础设施选项
type InfraOption func(i *infra)

// WithDB 设置DB数据库
func InfraDbOption(connectUrl string) InfraOption {
	return func(i *infra) {
		var err error
		i.Db, err = sql.Open("mysql", connectUrl)
		if err != nil {
			panic(err)
		}
		err = i.Db.Ping()
		if err != nil {
			panic(err)
		}
	}
}

// WithRdb 设置Rdb数据库
func InfraRDBOption(addr string) InfraOption {
	return func(i *infra) {
		rdb := redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   0,
		})
		res, err := rdb.Ping(context.TODO()).Result()
		if err != nil {
			panic(err)
		}
		if res != "PONG" {
			panic("redis connect fail")
		}
		i.Rdb = rdb
	}
}

// InitInfra 初始化基础设施
func (i *infra) InitInfra(options ...InfraOption) {
	for _, opt := range options {
		opt(i)
	}
}
