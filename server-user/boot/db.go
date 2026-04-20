package boot

import (
	"github.com/mazezen/ginframe/server-common/driver"
	"github.com/mazezen/ginframe/server-user/core"
	"github.com/mazezen/golog"

	"log"
)

// InitDb init gorm
func InitDb(dbDsn string, maxOpenConn, maxIdleConn int) {
	gdb, err := driver.CreateDb(dbDsn, maxOpenConn, maxIdleConn)
	if err != nil {
		log.Fatalf("Create mysql connection err: %v\n", err)
	}
	core.SetMysql(gdb)

	golog.Logger.Info("🚀🚀🚀🚀🚀🚀mysql success...🚀🚀🚀🚀🚀🚀")
}

// InitRedis init redis
func InitRedis(addr, pass string, rdb int) {
	rd, err := driver.CreateRedis(addr, pass, rdb)
	if err != nil {
		log.Fatalf("Create redis client err: %v\n", err)
	}
	core.SetRedis(rd)

	golog.Logger.Info("🚀🚀🚀🚀🚀🚀redis success...🚀🚀🚀🚀🚀🚀")
}

// InitMongo init mongodb
func InitMongo(addr string) {
	mongo, err := driver.InitMongo(addr)
	if err != nil {
		log.Fatalf("Create mongodb client err: %v\n", err)
	}
	core.SetMongo(mongo)
	golog.Logger.Info("🚀🚀🚀🚀🚀🚀 mongodb success...🚀🚀🚀🚀🚀🚀")
}

// InitLevelDb init leveldb
func InitLevelDb(path string) {
	ldb := driver.InitLevelDb(path)
	core.SetLevelDB(ldb)

	golog.Logger.Info("🚀🚀🚀🚀🚀🚀levedb success...🚀🚀🚀🚀🚀🚀")
}
