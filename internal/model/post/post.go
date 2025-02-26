package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"jank.com/jank_blog/internal/model/base"
)

// Post 博客文章模型
type Post struct {
	base.Base
	Title           string           `gorm:"type:varchar(255);not null;index" json:"title"`               // 标题
	Image           string           `gorm:"type:varchar(255)" json:"image"`                              // 图片
	Visibility      bool             `gorm:"type:boolean;not null;default:false;index" json:"visibility"` // 可见性，默认不可见
	ContentMarkdown string           `gorm:"type:text" json:"contentMarkdown"`                            // Markdown 内容
	ContentHTML     string           `gorm:"type:text" json:"contentHtml"`                                // 渲染后的 HTML 内容
	CategoryIDs     CategoryIDsArray `gorm:"type:text" json:"categoryIds"`                                // 分类 ID 数组
}

func (Post) TableName() string {
	return "posts"
}

// CategoryIDsArray 自定义类型
type CategoryIDsArray []int64

// Value 实现 driver.Valuer 接口
func (a CategoryIDsArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}
	return json.Marshal(a)
}

// Scan 实现 sql.Scanner 接口
func (a *CategoryIDsArray) Scan(value interface{}) error {
	if value == nil {
		*a = CategoryIDsArray{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return errors.New("不支持的类型")
	}

	return json.Unmarshal(bytes, a)
}
