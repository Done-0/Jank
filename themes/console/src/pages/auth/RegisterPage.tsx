/**
 * 用户注册页面
 */
import { useState } from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";
import { useMutation } from "@tanstack/react-query";
import { Link, useNavigate } from "@tanstack/react-router";

import {
  Button,
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  Input,
} from "@/components/ui";
import { AUTH_ROUTES } from "@/constants";
import { userService, verificationService } from "@/services";
import type { RegisterRequest } from "@/types";

// ===== 表单验证 =====
const schema = z.object({
  email: z.string().email("请输入有效邮箱"),
  nickname: z.string().min(2, "昵称至少2个字符").max(20, "昵称最多20个字符"),
  password: z.string().min(6, "密码至少6位").max(20, "密码最多20位"),
  email_verification_code: z.string().length(6, "验证码必须为6位"),
});

type FormData = z.infer<typeof schema>;

export function RegisterPage() {
  // ===== Hooks =====
  const navigate = useNavigate();
  const [countdown, setCountdown] = useState(0);
  const {
    register,
    handleSubmit,
    formState: { errors },
    watch,
  } = useForm<FormData>({
    resolver: zodResolver(schema),
  });

  // ===== 状态 =====
  const emailValue = watch("email");

  // ===== API 调用 =====
  const sendCodeMutation = useMutation({
    mutationFn: (email: string) => verificationService.sendEmailCode({ email }),
    onSuccess: () => {
      toast.success("验证码已发送，请查收邮件");
      setCountdown(60);
      const timer = setInterval(() => {
        setCountdown((prev) => {
          if (prev <= 1) {
            clearInterval(timer);
            return 0;
          }
          return prev - 1;
        });
      }, 1000);
    },
    onError: (error: any) => {
      console.error("Send verification code failed:", error);
      toast.error("发送验证码失败");
    },
  });

  const registerMutation = useMutation({
    mutationFn: (data: RegisterRequest) => userService.register(data),
    onSuccess: (data) => {
      toast.success(`注册成功，欢迎 ${data.nickname}！请登录`);
      navigate({ to: AUTH_ROUTES.LOGIN });
    },
    onError: (error: any) => {
      console.error("Register failed:", error);
      toast.error("注册失败");
    },
  });

  // ===== 事件处理 =====
  const onSubmit = (data: FormData) => registerMutation.mutate(data);
  const handleSendCode = () => {
    if (!emailValue) return toast.error("请先输入邮箱");
    sendCodeMutation.mutate(emailValue);
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <Card className="w-full max-w-sm">
        <CardHeader className="text-center space-y-2">
          <CardTitle className="text-2xl">注册</CardTitle>
          <p className="text-sm text-muted-foreground">创建您的新账户</p>
        </CardHeader>

        <CardContent>
          <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
            <div className="space-y-2">
              <label htmlFor="email" className="text-sm font-medium">
                邮箱
              </label>
              <Input
                id="email"
                type="email"
                placeholder="your@email.com"
                {...register("email")}
              />
              {errors.email && (
                <p className="text-sm text-destructive">
                  {errors.email.message}
                </p>
              )}
            </div>

            <div className="space-y-2">
              <label htmlFor="nickname" className="text-sm font-medium">
                昵称
              </label>
              <Input
                id="nickname"
                type="text"
                placeholder="请输入昵称"
                {...register("nickname")}
              />
              {errors.nickname && (
                <p className="text-sm text-destructive">
                  {errors.nickname.message}
                </p>
              )}
            </div>

            <div className="space-y-2">
              <label htmlFor="password" className="text-sm font-medium">
                密码
              </label>
              <Input
                id="password"
                type="password"
                placeholder="请输入密码"
                {...register("password")}
              />
              {errors.password && (
                <p className="text-sm text-destructive">
                  {errors.password.message}
                </p>
              )}
            </div>

            <div className="space-y-2">
              <label
                htmlFor="verification-code"
                className="text-sm font-medium"
              >
                邮箱验证码
              </label>
              <div className="flex gap-2">
                <Input
                  id="verification-code"
                  type="text"
                  placeholder="请输入验证码"
                  maxLength={6}
                  className="flex-1"
                  {...register("email_verification_code")}
                />
                <Button
                  type="button"
                  variant="outline"
                  size="sm"
                  onClick={handleSendCode}
                  disabled={countdown > 0 || sendCodeMutation.isPending}
                  className="shrink-0"
                >
                  {sendCodeMutation.isPending
                    ? "获取中..."
                    : countdown > 0
                    ? `${countdown}s`
                    : "获取验证码"}
                </Button>
              </div>
              {errors.email_verification_code && (
                <p className="text-sm text-destructive">
                  {errors.email_verification_code.message}
                </p>
              )}
            </div>

            <Button
              type="submit"
              className="w-full"
              disabled={registerMutation.isPending}
            >
              {registerMutation.isPending ? "注册中..." : "注册"}
            </Button>
          </form>

          <div className="mt-4 text-center text-sm text-muted-foreground">
            已有账户？{" "}
            <Link
              to={AUTH_ROUTES.LOGIN}
              className="text-primary hover:underline"
            >
              立即登录
            </Link>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
