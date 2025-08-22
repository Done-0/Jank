/**
 * 验证相关 Query Hooks
 */

import { useMutation, type UseMutationOptions } from "@tanstack/react-query";
import { verificationService } from "@/services";
import type { SendEmailCodeRequest, SendEmailCodeResponse } from "@/types";

// 发送邮箱验证码
export function useSendEmailCode(
  options?: Omit<
    UseMutationOptions<SendEmailCodeResponse, Error, SendEmailCodeRequest>,
    "mutationFn"
  >
) {
  return useMutation({
    mutationFn: (data: SendEmailCodeRequest) =>
      verificationService.sendEmailCode(data),
    ...options,
  });
}
