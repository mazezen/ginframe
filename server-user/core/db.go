package core

import (
	"gopkg.in/redis.v5"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
	Rd *redis.Client
)

func SetMysql(e *gorm.DB) {
	Db = e
}

func SetRedis(_rd *redis.Client) {
	Rd = _rd
}
