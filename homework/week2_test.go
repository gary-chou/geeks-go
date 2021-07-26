/**
 * week 2 test
 * @Author: GaryChou
 * @Date: 2021/7/26 2:13 下午
 */
package homework

import (
	"fmt"
	"log"
	"testing"
)

func TestGetUserRowsByPhone(t *testing.T) {
	phone := "199123456781"
	users, err := GetUserRowsByPhone(phone)
	if err != nil {
		log.Println(fmt.Sprintf("GetUserRowsByPhone Fail: %s, params: %s", err.Error(), phone))
		return
	}
	fmt.Println(users)
}
