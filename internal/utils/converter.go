package utils

import (
    "strconv"
    "strings"
)

// ConvertInt64SliceToString 将 int64 切片转换为逗号分隔的字符串
func ConvertInt64SliceToString(slice []int64) string {
    strSlice := make([]string, len(slice))
    for i, num := range slice {
        strSlice[i] = strconv.FormatInt(num, 10)
    }
    return strings.Join(strSlice, ",")
}
