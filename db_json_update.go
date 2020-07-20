package gox

import (
	`errors`
	`fmt`
)

func DBJsonUpdate(updateTableName, updateFiledName, conditionFiledName string, condition interface{}, conf map[string]interface{}, prefix string, style SeparatorStyle) (error, string) {

	// 检查配置是否为空
	// 如果为空，不再继续处理，会报SQL语句错误
	if 0 == len(conf) {
		return errors.New("cofing is nil"), ""
	}
	sqlStr := fmt.Sprintf(`UPDATE %v SET %v = JSON_SET(%v`, updateTableName, updateFiledName)
	if keyMap, err := Flatten(conf, prefix, style); nil != err {
		return err, ""
	} else {
		for k, v := range keyMap {
			switch v.(type) {
			case string:
				sqlStr += fmt.Sprintf(`,'$.%v','%v'`, k, v)
			default:
				sqlStr += fmt.Sprintf(`,'$.%v',%v`, k, v)
			}
		}
	}
	sqlStr += fmt.Sprintf(`) WHERE %v = '%v'`, conditionFiledName, condition)
	return nil, sqlStr
}
