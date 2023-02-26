# stitching

# stitching

## Introduction

A simple sql query splicing library, because my sql is not very good, so I may only develop it on the basis of satisfying my own business at present

## How to use

Install!

```shell
go get -u github.com/C-RiceCongee/stitching
```

```go
package models

type GetPostListByFilterParams struct {
	Title       string `json:"title,omitempty" form:"title,omitempty" db:"title"`
	CommunityId string `json:"community_id,omitempty" form:"community_id,omitempty" db:"community_id"`
	AuthorId    string `json:"author_id,omitempty" form:"author_id,omitempty" db:"author_id"`
	PageSize    int    `form:"page_size" json:"page_size" binding:"required"`
	PageNumber  int    `form:"page_number" json:"page_number" binding:"required"A`
	Sort        string `form:"sort" json:"sort" binding:"required"`
}
```

```go
var params = models.GetPostListByFilterParams{
	AuthorId: "123123",
	Title:    "123123",
}
sq1, sq2, namedQueryMap, namedQuerySlice := ConnectWhereConditions(params)
fmt.Println(where)
fmt.Println(queryMap)
fmt.Println(namedQueryMap)
fmt.Println(namedQuerySlice)
```

```shell
sq1             >  where title like :title and author_id = :author_id  # For sql named query
namedQueryMap   >  map[author_id:123123 title:123123] # For sql named query ,such as sq1

sq2             >	where title like ? and author_id = ? # Compatible with database/sql standard library queries
namedQuerySlice > [123123 123123]
```

```go
import (
  	...
   "github.com/C-RiceCongee/stitching/where"
)
```

## namedQuery

```go
whereCDS, _, namedQueryMap, _ := where.ConnectWhereConditions[parameters.GetPostListByFilterParams](*params)
namedQueryMap["offset"] = params.PageSize * (params.PageNumber - 1)
namedQueryMap["limit"] = params.PageSize
sql := fmt.Sprintf("select post_id,author_id,status,create_time,update_time,title,community_id from post %s limit :limit offset :offset;", whereCDS)
rows, err := setup.MysqlDB.NamedQuery(sql, namedQueryMap)
```

## GetQuery

```go
_, sqlWhereSlice, _, nameQuerySlice := where.ConnectWhereConditions[parameters.GetPostListByFilterParams](*params)
sql := fmt.Sprintf("select count(id) from post %s", sqlWhereSlice)
err := setup.MysqlDB.Get(&count, sql, nameQuerySlice...)
```



> This library will help you fill in the sql related to your structure and the parameter values of the query



## SUPPORT LIKE

You can use custom tags `sql` 

However, at present, is relatively simple and only supports `like` and `=`

By default, when nothing is configured, it is `=`

```diff
package models

type GetPostListByFilterParams struct {
-	Title       string `json:"title,omitempty" form:"title,omitempty" db:"title"`
+ Title       string `json:"title,omitempty" form:"title,omitempty" db:"title"  sql:"like"``
	CommunityId string `json:"community_id,omitempty" form:"community_id,omitempty" db:"community_id"`
	AuthorId    string `json:"author_id,omitempty" form:"author_id,omitempty" db:"author_id"`
	PageSize    int    `form:"page_size" json:"page_size" binding:"required"`
	PageNumber  int    `form:"page_number" json:"page_number" binding:"required"A`
	Sort        string `form:"sort" json:"sort" binding:"required"`
}
```

Now he generates statements and parameters like this

```shell
where title like :title and author_id = :author_id

where title like ? and author_id = ?

map[author_id:123123 title:%123123%]

[%123123% 123123]
```

