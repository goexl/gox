package gox

import (
	`database/sql`
	`errors`
	`fmt`
	`strings`
)

type (
	// JSONInitializer JSON初始化者
	JSONInitializer interface {
		// InitSQL 初始化SQL
		InitSQL(table string, field string) (string, error)
		// IsInitialized 是否已经初始化完成
		IsInitialized() bool
	}

	// JSONFielder JSON字段
	JSONFielder interface {
		// InitializeField 初始化字段
		InitializeField() string
	}

	// JSONInitialized JSON是否初始化
	JSONInitialized struct {
		// Initialized 是否初始化
		Initialized bool `json:"initialized,omitempty"`
	}
)

func (ji JSONInitialized) IsInitialized() bool {
	return ji.Initialized
}

func (ji JSONInitialized) InitializeField() string {
	return "initialized"
}

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

	if _, err = sqlBuilder.WriteString(fmt.Sprintf(`UPDATE %v SET %v = JSON_SET(COALESCE(%v, '{}')`, table, filed, filed)); nil != err {
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

// JSONInit JSON字段初始化
func JSONInit(db *sql.DB, initializer JSONInitializer, table string, field string) (err error) {
	if initializer.IsInitialized() {
		return
	}

	var (
		execSQL string
		result  sql.Result
	)

	if execSQL, err = initializer.InitSQL(table, field); nil != err {
		return
	}

	if result, err = db.Exec(execSQL); nil != err {
		return
	}
	if _, err = result.RowsAffected(); nil != err {
		return
	}

	return
}

// MySQLJsonInit JSON初始化
func MySQLJsonInit(
	table string, filed string,
	conditionFiled string, conditionValue interface{},
	initField string, initFieldValue interface{},
	paths ...string,
) (sql string, err error) {
	sqlBuilder := strings.Builder{}
	if _, err = sqlBuilder.WriteString(fmt.Sprintf("UPDATE %s SET %s = JSON_SET(%s,", table, filed, filed)); nil != err {
		return
	}

	for _, path := range paths {
		if _, err = sqlBuilder.WriteString(fmt.Sprintf(
			"'$.%s', JSON_OBJECT('%s', %v)",
			path,
			initField,
			initFieldValue,
		)); nil != err {
			return
		}
	}
	if _, err = sqlBuilder.WriteString(fmt.Sprintf(`) WHERE %v = '%v'`, conditionFiled, conditionValue)); nil != err {
		return
	}

	sql = sqlBuilder.String()

	return
}
