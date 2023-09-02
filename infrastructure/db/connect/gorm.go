package connect

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	sLog "log"
	"os"
	"time"
)

func NewDB() *gorm.DB {
	var loggerAdapter logger.Interface
	if config.Config.Debug {
		loggerAdapter = logger.New(
			//将标准输出作为Writer
			sLog.New(os.Stdout, "\r\n", sLog.LstdFlags),
			logger.Config{
				//设定慢查询时间阈值为1ms
				SlowThreshold: 1 * time.Microsecond,
				//设置日志级别，只有Warn和Info级别会输出慢查询日志
				LogLevel: logger.Info,
			},
		)
	} else {
		loggerAdapter = newGormLogger()
	}
	gormConf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: loggerAdapter,
	}

	db, err := gorm.Open(mysql.Open(config.Config.Mysql.User), gormConf)
	if err != nil {
		panic(err)
	}
	conn, err := db.DB()
	if err != nil {
		//logx.Channel(logx.Default).Error("获取MySQL连接错误", err)
		//os.Exit(1)
		panic("获取MySQL连接错误" + err.Error())
	}
	conn.SetMaxIdleConns(config.Config.Mysql.MaxIdleConns)
	conn.SetMaxOpenConns(config.Config.Mysql.MaxOpenConns)
	_ = db.Use(&TracePlugin{})
	return db
}
