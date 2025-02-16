package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// MapModelToVO 将模型数据映射到对应的 VO，返回具体类型的指针
func MapModelToVO(modelData interface{}, voType interface{}) (interface{}, error) {
	modelValue := reflect.ValueOf(modelData)
	voValue := reflect.ValueOf(voType)

	if voValue.Kind() != reflect.Ptr || voValue.IsNil() {
		return nil, fmt.Errorf("voType 必须为指向结构体的指针")
	}
	voValue = voValue.Elem()

	// 如果 modelData 是指针类型，解引用获取实际的结构体值
	if modelValue.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
	}

	if modelValue.Kind() != reflect.Struct || voValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("modelData 和 voType 必须为结构体类型")
	}

	// 遍历 modelData 中的字段
	for i := 0; i < modelValue.NumField(); i++ {
		modelField := modelValue.Field(i)
		fieldType := modelValue.Type().Field(i)

		// 递归处理嵌入的结构体（如 BaseModel）
		if modelField.Kind() == reflect.Struct {
			embeddedModelType := modelField.Type()
			for j := 0; j < embeddedModelType.NumField(); j++ {
				embeddedField := embeddedModelType.Field(j)
				voField := voValue.FieldByName(embeddedField.Name)
				if voField.IsValid() && voField.CanSet() {
					voField.Set(modelField.Field(j))
				}
			}
			continue
		}

		// 查找对应的 VO 字段并赋值
		voField := voValue.FieldByName(fieldType.Name)
		if voField.IsValid() && voField.CanSet() {
			// 赋值：如果类型兼容或是特定类型转换
			if modelField.Type().AssignableTo(voField.Type()) {
				voField.Set(modelField)
			} else if modelField.Kind() == reflect.String && voField.Kind() == reflect.Slice && voField.Type().Elem().Kind() == reflect.Int64 {
				str := modelField.String()
				if str != "" {
					strs := strings.Split(str, ",")
					ints := make([]int64, len(strs))
					for j, s := range strs {
						if val, err := strconv.ParseInt(s, 10, 64); err == nil {
							ints[j] = val
						}
					}
					voField.Set(reflect.ValueOf(ints))
				}
			}
		}
	}

	return voValue.Addr().Interface(), nil
}
