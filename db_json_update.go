package gox

import (
	`errors`
	`fmt`
)

// MySQLJsonUpdateWithConfig 生成MySQL JSON增量更新SQL语句
func MySQLJsonUpdateWithConfig(
	table string, filed string,
	conditionFiled string, conditionValue interface{},
	data map[string]interface{},
	prefix string, style SeparatorStyle,
) (sql string, err error) {
	var keyMap map[string]interface{}

	// 检查配置是否为空
	// 如果为空，不再继续处理，会报SQL语句错误
	if 0 == len(data) {
		err = errors.New("配置数据为空")
		return
	}

	sql = fmt.Sprintf(`UPDATE %v SET %v = JSON_SET(%v`, table, filed, filed)
	if keyMap, err = Flatten(data, prefix, style); nil != err {
		return
	}

	for k, v := range keyMap {
		switch v.(type) {
		case string:
			sql += fmt.Sprintf(`,'$.%v','%v'`, k, v)
		default:
			sql += fmt.Sprintf(`,'$.%v',%v`, k, v)
		}
	}

	sql += fmt.Sprintf(`) WHERE %v = '%v'`, conditionFiled, conditionValue)

	return
}

// MySQLJsonUpdate 生成MySQL JSON增量更新SQL语句的快捷方式
func MySQLJsonUpdate(
	table string, filed string,
	conditionFiled string, conditionValue interface{},
	data map[string]interface{},
) (sql string, err error) {
	return MySQLJsonUpdateWithConfig(table, filed, conditionFiled, conditionValue, data, "", DotStyle)
}
