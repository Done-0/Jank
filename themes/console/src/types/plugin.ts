/**
 * 插件相关类型定义
 */

// ===== 请求类型 (Request) =====

// RegisterPluginRequest 注册插件请求
export interface RegisterPluginRequest {
  id: string; // 插件 ID
  rebuild?: boolean; // 强制重新编译
}

// UnregisterPluginRequest 注销插件请求
export interface UnregisterPluginRequest {
  id: string; // 插件 ID
}

// GetPluginRequest 获取插件信息请求
export interface GetPluginRequest {
  id: string; // 插件 ID
}

// ListPluginsRequest 插件列表请求
export interface ListPluginsRequest {
  status?: string; // 插件状态筛选
  page_no: number; // 页码（int64），从1开始
  page_size: number; // 每页数量（int64），最大100
}

// ExecutePluginRequest 执行插件请求
export interface ExecutePluginRequest {
  id: string; // 插件 ID
  method: string; // 方法名
  args?: Record<string, any>; // 方法参数
}

// ===== 响应类型 (Response) =====

// RegisterPluginResponse 注册插件响应
export interface RegisterPluginResponse {
  message: string; // 注册结果消息
}

// UnregisterPluginResponse 注销插件响应
export interface UnregisterPluginResponse {
  message: string; // 注销结果消息
}

// GetPluginResponse 插件信息
export interface GetPluginResponse {
  // 基本信息
  id: string; // 插件唯一标识
  name: string; // 插件显示名称
  version: string; // 插件版本号
  author: string; // 插件作者
  description?: string; // 插件描述
  repository: string; // 插件仓库地址
  binary: string; // 二进制文件路径
  type: string; // 插件类型（provider/filter/handler/notifier）

  // 配置信息
  auto_start: boolean; // 是否自动启动
  start_timeout?: number; // 启动超时时间（int64 毫秒）
  min_port?: number; // 最小端口号（uint）
  max_port?: number; // 最大端口号（uint）
  auto_mtls?: boolean; // 是否启用自动 MTLS
  managed?: boolean; // 是否为托管模式

  // 运行时信息
  status: string; // 当前状态
  started_at?: number; // 启动时间戳（int64 Unix时间戳）
  process_id?: string; // 进程标识
  protocol?: string; // 通信协议
  is_exited?: boolean; // 是否已退出
  negotiated_version?: number; // 协商的协议版本（int）
  process_pid?: number; // 系统进程 PID（int）
  protocol_version?: number; // 协议版本（int）
  network_addr?: string; // 网络地址
}

// ListPluginsResponse 插件列表响应
export interface ListPluginsResponse {
  total: number; // 总数（int64）
  page_no: number; // 当前页码（int64）
  page_size: number; // 每页大小（int64）
  list: GetPluginResponse[]; // 插件列表
}

// ExecutePluginResponse 执行插件响应
export interface ExecutePluginResponse {
  method: string; // 执行的方法名
  data: Record<string, any>; // 插件返回的业务数据
}

// StartPluginResponse 启动插件响应
export interface StartPluginResponse {
  message: string; // 启动结果消息
}

// StopPluginResponse 停止插件响应
export interface StopPluginResponse {
  message: string; // 停止结果消息
}

// ===== 客户端状态类型 =====

// PluginState 插件状态（客户端使用）
export interface PluginState {
  plugins: GetPluginResponse[];
  selectedPlugin: GetPluginResponse | null;
  loading: boolean;
}
