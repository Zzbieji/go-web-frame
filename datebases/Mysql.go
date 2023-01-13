package datebases

import (
	. "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitMysql() *gorm.DB {
	config := NewConfig()
	config.DBName = "test"
	config.User = "root"
	config.Passwd = "xxxxxx"
	config.Net = "tcp"
	config.Addr = "xxxxxxxx"
	config.Params["charset"] = "utf8mb4"
	config.Params["parseTime"] = "True"
	config.Params["loc"] = "Local"
	dsn := config.FormatDSN()
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "go-sql-driver/mysql",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db

}
