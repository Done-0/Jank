/**
 * 验证相关类型定义
 */

// ===== 请求类型 (Request) =====

// SendEmailCodeRequest 发送邮箱验证码请求
export interface SendEmailCodeRequest {
  email: string; // 邮箱地址
}

// ===== 响应类型 (Response) =====

// SendEmailCodeResponse 发送邮箱验证码响应
export interface SendEmailCodeResponse {
  message: string; // 响应消息
}

// ===== 客户端状态类型 =====

// VerificationState 验证状态（客户端使用）
export interface VerificationState {
  email: string;
  codeSent: boolean;
  countdown: number;
  loading: boolean;
}
