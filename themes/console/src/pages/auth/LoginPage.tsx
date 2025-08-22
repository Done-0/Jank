/**
 * 用户登录页面
 */
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
import { CONSOLE_ROUTES } from "@/constants/routes";
import { userService } from "@/services";
import type { LoginRequest } from "@/types";

// ===== 表单验证 =====
const schema = z.object({
  email: z.string().email("请输入有效邮箱"),
  password: z.string().min(6, "密码至少6位"),
});

type FormData = z.infer<typeof schema>;

export function LoginPage() {
  // ===== Hooks =====
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(schema),
  });

  // ===== API 调用 =====
  const loginMutation = useMutation({
    mutationFn: (data: LoginRequest) => userService.login(data),
    onSuccess: () => {
      toast.success("登录成功");
      navigate({ to: CONSOLE_ROUTES.ROOT });
    },
    onError: (error: any) => {
      console.error("Login failed:", error);
      toast.error("登录失败");
    },
  });

  // ===== 事件处理 =====
  const onSubmit = (data: FormData) => loginMutation.mutate(data);

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <Card className="w-full max-w-sm">
        <CardHeader className="text-center space-y-2">
          <CardTitle className="text-2xl">登录</CardTitle>
          <p className="text-sm text-muted-foreground">使用邮箱和密码登录</p>
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

            <Button
              type="submit"
              className="w-full"
              disabled={loginMutation.isPending}
            >
              {loginMutation.isPending ? "登录中..." : "登录"}
            </Button>
          </form>

          <div className="mt-4 text-center text-sm text-muted-foreground">
            还没有账户？{" "}
            <Link
              to={AUTH_ROUTES.REGISTER}
              className="text-primary hover:underline"
            >
              立即注册
            </Link>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
