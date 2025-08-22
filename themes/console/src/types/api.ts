/**
 * API 通用类型定义
 */

// ApiError 统一错误信息结构
export interface ApiError {
  code: string; // 错误码
  message: string; // 错误信息
}

// ApiResponse 统一响应结构
export interface ApiResponse<T = unknown> {
  error?: ApiError; // 错误信息（失败时返回）
  data?: T; // 响应数据
  requestId: string; // 请求 UUID
  timeStamp: number; // 响应时间戳（Unix时间戳）
}
