import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { toast } from "sonner";
import { Eye, EyeOff, Loader2, Shield } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useUserStore } from "@/stores";
import { useResetPassword } from "@/hooks/use-user";
import { useSendEmailCode } from "@/hooks/use-verification";

// 修改密码表单验证
const resetPasswordSchema = z
  .object({
    old_password: z.string().min(6, "原密码至少6位").max(20, "原密码最多20位"),
    new_password: z.string().min(6, "新密码至少6位").max(20, "新密码最多20位"),
    confirm_password: z.string(),
    email_verification_code: z.string().length(6, "验证码必须为6位"),
  })
  .refine((data) => data.new_password === data.confirm_password, {
    message: "两次输入的密码不一致",
    path: ["confirm_password"],
  })
  .refine((data) => data.old_password !== data.new_password, {
    message: "新密码不能与原密码相同",
    path: ["new_password"],
  });

type ResetPasswordForm = z.infer<typeof resetPasswordSchema>;

export function SecurityCard() {
  // ===== 状态管理 =====
  const { user } = useUserStore();
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [countdown, setCountdown] = useState(0);

  // ===== 表单管理 =====
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<ResetPasswordForm>({
    resolver: zodResolver(resetPasswordSchema),
    defaultValues: {
      old_password: "",
      new_password: "",
      confirm_password: "",
      email_verification_code: "",
    },
  });

  const userEmail = user?.email || "";

  // ===== Mutations =====
  const resetPasswordMutation = useResetPassword({
    onSuccess: () => {
      toast.success("密码修改成功");
      setIsDialogOpen(false);
      reset();
    },
    onError: (error: any) => {
      console.error("Reset password error:", error);
      const errorMessage = error?.response?.data?.message || error?.message;
      if (errorMessage) {
        console.error("Error details:", errorMessage);
      }
      toast.error("密码修改失败，请稍后重试");
    },
  });

  const sendEmailCodeMutation = useSendEmailCode({
    onSuccess: () => {
      toast.success("验证码已发送，请查收邮件");
      startCountdown();
    },
    onError: (error: any) => {
      console.error("Send email code error:", error);
      const errorMessage = error?.response?.data?.message || error?.message;
      if (errorMessage) {
        console.error("Error details:", errorMessage);
      }
      toast.error("验证码发送失败，请稍后重试");
    },
  });

  // ===== 事件处理函数 =====
  const startCountdown = () => {
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
  };

  const handleSendVerificationCode = () => {
    if (!userEmail) {
      toast.error("用户邮箱信息缺失");
      return;
    }
    sendEmailCodeMutation.mutate({ email: userEmail });
  };

  const handleResetPassword = (data: ResetPasswordForm) => {
    resetPasswordMutation.mutate({
      old_password: data.old_password,
      new_password: data.new_password,
      email_verification_code: data.email_verification_code,
    });
  };

  const openPasswordDialog = () => {
    reset({
      old_password: "",
      new_password: "",
      confirm_password: "",
      email_verification_code: "",
    });
    setIsDialogOpen(true);
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          <Shield className="h-5 w-5" />
          安全设置
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="flex items-center justify-between">
          <div>
            <p className="font-medium">登录密码</p>
            <p className="text-sm text-muted-foreground">
              定期更换密码可以提高账户安全性
            </p>
          </div>
          <Button onClick={openPasswordDialog}>修改密码</Button>
        </div>

        <div className="flex items-center justify-between">
          <div>
            <p className="font-medium">邮箱验证</p>
            <p className="text-sm text-muted-foreground">
              当前邮箱：{user?.email || "未设置"}
            </p>
          </div>
          <Button variant="outline" disabled>
            已验证
          </Button>
        </div>

        <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
          <DialogContent className="sm:max-w-md">
            <DialogHeader>
              <DialogTitle>修改密码</DialogTitle>
              <DialogDescription>
                请输入原密码、邮箱验证码和新密码来修改您的登录密码
              </DialogDescription>
            </DialogHeader>
            <form
              onSubmit={handleSubmit(handleResetPassword)}
              className="space-y-6"
            >
              <div className="space-y-2">
                <Label htmlFor="old_password">原密码</Label>
                <div className="relative">
                  <Input
                    id="old_password"
                    type={showPassword ? "text" : "password"}
                    placeholder="请输入原密码"
                    {...register("old_password")}
                  />
                  <Button
                    type="button"
                    variant="ghost"
                    size="sm"
                    className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"
                    onClick={() => setShowPassword(!showPassword)}
                  >
                    {showPassword ? (
                      <EyeOff className="h-4 w-4" />
                    ) : (
                      <Eye className="h-4 w-4" />
                    )}
                  </Button>
                </div>
                {errors.old_password && (
                  <p className="text-sm text-destructive">
                    {errors.old_password.message}
                  </p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="email_verification_code">邮箱验证码</Label>
                <div className="flex gap-2">
                  <Input
                    id="email_verification_code"
                    placeholder="请输入6位验证码"
                    {...register("email_verification_code")}
                  />
                  <Button
                    type="button"
                    variant="outline"
                    onClick={handleSendVerificationCode}
                    disabled={
                      countdown > 0 ||
                      sendEmailCodeMutation.isPending ||
                      !userEmail
                    }
                    className="whitespace-nowrap"
                  >
                    {sendEmailCodeMutation.isPending
                      ? "发送中..."
                      : countdown > 0
                      ? `${countdown}s`
                      : "发送验证码"}
                  </Button>
                </div>
                {errors.email_verification_code && (
                  <p className="text-sm text-destructive">
                    {errors.email_verification_code.message}
                  </p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="new_password">新密码</Label>
                <div className="relative">
                  <Input
                    id="new_password"
                    type={showPassword ? "text" : "password"}
                    placeholder="请输入新密码"
                    {...register("new_password")}
                  />
                  <Button
                    type="button"
                    variant="ghost"
                    size="sm"
                    className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"
                    onClick={() => setShowPassword(!showPassword)}
                  >
                    {showPassword ? (
                      <EyeOff className="h-4 w-4" />
                    ) : (
                      <Eye className="h-4 w-4" />
                    )}
                  </Button>
                </div>
                {errors.new_password && (
                  <p className="text-sm text-destructive">
                    {errors.new_password.message}
                  </p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="confirm_password">确认新密码</Label>
                <div className="relative">
                  <Input
                    id="confirm_password"
                    type={showConfirmPassword ? "text" : "password"}
                    placeholder="请再次输入新密码"
                    {...register("confirm_password")}
                  />
                  <Button
                    type="button"
                    variant="ghost"
                    size="sm"
                    className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"
                    onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                  >
                    {showConfirmPassword ? (
                      <EyeOff className="h-4 w-4" />
                    ) : (
                      <Eye className="h-4 w-4" />
                    )}
                  </Button>
                </div>
                {errors.confirm_password && (
                  <p className="text-sm text-destructive">
                    {errors.confirm_password.message}
                  </p>
                )}
              </div>

              <div className="flex flex-col sm:flex-row gap-3">
                <Button
                  type="submit"
                  disabled={resetPasswordMutation.isPending}
                  className="flex-1"
                >
                  {resetPasswordMutation.isPending && (
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  )}
                  {resetPasswordMutation.isPending ? "修改中..." : "确认修改"}
                </Button>
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => setIsDialogOpen(false)}
                  disabled={resetPasswordMutation.isPending}
                  className="flex-1"
                >
                  取消
                </Button>
              </div>
            </form>
          </DialogContent>
        </Dialog>
      </CardContent>
    </Card>
  );
}
