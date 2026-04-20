package driver

import (
	_gorm "github.com/mazezen/ginframe/server-common/pkg/gorm"
	_redis "github.com/mazezen/ginframe/server-common/pkg/redis"
	"gopkg.in/redis.v5"
	"gorm.io/gorm"
)

// CreateDb create mysql db
func CreateDb(dbDsn string, maxOpenConn, maxIdleConn int) (*gorm.DB, error) {
	arg := []int{maxOpenConn, maxIdleConn}
	db, err := _gorm.InitGormMysql(dbDsn, arg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateRedis 初始化redis 连接
func CreateRedis(addr, pass string, rdb int) (*redis.Client, error) {
	rd, err := _redis.InitRedis(addr, pass, rdb)
	if err != nil {
		return nil, err
	}
	return rd, nil
}
