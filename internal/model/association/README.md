# 跨模块中间表 (Cross-Module Association Tables)

## 简介

本目录包含跨模块之间的关联关系模型，用于表示不同业务实体之间的关系。这些模型通常作为中间表实现多对多关系。

## 命名规范

关联模型的命名遵循以下规则：

1. **表名**: 采用 `主表名_关联表名` 的格式，如 `post_categories`
2. **结构体名**: 采用 `主表实体关联表实体` 的 Pascal 命名法，如 `PostCategory`
3. **字段命名**:
   - 包含两个关联实体的 ID 字段，命名为 `主表名ID` 和 `关联表名ID`
   - 例如: `PostID` 和 `CategoryID`
4. **索引**: 关联字段通常需要创建索引以提高查询性能

## 模型实现

关联模型通常包含以下要素：

- 继承自 `base.Base` 基础模型
- 包含关联实体的外键字段
- 实现 `TableName()` 方法指定表名
- 使用 `gorm` 标签定义字段属性和索引

## 示例

```go
// PostCategory 文章-类目关联模型
type PostCategory struct {
	base.Base
	PostID     int64 `gorm:"type:bigint;not null;index" json:"post_id"`     // 文章ID
	CategoryID int64 `gorm:"type:bigint;not null;index" json:"category_id"` // 类目ID
}

func (PostCategory) TableName() string {
	return "post_categories"
}
```
