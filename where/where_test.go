package where

import (
	"fmt"
	"testing"

	"github.com/C-RiceCongee/stitching/models"
)

func TestConnectWhereConditions(t *testing.T) {
	var params = models.GetPostListByFilterParams{
		AuthorId: "123123",
		Title:    "123123",
	}

	/*
		1. 反射不能使用指针类型！

	*/
	where, queryMap, slice, _ := ConnectWhereConditions[models.GetPostListByFilterParams](params)
	fmt.Println(where)
	fmt.Println(queryMap)
	fmt.Println(slice)
}
