# Go 语言编码规范

## 一、命名规范

### 基本原则

- 命名必须以字母（A-Z、a-z）或下划线开头，后续可使用字母、下划线或数字（0-9）。
- 严禁在命名中使用特殊符号，如 @、$、% 等。
- Go 语言区分大小写，首字母大写的标识符可被外部包访问（公开），首字母小写则仅包内可访问（私有）。

### 1. 包命名（package）

- 包名必须与目录名保持一致，应选择简洁、有意义且不与标准库冲突的名称。
- 包名必须全部小写，多个单词可使用下划线分隔或采用混合式小写（不推荐使用驼峰式）。

```go
package demo
package main
```

### 2. 文件命名

- 文件名应有明确含义，简洁易懂。
- 必须使用小写字母，多个单词间使用下划线分隔。

```go
my_test.go
```

### 3. 结构体命名

- 必须采用驼峰命名法，首字母根据访问控制需求决定大小写。
- 结构体声明和初始化必须采用多行格式，示例如下：

```go
// 多行声明
type User struct {
	Username string
	Email    string
}

// 多行初始化
user := User{
	Username: "admin",
	Email:    "admin@example.com",
}
```

### 4. 接口命名

- 必须采用驼峰命名法，首字母根据访问控制需求决定大小写。
- 单一功能的接口名应以 "er" 作为后缀，例如 Reader、Writer。

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

### 5. 变量命名

- 必须采用驼峰命名法，首字母根据访问控制需求决定大小写。
- 特有名词的处理规则：
    - 如果变量为私有且特有名词为首个单词，则使用小写，如 apiClient。
    - 其他情况应保持该名词原有的写法，如 APIClient、repoID、UserID。
    - 错误示例：UrlArray，应写为 urlArray 或 URLArray。
- 布尔类型变量名必须以 Has、Is、Can 或 Allow 开头。

```go
var isExist bool
var hasConflict bool
var canManage bool
var allowGitHook bool
```

### 6. 常量命名

- 常量必须全部使用大写字母，并使用下划线分隔单词。

```go
const APP_VER = "1.0.0"
```

- 枚举类型的常量，应先创建相应类型：

```go
type Scheme string

const (
    HTTP  Scheme = "http"
    HTTPS Scheme = "https"
)
```

### 7. 关键字

Go 语言的关键字：break、case、chan、const、continue、default、defer、else、fallthrough、for、func、go、goto、if、import、interface、map、package、range、return、select、struct、switch、type、var

## 二、注释规范

Go 语言支持 C 风格的注释语法，包括 `/**/` 和 `//`。

- 行注释（//）是最常用的注释形式。
- 块注释（/\* \*/）主要用于包注释，不可嵌套使用，通常用于文档说明或注释大段代码。

### 1. 包注释

- 每个包必须有一个包注释，位于 package 子句之前。
- 包内如果有多个文件，包注释只需在一个文件中出现（建议是与包同名的文件）。
- 包注释必须包含以下信息（按顺序）：
    - 包的基本简介（包名及功能说明）
    - 创建者信息，格式：创建者：[GitHub 用户名]
    - 创建时间，格式：创建时间：yyyy-MM-dd

```go
// Package router  Jank Blog 路由功能定义（路由注册/中间件加载）
// 创建者：Done-0
// 创建时间：2025-03-25
```

### 2. 结构体与接口注释

- 每个自定义结构体或接口必须有注释说明，放在定义的前一行。
- 注释格式为：[结构体名/接口名]，[说明]。
- 结构体的每个成员变量必须有说明，放在成员变量后面并保持对齐。
- 例如：下方的 `User` 为结构体名，`用户对象，定义了用户的基础信息` 为说明。

```go
// User，用户对象，定义了用户的基础信息
type User struct {
    Username  string // 用户名
    Email     string // 邮箱
}
```

### 3. 函数与方法注释

每个函数或方法必须有注释说明，包含以下内容（按顺序）：
  - 简要说明：以函数名开头，使用空格分隔说明部分
  - 参数列表：每行一个参数，参数名开头，“: ”分隔说明部分
  - 返回值：每行一个返回值

```go
// NewtAttrModel 属性数据层操作类的工厂方法
// 参数：
//      ctx: 上下文信息
// 返回值：
//      *AttrModel: 属性操作类指针
func NewAttrModel(ctx *common.Context) *AttrModel {
}
```

### 4. 代码逻辑注释

- 对于关键位置或复杂逻辑处理，必须添加逻辑说明注释。

```go
// 从 Redis 中批量读取属性，对于没有读取到的 id，记录到一个数组里面，准备从 DB 中读取
// 后续代码...
```

### 5. 注释风格

- 统一使用中文注释。
- 中英文字符之间必须使用空格分隔，包括中文与英文、中文与英文标点之间。

```go
// 从 Redis 中批量读取属性，对于没有读取到的 id，记录到一个数组里面，准备从 DB 中读取
```

- 建议全部使用单行注释。
- 单行注释不得超过 120 个字符。

## 三、代码风格

### 1. 缩进与折行

- 缩进必须使用 gofmt 工具格式化（使用 tab 缩进）。
- 每行代码不应超过 120 个字符，超过时应使用换行并保持格式优雅。

> 使用 Goland 开发工具时，可通过快捷键 Control + Alt + L 格式化代码。

### 2. 语句结尾

- Go 语言不需要使用分号结尾，一行代表一条语句。
- 多条语句写在同一行时，必须使用分号分隔。


```go
package main

func main() {
  var a int = 5; var b int = 10
  // 多条语句写在同一行时，必须使用分号分隔
  c := a + b; fmt.Println(c)
}
```

- 代码简单时可以使用多行语句，但建议使用单行语句

```go
package main

func main() {
    var a int = 5
    var b int = 10

    c := a + b
    fmt.Println(c)
}
```

### 3. 括号与空格

- 左大括号不得换行（Go 语法强制要求）。
- 所有运算符与操作数之间必须留有空格。

```go
// 正确示例
if a > 0 {
    // 代码块
}

// 错误示例
if a>0  // a、0 和 > 之间应有空格
{       // 左大括号不可换行，会导致语法错误
    // 代码块
}
```

### 4. import 规范

- 单个包引入时，建议使用括号格式：

```go
import (
    "fmt"
)
```

- 多个包引入时，应按以下顺序分组，并用空行分隔：
    1. 标准库包
    2. 第三方包
    3. 项目内部包

```go
import (
    "context"
    "fmt"
    "sync"
    "time"
    
    "github.com/labstack/echo/v4"
    "golang.org/x/crypto/bcrypt"
    
    "jank.com/jank_blog/internal/global"
    model "jank.com/jank_blog/internal/model/account"
    "jank.com/jank_blog/internal/utils"
    "jank.com/jank_blog/pkg/serve/controller/account/dto"
    "jank.com/jank_blog/pkg/serve/mapper"
    "jank.com/jank_blog/pkg/vo/account"
)
```

- 禁止使用相对路径引入外部包：

```go
// 错误示例
import "../net" // 禁止使用相对路径引入外部包

// 正确示例
import "github.com/repo/proj/src/net"
```

- 包名和导入路径不匹配，建议使用别名：

```go
// 错误示例
import "jank.com/jank_blog/internal/model/account" // 此文件的实际包名为 model

// 正确示例
import model "jank.com/jank_blog/internal/model/account" // 使用 model 别名
```

### 5. 错误处理

- 不得丢弃任何有返回 err 的调用，禁止使用 `_` 丢弃错误，必须全部处理。
- 错误处理原则：
    - 一旦发生错误，应立即返回（尽早 return）。
    - 除非确切了解后果，否则不要使用 panic。
    - 英文错误描述必须全部小写，不需要标点结尾。
    - 必须采用独立的错误流进行处理。

```go
// 错误示例
if err != nil {
    // 错误处理
} else {
    // 正常代码
}

// 正确示例
if err != nil {
    // 错误处理
    return // 或 continue 等
}
// 正常代码
```

### 6. 测试规范

- 测试文件命名必须以 `_test.go` 结尾，如 `example_test.go`。
- 测试函数名称必须以 `Test` 开头，如 `TestExample`。
- 每个重要函数都应编写测试用例，与正式代码一起提交，便于回归测试。

## 四、常用工具

Go 语言提供了多种工具帮助开发者遵循代码规范：

### gofmt

大部分格式问题可通过 gofmt 解决，它能自动格式化代码，确保所有 Go 代码与官方推荐格式保持一致。所有格式相关问题均以 gofmt 结果为准。

### goimports

强烈建议使用 goimports，它在 gofmt 基础上增加了自动删除和引入包的功能。

```bash
go get golang.org/x/tools/cmd/goimports
```

### go vet

vet 工具可静态分析源码中的各种问题，如多余代码、提前 return 的逻辑、struct 的 tag 是否符合标准等。

```bash
go get golang.org/x/tools/cmd/vet
```

使用方法：

```bash
go vet .
```
