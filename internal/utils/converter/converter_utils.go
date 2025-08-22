// Package converter 提供类型转换工具函数
// 创建者：Done-0
// 创建时间：2025-08-09
package converter

import (
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

// ToAnyMap 将 Go map 转换为 protobuf Any map
// 参数：
//
//	src: 源 map
//
// 返回值：
//
//	map[string]*anypb.Any: 转换后的 protobuf Any map
//	error: 转换过程中的错误
func ToAnyMap(src map[string]any) (map[string]*anypb.Any, error) {
	if src == nil {
		return nil, nil
	}

	result := make(map[string]*anypb.Any, len(src))
	for k, v := range src {
		pbValue, err := structpb.NewValue(v)
		if err != nil {
			return nil, err
		}

		anyValue, err := anypb.New(pbValue)
		if err != nil {
			return nil, err
		}

		result[k] = anyValue
	}
	return result, nil
}

// FromAnyMap 将 protobuf Any map 转换为 Go map
// 参数：
//
//	src: 源 protobuf Any map
//
// 返回值：
//
//	map[string]any: 转换后的 Go map
//	error: 转换过程中的错误
func FromAnyMap(src map[string]*anypb.Any) (map[string]any, error) {
	if src == nil {
		return nil, nil
	}

	result := make(map[string]any, len(src))
	for k, v := range src {
		if v == nil {
			result[k] = nil
			continue
		}

		var pbValue structpb.Value
		if err := v.UnmarshalTo(&pbValue); err != nil {
			result[k] = v.String()
			continue
		}

		result[k] = pbValue.AsInterface()
	}
	return result, nil
}
