/**
 * 验证码相关常量定义
 */

// ===== 验证码类型 =====
export const VERIFICATION_TYPE = {
  EMAIL: "email",
} as const;

// ===== 验证码用途 =====
export const VERIFICATION_PURPOSE = {
  REGISTER: "register",
  RESET_PASSWORD: "reset_password",
  LOGIN: "login",
} as const;

// ===== 验证码查询键 =====
export const VERIFICATION_QUERY_KEYS = {
  SEND_CODE: "sendVerificationCode",
} as const;

// ===== 类型定义 =====
export type VerificationType =
  (typeof VERIFICATION_TYPE)[keyof typeof VERIFICATION_TYPE];
export type VerificationPurpose =
  (typeof VERIFICATION_PURPOSE)[keyof typeof VERIFICATION_PURPOSE];
