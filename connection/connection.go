package connection

import (
	// "gorm.io/driver/sqlserver"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var (
	DB *gorm.DB
	Err error
)

func Connect(){
	// dsn := "sqlserver://digitalporto:Bre@k2312@172.18.132.122:1433?database=DB_DIGIPORTO_TEST"
	// dsn := "sqlserver://root:12345@localhost:1433?database=DB_DIGIPORTO_TEST"
	dsn := "root:@tcp(localhost:3306)/apm?charset=utf8mb4&parseTime=true&loc=Local"
	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Err != nil {
		panic("failed to connect database")
	}
	// return db, err
}