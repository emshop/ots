package dbs

import (
	"fmt"
	"strings"

	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Query 将前一流程的结果作为后序流程的输入，批量执行一组SQL语句
func Query(db db.IDBExecuter, input types.IXMap, sqls ...string) (types.XMaps, error) {
	if len(sqls) == 0 {
		return nil, fmt.Errorf("未传入任何SQL语句")
	}
	for i, sql := range sqls {
		if len(sql) < 6 {
			return nil, fmt.Errorf("sql语句错误%s", sql)
		}
		prefix := strings.Trim(strings.TrimSpace(strings.TrimLeft(sql, "\n")), "\t")[:6]
		switch strings.ToUpper(prefix) {
		case "SELECT":
			outputs, err := db.Query(sql, input.ToMap())
			if err != nil {
				return nil, err
			}
			if outputs.Len() == 0 {
				return nil, fmt.Errorf("%s数据不存在%w input:%+v", sql, xerr.ErrNOTEXISTS, input)
			}
			input.Merge(outputs.Get(0))
			if i == len(sqls)-1 {
				return outputs, nil
			}
		default:
			return nil, fmt.Errorf("不支持的SQL语句，或SQL语句前包含有特殊字符:%s", sql)
		}
	}
	return nil, nil
}

//Executes 将前一流程的结果作为后序流程的输入，批量执行一组SQL语句
func Executes(db db.IDBExecuter, input types.IXMap, sqls ...string) (types.IXMap, error) {
	if len(sqls) == 0 {
		return nil, fmt.Errorf("未传入任何SQL语句")
	}
	for _, sql := range sqls {

		if len(sql) < 6 {
			return nil, fmt.Errorf("sql语句错误%s", sql)
		}
		prefix := strings.Trim(strings.TrimSpace(strings.TrimLeft(sql, "\n")), "\t")[:6]
		switch strings.ToUpper(prefix) {
		case "SELECT":
			outputs, err := db.Query(sql, input.ToMap())
			if err != nil {
				return nil, err
			}
			if outputs.Len() == 0 {
				return nil, fmt.Errorf("%s数据不存在%w input:%+v", sql, xerr.ErrNOTEXISTS, input)
			}
			input.Merge(outputs.Get(0))
		case "UPDATE", "INSERT":
			rows, err := db.Execute(sql, input.ToMap())
			if err != nil {
				return nil, err
			}
			if rows == 0 {
				return nil, fmt.Errorf("%s数据修改失败%w input:%+v", sql, xerr.ErrNOTEXISTS, input)
			}
		default:
			return nil, fmt.Errorf("不支持的SQL语句，或SQL语句前包含有特殊字符:%s", sql)
		}
	}
	return input, nil
}
