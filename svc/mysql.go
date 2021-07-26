/**
 * service
 * @Author: GaryChou
 * @Date: 2021/7/26 1:19 下午
 */
package svc

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	host     = "127.0.0.1:3306"
	username = "root"
	password = "zzm666"
	dbName   = "geek"
)

func GetDBIns() (dbIns *sql.DB, err error) {
	//简易版
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", username, password, host, dbName)
	dbIns, err = sql.Open("mysql", dbInfo)
	if err != nil {
		return nil, err
	}
	dbIns.SetConnMaxLifetime(100)
	dbIns.SetMaxIdleConns(10)
	//TODO 单例、连接池、并发控制、超时控制、容错
	return
}
