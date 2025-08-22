# Jank 主题开发规范

基于主流分层架构和脚本化构建的主题开发规范。

## 🎯 系统架构

```bash
HTTP API → ThemeServiceImpl → ThemeManagerImpl → File System
                     ↓                    ↓
            Business Logic        Core Management
            (Build & Rebuild)     (Switch & Route)
```

**分层设计：**
- **Service 层**：处理业务逻辑，包括构建参数校验和 rebuild 逻辑
- **Manager 层**：纯粹的资源管理，接口统一为 `SwitchTheme(id string)`
- **Utils 层**：通用构建工具，支持脚本化构建流程

## 📁 主题目录结构

主题采用标准化目录结构，支持 **ID 与目录名解耦**：

```
themes/
├── awesome-dark-theme/      # 目录名（可任意命名，如 Git 仓库名）
│   ├── theme.json          # 主题配置文件
│   ├── src/                # 源代码目录
│   ├── dist/               # 构建输出目录
│   └── scripts/            # 构建脚本目录
│       └── build.sh        # 构建脚本
├── moon-theme-v2/          # 版本化目录名
│   └── theme.json          # { "id": "com.company.themes.moon" }
└── default/                # 默认主题
    └── theme.json          # { "id": "com.jank.themes.default" }
```

**重要约定：**
- **主题 ID 与目录名完全解耦**：系统通过扫描目录读取 `theme.json` 获取真实 ID
- **推荐使用域名反转格式 ID**：如 `com.company.themes.theme-name`
- **目录名可任意命名**：支持 Git 仓库名、版本化目录等
- **ID 必须全局唯一**：系统通过 ID 进行主题管理和切换

## 主题配置 (theme.json)

```json
{
  "id": "com.jank.themes.example",
  "name": "示例主题",
  "version": "1.0.0",
  "author": "Your Name",
  "description": "A modern dark theme with beautiful UI components",
  "repository": "https://github.com/username/awesome-dark-theme",
  "preview": "/assets/preview.png",
  "type": "frontend",
  "index_file_path": "/dist/index.html",
  "static_dir_path": "/dist/assets"
}
```

### 核心字段
- `id`: 唯一标识符
- `name`: 显示名称
- `author`: 作者信息
- `description`: 主题描述
- `repository`: 仓库地址
- `preview`: 预览图路径
- `type`: 主题类型 (`frontend` | `console`)
- `index_file_path`: 主页文件路径
- `static_dir_path`: 静态资源目录

## 双主题架构

系统支持两种完全隔离的主题类型：

- **Frontend 主题** (`type: "frontend"`)：用户前端界面，访问路径 `/`
- **Console 主题** (`type: "console"`)：管理后台界面，访问路径 `/console`

## 主题切换 API

```javascript
// 切换主题
await fetch('/api/v1/theme/switch', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ 
    id: 'com.jank.themes.moon',
    theme_type: 'frontend'
  })
});

// 获取当前主题
const response = await fetch('/api/v1/theme/get');

// 列举主题
const themes = await fetch('/api/v1/theme/list?page_no=1&page_size=100');
```

## 权限控制

- Frontend页面只能切换frontend主题
- Console页面只能切换console主题  
- Console可管理所有主题但不影响当前页面类型
- 主题类型完全隔离，确保界面安全

## 开发要点

- 所有主题文件放在 `dist/` 目录
- 静态资源放在 `dist/assets/`
- `theme.json` 路径必须指向 `dist/`
- Frontend主题避免使用 `/console` 开头的路由
- Console主题内部路由不包含 `/console` 前缀
- 使用相对路径引用静态资源
