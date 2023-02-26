package where

import (
	"fmt"
	"reflect"
	"strings"
)

// ConnectWhereConditions
//
//	 	1.反射拼装sql！
//		2.Map 适配sql标准库的 GET ，Select  nameexec  namedquery Slice 适配多参数一般查询 展开切片~
func ConnectWhereConditions[T any](params T) (
	sqlWhereMap string,
	sqlWhereSlice string,
	namedQueryMap map[string]interface{},
	namedQuerySlice []interface{},
) {
	// 通过反射推断 值和Type
	v := reflect.ValueOf(params)
	t := reflect.TypeOf(params)
	namedQueryMap = make(map[string]interface{})
	namedQuerySlice = make([]interface{}, 0)
	addIdx := 0
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		if field.IsValid() {
			// 字段值类型！
			kind := field.Kind()
			fieldDBName := t.Field(i)
			//获取 tag 这个一般是指定的对应的数据库的字段！
			tagDBName := fieldDBName.Tag.Get("db")
			if len(tagDBName) == 0 {
				continue
			}
			tagSQLName := fieldDBName.Tag.Get("sql")
			switch kind {
			case reflect.String:
				if len(field.String()) != 0 {
					// 根据 sql tag 拿到查询的方式！是 = 还是 like !
					fieldRealValue := v.FieldByName(fieldDBName.Name).String()
					// tagSQLName 没有代码默认传递长度！
					fmt.Println(tagDBName)
					if len(tagSQLName) == 0 {
						tagSQLName = "="
						namedQueryMap[tagDBName] = fieldRealValue
						namedQuerySlice = append(namedQuerySlice, fieldRealValue)
					} else {
						// 目前支持like
						namedQueryMap[tagDBName] = "%" + fieldRealValue + "%"
						namedQuerySlice = append(namedQuerySlice, "%"+fieldRealValue+"%")
					}
					if addIdx == 0 {
						sqlWhereMap += fmt.Sprintf("where %s %s :%s", tagDBName, tagSQLName, tagDBName)
						sqlWhereSlice += fmt.Sprintf("where %s %s ?", tagDBName, tagSQLName)
						addIdx += 1
					} else {
						sqlWhereMap += fmt.Sprintf(" and %s %s :%s", tagDBName, tagSQLName, tagDBName)
						sqlWhereSlice += fmt.Sprintf(" and %s %s ?", tagDBName, tagSQLName)
					}
				} else {
					continue
				}
			}
		}
		// 2. 获取 字段的名称！fieldName
	}
	fmt.Println(namedQuerySlice)
	return strings.Trim(sqlWhereMap, "\n"), sqlWhereSlice, namedQueryMap, namedQuerySlice
}
