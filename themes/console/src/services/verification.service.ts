/**
 * 验证码服务
 */

import { VERIFICATION_ENDPOINTS } from "@/api";
import { apiClient } from "@/lib/api-client";
import type {
  ApiResponse,
  SendEmailCodeRequest,
  SendEmailCodeResponse,
} from "@/types";

class VerificationService {
  // ===== 验证码管理 =====

  // 发送邮箱验证码
  async sendEmailCode(
    request: SendEmailCodeRequest
  ): Promise<SendEmailCodeResponse> {
    const response = await apiClient.get<ApiResponse<SendEmailCodeResponse>>(
      VERIFICATION_ENDPOINTS.SEND_EMAIL_CODE,
      { params: request }
    );
    return response.data.data!;
  }
}

export const verificationService = new VerificationService();
