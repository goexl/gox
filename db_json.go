package gox

import (
	`errors`
	`fmt`
	`strings`
)

const (
	// JSONInitialStatusUnInitial 未初始化
	JSONInitialStatusUnInitial JSONInitialStatus = 0
	// JSONInitialStatusJSONInitialized 已初始化
	JSONInitialStatusJSONInitialized JSONInitialStatus = 1
)

type (
	// JSONInitialStatus JSON初始化状态
	JSONInitialStatus uint8

	// JSONInitial JSON是否初始化
	JSONInitial struct {
		// Status 是否初始化
		Status JSONInitialStatus `json:"status,omitempty"`
	}
)

// MySQLJsonUpdateWithConfig 生成MySQL JSON增量更新SQL语句
func MySQLJsonUpdateWithConfig(
	table string, filed string,
	conditionFiled string, conditionValue interface{},
	data map[string]interface{},
	prefix string, style SeparatorStyle,
) (sql string, err error) {
	var (
		sqlBuilder = strings.Builder{}
		keyMap     map[string]interface{}
	)

	// 检查配置是否为空
	// 如果为空，不再继续处理，会报SQL语句错误
	if 0 == len(data) {
		err = errors.New("配置数据为空")

		return
	}

	if _, err = sqlBuilder.WriteString(fmt.Sprintf(`UPDATE %v SET %v = JSON_SET(%v`, table, filed, filed)); nil != err {
		return
	}
	if keyMap, err = Flatten(data, prefix, style); nil != err {
		return
	}

	for k, v := range keyMap {
		switch v.(type) {
		case string:
			_, err = sqlBuilder.WriteString(fmt.Sprintf(`,'$.%v','%v'`, k, v))
		default:
			_, err = sqlBuilder.WriteString(fmt.Sprintf(`,'$.%v',%v`, k, v))
		}

		if nil != err {
			return
		}
	}

	if _, err = sqlBuilder.WriteString(fmt.Sprintf(`) WHERE %v = '%v'`, conditionFiled, conditionValue)); nil != err {
		return
	}

	sql = sqlBuilder.String()

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

// MySQLJsonInit JSON初始化
func MySQLJsonInit(table string, filed string, paths ...string) (sql string, err error) {
	sqlBuilder := strings.Builder{}
	if _, err = sqlBuilder.WriteString(fmt.Sprintf("UPDATE %s SET %s = JSON_SET(%s,", table, filed, filed)); nil != err {
		return
	}

	for _, path := range paths {
		if _, err = sqlBuilder.WriteString(fmt.Sprintf(
			"'$.%s', JSON_OBJECT('status', %d)",
			path,
			JSONInitialStatusJSONInitialized,
		)); nil != err {
			return
		}
	}
	if _, err = sqlBuilder.WriteString(")"); nil != err {
		return
	}

	sql = sqlBuilder.String()

	return
}
