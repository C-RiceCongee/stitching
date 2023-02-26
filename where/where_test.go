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
	where, queryMap, namedQueryMap, namedQuerySlice := ConnectWhereConditions(params)
	fmt.Println(where)
	fmt.Println(queryMap)
	fmt.Println(namedQueryMap)
	fmt.Println(namedQuerySlice)
}
