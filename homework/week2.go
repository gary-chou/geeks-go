/**
 * week 2
 * @Author: GaryChou
 * @Date: 2021/7/26 1:03 下午
 */
package homework

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/zhouzhimin/geeks-go/svc"
	"log"
	"time"
)

type User struct {
	Id        int
	Name      string
	Phone     string
	createdAt time.Time
	updatedAt time.Time
}

func GetUserRowsByPhone(phone string) (users []User, err error) {
	//获取数据库实例
	dbIns, dbErr := svc.GetDBIns()
	if dbErr != nil {
		return nil, errors.New(fmt.Sprintf("Get DBInstance Fail: %s", dbErr.Error()))
	}
	//查询
	sqlStr := "select id,name,phone from `user` where `phone` = ?"
	rows, queryErr := dbIns.Query(sqlStr, phone)
	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
			log.Println(fmt.Sprintf("LOG INFO: DBIns Query Rows: 「%s」, sql: 「%s」, params: 「%s」", sql.ErrNoRows, sqlStr, phone))
			return users, nil
		}
		return nil, errors.New(fmt.Sprintf("DBIns Query Fail: %s", queryErr.Error()))
	}
	defer rows.Close()

	//遍历
	for rows.Next() {
		var user User
		scanErr := rows.Scan(&user.Id, &user.Name, &user.Phone)
		if scanErr != nil {
			return nil, errors.New(fmt.Sprintf("DBRows Scan Fail: %s", scanErr.Error()))
		}
		users = append(users, user)
	}
	return
}
