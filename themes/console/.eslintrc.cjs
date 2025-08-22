module.exports = {
  root: true,
  env: {
    browser: true, // 浏览器环境
    es2022: true, // 支持最新 ES 特性
    node: true, // Node.js 环境支持
  },
  extends: [
    'eslint:recommended', // ESLint 核心推荐规则
    'plugin:@typescript-eslint/recommended', // TypeScript 基础推荐规则
    'plugin:react/recommended', // React 推荐规则
    'plugin:react-hooks/recommended', // React Hooks 规则
    'plugin:react/jsx-runtime', // 支持新的 JSX 转换（React 17+）
  ],
  ignorePatterns: [
    'dist', // 构建产物
    'build', // 构建目录
    'node_modules', // 依赖包
    '*.config.js', // 配置文件
    '*.config.mjs', // ES 模块配置文件
    '.eslintrc.cjs', // ESLint 配置文件本身
    'public', // 静态资源目录
    'coverage', // 测试覆盖率报告
  ],
  parser: '@typescript-eslint/parser', // 使用 TypeScript 解析器
  parserOptions: {
    ecmaVersion: 'latest', // 支持最新 ECMAScript 版本
    sourceType: 'module', // 使用 ES 模块
    ecmaFeatures: {
      jsx: true, // 支持 JSX
    },
  },
  plugins: [
    '@typescript-eslint', // TypeScript 插件
    'react-refresh', // React 热更新插件
  ],
  settings: {
    react: {
      version: 'detect', // 自动检测 React 版本
    },
  },
  rules: {
    // ==================== TypeScript 规则 - 超宽松 ====================
    '@typescript-eslint/no-unused-vars': 'off', // 完全关闭未使用变量检查
    '@typescript-eslint/no-explicit-any': 'off', // 允许使用 any
    '@typescript-eslint/consistent-type-imports': 'off', // 不强制 type 导入
    '@typescript-eslint/no-non-null-assertion': 'off', // 允许非空断言

    // ==================== React 规则 - 只保留必要的 ====================
    'react/prop-types': 'off', // 关闭 prop-types（使用 TypeScript）
    'react/react-in-jsx-scope': 'off', // 新版 React 不需要导入 React
    'react/display-name': 'off', // 不强制 displayName
    'react/jsx-key': 'error', // 列表项必须有 key 属性（这个很重要）
    'react/jsx-no-undef': 'error', // 禁止使用未定义的 JSX 元素（重要）

    // ==================== React Hooks 规则 ====================
    'react-hooks/rules-of-hooks': 'error', // Hooks 使用规则（重要）
    'react-hooks/exhaustive-deps': 'warn', // 检查 useEffect 依赖项

    // ==================== React Refresh 规则 ====================
    'react-refresh/only-export-components': 'off', // 关闭 React Refresh 限制

    // ==================== 通用规则 - 超宽松 ====================
    'no-console': 'off', // 允许 console
    'no-debugger': 'warn', // 只警告 debugger
    'no-var': 'warn', // 建议不使用 var
    'prefer-const': 'off', // 不强制 const
    'no-multiple-empty-lines': 'off', // 允许多个空行
    'comma-dangle': 'off', // 不强制尾随逗号
    'semi': 'off', // 不强制分号规则
    'quotes': 'off', // 不强制引号类型
    'no-trailing-spaces': 'off', // 允许行尾空格
  },
  overrides: [
    {
      // JavaScript 文件特殊配置
      files: ['*.js', '*.jsx'],
      rules: {
        '@typescript-eslint/no-var-requires': 'off', // JS 文件允许 require
      },
    },
    {
      // 测试文件特殊配置
      files: [
        '*.test.ts',
        '*.test.tsx',
        '*.spec.ts',
        '*.spec.tsx',
        '**/__tests__/**/*',
        '**/__mocks__/**/*',
      ],
      env: {
        jest: true, // Jest 测试环境
      },
      rules: {
        // 测试文件完全放开
      },
    },
    {
      // 配置文件特殊配置
      files: [
        '*.config.ts',
        '*.config.js',
        '*.config.mjs',
        'vite.config.*',
        'tailwind.config.*',
        'postcss.config.*',
      ],
      rules: {
        '@typescript-eslint/no-var-requires': 'off',
      },
    },
  ],
}
