## 统一响应格式：

- 正确响应：

```json
{
  "data": T,
  "requestId": "string",
  "timeStamp": number
}
```

- 错误响应：

```json
{
    "code": number,
    "msg": "string",
    "data": T,
    "requestId": "string",
    "timeStamp": number
}
```

# account 账号模块：

## 账户部分：

1. **registerAccount** 注册账号：
   - 请求方式：POST
   - 请求路径：/api/v1/account/registerAccount
   - 请求参数 json：
     - email：string 类型，邮箱
     - nickname：string 类型，昵称
     - password：string 类型，密码
     - phone：string 类型，手机号
     - email_verification_code：string 类型，邮箱验证码
     - img_verification_code：string 类型，图片验证码，大小写不敏感
   - 响应示例：
     ```json
     {
       "data": {
         "nickname": "fender",
         "email": "927171598@qq.com"
       },
       "requestId": "TdGlsTqcsEBbUvhClaRQnAYXVbCRZjjB",
       "timeStamp": 1740052911
     }
     ```
2. **getAccount** 获取账号信息：

   - 请求方式：POST
   - 请求路径：/api/v1/account/getAccount
   - 请求参数 json：
     - email：string 类型，邮箱
   - 响应示例：
     ```json
     {
       "data": {
         "email": "927171598@qq.com",
         "nickname": "fender",
         "phone": "110",
         "avatar": "https://haowallpaper.com/link/common/file/previewFileImg/15942630369381760"
       },
       "requestId": "FRjzgvEFXlsHWPKvvOCdtgAmiOidCwHt",
       "timeStamp": 1740053250
     }
     ```

3. **loginAccount** 登录账号：

   - 请求方式：POST
   - 请求路径：/api/v1/account/loginAccount
   - 请求参数 json：
     - email：string 类型，邮箱
     - password：string 类型，密码
     - img_verification_code：string 类型，图片验证码，大小写不敏感
   - 响应示例：

   ```json
   {
     "data": {
       "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAwNjA3NTUsInVzZXJJZCI6Mn0.Ejv6v1ceDeArV-5zWjEExQUIwm-BfvwwHMRIH6hm3B4",
       "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAyMjYzNTUsInVzZXJJZCI6Mn0.ZpbSXypjzG302IDff6BRLGM92Ieiywz8GimiZViwPr0"
     },
     "requestId": "WgXCIzQcTeEXXLFXKbxxTCrVVlnPsbvI",
     "timeStamp": 1740053555
   }
   ```

4. **logoutAccount** 登出账号：

   - 请求方式：POST
   - 请求路径：/api/v1/account/logoutAccount
   - **请求头** Bearer Token：
   - 响应示例：

   ```json
   {
     "data": "用户注销成功",
     "requestId": "BNqxozHafYcfghrdbzaJyRMgZFiyUgee",
     "timeStamp": 1740063607
   }
   ```

5. **resetPassword** 重置密码：
   - 请求方式：POST
   - 请求路径：/api/v1/account/resetPassword
   - 请求参数 json：
     - email：string 类型，邮箱
     - new_password：string 类型，新密码
     - again_new_password：string 类型，再次输入新密码
     - email_verification_code：string 类型，邮箱验证码
   - 响应示例：
     ```json
     {
       "data": "密码重置成功",
       "requestId": "ZybJbcMxXCMJPhoJnZBABjiQMKTyvJNk",
       "timeStamp": 1740063697
     }
     ```

## post 文章模块

- 统一响应格式：

```json
{
    "data": {
        "id": number,
        "title": "string",
        "image": "string",
        "visibility": "string",
        "content_html": "string",
        "category_ids": []int64
    },
    "requestId": "string",
    "timeStamp": number
}
```

> visibility 只有两种取值："public" 和 "private"，分别表示公开和私密。

1. **GetAllPosts** 获取包含所有文章的列表：

   - 请求方式：GET
   - 请求路径：/api/v1/post/getAllPosts
   - 请求参数 query：
     - pageSize：每页显示的文章数量，默认值：5
     - page：当前页码，默认值：1
   - 响应示例：
     `json
     {
       "data": {
         "currentPage": 2,
         "posts": [
           {
             "id": 6,
             "title": "文章标题6",
             "image": "https://haowallpaper.com/link/common/file/previewFileImg/16019647630462336",
             "visibility": "public",
             "content_markdown": "",
             "content_html": "<h1>Markdown 格式示例</h1>\n\n<h2>1. 标题</h2>\n<h1>这是一级标题</h1>\n<h2>这是二级标题</h2>\n<h3>这是三级标题</h3>\n<h4>这是四��",
             "category_ids": [
               13
             ]
           }
         ],
         "totalPages": 2
       },
       "requestId": "VjDkicQKtuIJGoDUCzwGiAkLgVpxSgvW",
       "timeStamp": 1740042288
     }
     `
     > 注：为了减少传输体积和提供预览效果，此接口对于 content_html 字段只会返回存储在数据库的 HTML 的前 150 个字符。

2. **getOnePost** 获取单篇文章详情：

   - 请求方式：POST
   - 请求路径：/api/v1/post/getOnePost
   - 请求参数 json：
     - id：number 类型，文章 ID
     - title：string 类型，文章标题
       > 注：id 和 title 至少需要有一个，如果两个都传，则以 id 为准。
   - 响应示例：
     `json
     {
       "data": {
       "id": 1,
       "title": "文章标题1",
       "image": "https://haowallpaper.com/link/common/file/previewFileImg/15942630369381760",
       "visibility": "public",
       "content_markdown": "",
       "content_html": "<p>文章内容 1</p>",
       "category_ids": [
         12
       ]
     },
       "requestId": "YWOzpncbNgdQINiDMPcYpwvtaFFQrAPI",
       "timeStamp": 1740043295
     }
     `
     > 注：此接口对于 content_html 字段会返回完整的 HTML 内容。

3. **createOnePost** 创建文章：

   - 请求方式：POST
   - 请求路径：/api/v1/post/createOnePost
   - 请求参数 form-data：
     - title：string 类型，文章标题
     - image：string 类型，文章图片 URL
     - visibility：boolean 类型，文章可见性，取值：0 或 1，也可以 true 或 false，0 表示公开，1 表示私密
     - content_markdown: string 类型，文章内容的 Markdown 格式
     - category_ids：string 类型，文章所属类目 ID 列表，示例：[12]
       > 注：category_ids 字段为 string 类型，必须加上方括号，否则会导致解析失败，此外就是暂时只支持传递一个类目。
   - 响应示例：
     ```json
     {
       "data": {
         "id": 7,
         "title": "文章标题7",
         "image": "https://haowallpaper.com/link/common/file/previewFileImg/16019647630462336",
         "visibility": "public",
         "content_markdown": "",
         "content_html": "<h1>Markdown 格式示例</h1>\n\n<h2>1. 标题</h2>\n<h1>这是一级标题</h1>\n<h2>这是二级标题</h2>\n<h3>这是三级标题</h3>\n<h4>这是四��",
         "category_ids": [13]
       },
       "requestId": "VjDkicQKtuIJGoDUCzwGiAkLgVpxSgvW",
       "timeStamp": 1740042288
     }
     ```

4. **updateOnePost** 更新文章：

   - 请求方式：POST
   - 请求路径：/api/v1/post/updateOnePost
   - 请求参数 json：
     - id：number 类型，文章 ID
     - title：string 类型，文章标题
     - image：string 类型，文章图片 URL
     - visibility：boolean 类型，文章可见性，取值：0 或 1，也可以 true 或 false，0 表示公开，1 表示私密
     - content_markdown: string 类型，文章内容的 Markdown 格式，支持文件路径和直接输入 markdown 文件内容
     - category_ids：string 类型，文章所属类目 ID
       > 除了 id 为必填项外，其他字段都为可选，如果不传则不修改相应字段。
       > 注：category_ids 字段为 string 类型，必须加上方括号，否则会导致解析失败，此外就是暂时只支持传递一个类目。
   - 响应示例：
     ```json
     {
       "data": {
         "id": 1,
         "title": "文章标题1",
         "image": "https://haowallpaper.com/link/common/file/previewFileImg/15942630369381760",
         "visibility": "public",
         "content_markdown": "",
         "content_html": "<p>文章内容 1</p>",
         "category_ids": [12]
       },
       "requestId": "YWOzpncbNgdQINiDMPcYpwvtaFFQrAPI",
       "timeStamp": 1740043295
     }
     ```

5. **deleteOnePost** 删除文章：
   - 请求方式：POST
   - 请求路径：/api/v1/post/deleteOnePost
   - 请求参数 json：
     - id：number 类型，文章 ID
   - 响应示例：
     ```json
     {
       "data": "文章删除成功",
       "requestId": "zWaMCAOkBoYiojZppBSJYZDDNnkCCmrr",
       "timeStamp": 1740048955
     }
     ```

## category 类目模块：

- 统一响应格式：

```json
{
  "data": {
    "id": number,
    "name": "string",
    "description": "string",
    "parent_id": number,
    "path": "string",
    "children": []int64
  },
  "requestId": "string",
  "timeStamp": number
}
```

1. **getOneCategory** 获取单个类目详情：

   - 请求方式：GET
   - 请求路径：/api/v1/category/getOneCategory
   - 请求参数 query：
     - id：number 类型，类目 ID
   - 响应示例：

   ```json
   {
     "data": {
       "id": 1,
       "name": "测试类目1",
       "description": "测试类目1",
       "parent_id": 0,
       "path": "",
       "children": null
     },
     "requestId": "wSdVGCQSbtWQuOdjrzpAjWzLIBPNVIwK",
     "timeStamp": 1740064345
   }
   ```

2. **getCategoryTree** 获取类目树：

   - 请求方式：GET
   - 请求路径：/api/v1/category/getCategoryTree
   - 响应示例：

   ```json
   {
     "data": [
       {
         "id": 12,
         "name": "测试类目5",
         "description": "测试类目5",
         "parent_id": 0,
         "path": "",
         "children": [
           {
             "id": 19,
             "name": "测试类目17",
             "description": "测试类目17",
             "parent_id": 12,
             "path": "/12",
             "children": [
               {
                 "id": 22,
                 "name": "测试类目18",
                 "description": "测试类目18",
                 "parent_id": 19,
                 "path": "/12/19",
                 "children": null
               }
             ]
           }
         ]
       },
       {
         "id": 13,
         "name": "测试类目5",
         "description": "测试类目5",
         "parent_id": 0,
         "path": "",
         "children": null
       }
     ],
     "requestId": "AsFZwohhwIHSjTmQKLevGDyXJObyJXMC",
     "timeStamp": 1740114788
   }
   ```

3. **createOneCategory** 创建类目：

   - 请求方式：POST
   - 请求路径：/api/v1/category/createOneCategory
   - 请求参数 json：
     - name：string 类型，类目名称
     - description：string 类型，类目描述
     - parent_id：number 类型，父类目 ID
   - 响应示例：

   ```json
   {
     "data": {
       "id": 22,
       "name": "测试类目18",
       "description": "测试类目18",
       "parent_id": 19,
       "path": "/12/19",
       "children": null
     },
     "requestId": "JgXNIfiRoIpSuDvTKGUrkpiPhWsJvKCd",
     "timeStamp": 1740114784
   }
   ```

4. **updateOneCategory** 更新类目：

   - 请求方式：POST
   - 请求路径：/api/v1/category/updateOneCategory
   - 请求参数 json：
     - id：number 类型，类目 ID
     - name：string 类型，类目名称
     - description：string 类型，类目描述
     - parent_id：number 类型，父类目 ID，根目录为 0，不传则不修改父类目
   - 响应示例：

   ```json
   {
     "data": {
       "id": 21,
       "name": "测试类目001",
       "description": "测试类目001",
       "parent_id": 0,
       "path": "/0"
     },
     "requestId": "ApUWxYagOuFFhUlvJszyhnDiXyOwemez",
     "timeStamp": 1740115260
   }
   ```

   > 注：更新类目时，如果不传递 parent_id 字段，则表示不修改父类目，如果 parent_id 字段传 0，则表示将类目置于根目录下。

5. **deleteOneCategory** 删除类目：

   - 请求方式：POST
   - 请求路径：/api/v1/category/deleteOneCategory
   - 请求参数 json：
     - id：number 类型，类目 ID
   - 响应示例：

   ```json
   {
     "data": [
       {
         "id": 21,
         "name": "测试类目002",
         "description": "测试类目001",
         "parent_id": 0,
         "path": "/0",
         "children": null
       }
     ],
     "requestId": "yqiGGDEXkeSQnvwWrotwBWZIQOCsgLOY",
     "timeStamp": 1740115579
   }
   ```

6. **getCategoryChildrenTree** 获取类目子树：
   - 请求方式：GET
   - 请求路径：/api/v1/category/getCategoryChildrenTree
   - 请求参数 query：
     - id：number 类型，类目 ID
   - 响应示例：
   ```json
   {
     "data": [
       {
         "id": 19,
         "name": "测试类目17",
         "description": "测试类目17",
         "parent_id": 12,
         "path": "/12",
         "children": [
           {
             "id": 22,
             "name": "测试类目18",
             "description": "测试类目18",
             "parent_id": 19,
             "path": "/12/19",
             "children": null
           }
         ]
       },
       {
         "id": 20,
         "name": "测试类目17",
         "description": "测试类目17",
         "parent_id": 12,
         "path": "/12",
         "children": null
       }
     ],
     "requestId": "sOgOxUNndvRjzVxkCTpSzLmZWuSZOYSd",
     "timeStamp": 1740115733
   }
   ```

## verification 验证码模块：

1. **SendImgVerificationCode** 发送图形验证码：
   - 请求方式：GET
   - 请求路径：/api/v1/verification/sendImgVerificationCode
   - 请求参数 query：
     - email：string 类型，邮箱地址
   - 响应示例：
     ```json
     {
       "data": {
         "imgBase64": "data:image/png;base64,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"
       },
       "requestId": "kXmCrFwABkpNyMjsvhmSrQmyZeXPdhrh",
       "timeStamp": 1740049437
     }
     ```
2. **SendEmailVerificationCode** 验证邮箱验证码：
   - 请求方式：GET
   - 请求路径：/api/v1/verification/sendEmailVerificationCode
   - 请求参数 query：
     - email：string 类型，邮箱地址
   - 响应示例：
     ```json
     {
       "data": "邮箱验证码发送成功, 请注意查收！",
       "requestId": "tPFcJhDOJSHXDSMdULLfRlGDHFMShFYe",
       "timeStamp": 1740049546
     }
     ```

## comment 评论模块：

1. **getOneComment** 获取单条评论：

   - 请求方式：GET
   - 请求路径：/api/v1/comment/getOneComment
   - 请求参数 query：
     - comment_id：string 类型，评论 ID
   - 响应示例：

   ```json
   {
     "data": {
       "id": 1,
       "content": "测试评论 1",
       "user_id": 1,
       "post_id": 5,
       "reply_to_comment_id": 0,
       "replies": null
     },
     "requestId": "OfSQzMSJwOTXTPCkUDqyDWUOcmkrABQd",
     "timeStamp": 1740116840
   }
   ```

2. **getCommentGraph** 获取文章下的评论：

   - 请求方式：GET
   - 请求路径：/api/v1/comment/getCommentGraph
   - 请求参数 query：
     - post_id：string 类型，文章 ID
   - 响应示例：

   ```json
   {
     "data": [
       {
         "id": 2,
         "content": "测试评论 1",
         "user_id": 1,
         "post_id": 5,
         "reply_to_comment_id": 0,
         "replies": []
       },
       {
         "id": 3,
         "content": "测试评论 1",
         "user_id": 1,
         "post_id": 5,
         "reply_to_comment_id": 0,
         "replies": [
           {
             "id": 4,
             "content": "测试评论 1",
             "user_id": 1,
             "post_id": 5,
             "reply_to_comment_id": 3,
             "replies": [
               {
                 "id": 6,
                 "content": "测试评论 1",
                 "user_id": 1,
                 "post_id": 5,
                 "reply_to_comment_id": 4,
                 "replies": []
               }
             ]
           }
         ]
       },
       {
         "id": 5,
         "content": "测试评论 1",
         "user_id": 1,
         "post_id": 5,
         "reply_to_comment_id": 0,
         "replies": []
       }
     ],
     "requestId": "KKpQuPxAvBJEIKQlJRzpgKgRJktusvez",
     "timeStamp": 1740122827
   }
   ```

3. **createOneComment** 创建评论：

   - 请求方式：POST
   - 请求路径：/api/v1/comment/createOneComment
   - 请求参数 json：
     - content：string 类型，评论内容
     - user_id：number 类型，用户 ID
     - post_id：number 类型，文章 ID
     - reply_to_comment_id：number 类型，回复的评论 ID
       > 注：reply_to_comment_id 为 0 时，表示对文章进行评论，reply_to_comment_id 不为 0 时，表示对文章的评论进行回复，默认为 0
   - 响应示例：

   ```json
   {
     "data": {
       "id": 4,
       "content": "测试评论 1",
       "user_id": 1,
       "post_id": 5,
       "reply_to_comment_id": 3,
       "replies": null
     },
     "requestId": "wHIXMZsGpoZSrYjBxUVhzFSEJefAnKAo",
     "timeStamp": 1740117085
   }
   ```

4. **deleteOneComment** 删除评论：
   - 请求方式：POST
   - 请求路径：/api/v1/comment/deleteOneComment
   - 请求参数 json：
     - id：number 类型，评论 ID
   - 响应示例：
   ```json
   {
     "data": {
       "id": 1,
       "content": "测试评论 1",
       "user_id": 1,
       "post_id": 5,
       "reply_to_comment_id": 0,
       "replies": null
     },
     "requestId": "yAHOICsujeqRgXDGlpWExgcIShifmbuR",
     "timeStamp": 1740116847
   }
   ```
