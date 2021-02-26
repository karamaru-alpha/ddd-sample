package mysql

import (
	"fmt"
	"time"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Conn 各repositoryで利用するDB接続(Connection)情報
var Conn *gorm.DB

func init() {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(ddd-sample-db:3306)"
	DBNAME := "ddd_sample"

	CONNECT := fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8mb4&parseTime=true&loc=Local",
		USER,
		PASS,
		PROTOCOL,
		DBNAME,
	)

	count := 0
	var err error
	Conn, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			Conn, err = gorm.Open(DBMS, CONNECT)
		}
	}

	Conn.LogMode(true)
	Conn.Set("gorm:table_options", "ENGINE=InnoDB")

	fmt.Println("DB接続成功")
}
