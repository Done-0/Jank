// Package utils 提供参数验证工具
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import "github.com/go-playground/validator/v10"

// ValidErrRes 验证错误结果结构体
type ValidErrRes struct {
	Error bool        // 是否存在错误
	Field string      // 错误字段名
	Tag   string      // 错误标签
	Value interface{} // 错误值
}

// NewValidator 全局验证器实例
var NewValidator = validator.New()

// Validator 参数验证器
// 参数：
//   - data: 待验证的数据
//
// 返回值：
//   - []ValidErrRes: 验证错误结果数组
func Validator(data interface{}) []ValidErrRes {
	var Errors []ValidErrRes
	errs := NewValidator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var el ValidErrRes
			el.Error = true
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Value()

			Errors = append(Errors, el)
		}
	}
	return Errors
}
