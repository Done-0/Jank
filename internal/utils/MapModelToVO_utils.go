package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// MapModelToVO 将模型数据映射到对应的 VO，返回具体类型的指针
func MapModelToVO(modelData interface{}, voPtr interface{}) (interface{}, error) {
	modelVal := reflect.ValueOf(modelData)
	voVal := reflect.ValueOf(voPtr)

	// 检查 voPtr 是否为指向结构体的指针
	if voVal.Kind() != reflect.Ptr || voVal.IsNil() {
		return nil, fmt.Errorf("voPtr 必须为指向结构体的指针")
	}
	voVal = voVal.Elem()

	// 如果 modelData 是指针类型，解引用获取实际的结构体值
	if modelVal.Kind() == reflect.Ptr {
		modelVal = modelVal.Elem()
	}

	// 确保 modelData 和 voPtr 都是结构体类型
	if modelVal.Kind() != reflect.Struct || voVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("modelData 和 voPtr 必须为结构体类型")
	}

	numFields := modelVal.NumField()
	modelType := modelVal.Type()

	// 控制并发的最大数量
	maxConcurrency := 8
	sem := make(chan struct{}, maxConcurrency)

	var wg sync.WaitGroup
	var mu sync.Mutex

	// 遍历 modelData 中的字段并并行处理
	for i := 0; i < numFields; i++ {
		wg.Add(1)
		sem <- struct{}{} // 控制并发

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }() // 释放信号量

			modelField := modelVal.Field(i)
			fieldType := modelType.Field(i)
			voField := voVal.FieldByName(fieldType.Name)

			if voField.IsValid() && voField.CanSet() {
				// 赋值操作
				if modelField.Type().AssignableTo(voField.Type()) {
					mu.Lock()
					voField.Set(modelField)
					mu.Unlock()
				} else if modelField.Kind() == reflect.String && voField.Kind() == reflect.Slice && voField.Type().Elem().Kind() == reflect.Int64 {
					str := modelField.String()
					if str != "" {
						strArray := strings.Split(str, ",")
						intArray := make([]int64, len(strArray))
						for j, s := range strArray {
							if val, err := strconv.ParseInt(s, 10, 64); err == nil {
								intArray[j] = val
							}
						}
						mu.Lock()
						voField.Set(reflect.ValueOf(intArray))
						mu.Unlock()
					}
				}
			}

			// 处理嵌套结构体字段
			if modelField.Kind() == reflect.Struct {
				embeddedModelType := modelField.Type()
				embeddedNumFields := embeddedModelType.NumField()

				for j := 0; j < embeddedNumFields; j++ {
					embeddedField := embeddedModelType.Field(j)
					voField := voVal.FieldByName(embeddedField.Name)

					if voField.IsValid() && voField.CanSet() {
						mu.Lock()
						voField.Set(modelField.Field(j))
						mu.Unlock()
					}
				}
			}
		}(i)
	}

	wg.Wait()

	return voVal.Addr().Interface(), nil
}
