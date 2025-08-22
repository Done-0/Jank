import { useState } from "react";
import { User, Edit2, Camera } from "lucide-react";
import { toast } from "sonner";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { USER_ROLE_LABELS } from "@/constants";
import { useUpdateProfile } from "@/hooks";
import type { GetProfileResponse } from "@/types";

interface BasicInfoCardProps {
  user: GetProfileResponse;
}

export function BasicInfoCard({ user }: BasicInfoCardProps) {
  // ===== 状态管理 =====
  const [isAvatarDialogOpen, setIsAvatarDialogOpen] = useState(false);
  const [isNicknameDialogOpen, setIsNicknameDialogOpen] = useState(false);
  const [avatarValue, setAvatarValue] = useState(user.avatar || "");
  const [nicknameValue, setNicknameValue] = useState(user.nickname);

  // ===== 使用现有 hooks =====
  const updateProfileMutation = useUpdateProfile({
    onSuccess: () => {
      toast.success("信息更新成功");
      setIsAvatarDialogOpen(false);
      setIsNicknameDialogOpen(false);
    },
    onError: (error: any) => {
      console.error("Profile update error:", error);
      const errorMessage = error?.response?.data?.message || error?.message;
      if (errorMessage) {
        console.error("Error details:", errorMessage);
      }
      toast.error("更新失败，请稍后重试");
    },
  });

  // ===== 事件处理函数 =====
  const handleUpdateAvatar = (e: React.FormEvent) => {
    e.preventDefault();
    updateProfileMutation.mutate({
      nickname: user.nickname,
      avatar: avatarValue,
    });
  };

  const handleUpdateNickname = (e: React.FormEvent) => {
    e.preventDefault();
    if (!nicknameValue.trim()) {
      toast.error("昵称不能为空");
      return;
    }
    updateProfileMutation.mutate({
      nickname: nicknameValue,
      avatar: user.avatar,
    });
  };

  const openAvatarDialog = () => {
    setAvatarValue(user.avatar || "");
    setIsAvatarDialogOpen(true);
  };

  const openNicknameDialog = () => {
    setNicknameValue(user.nickname);
    setIsNicknameDialogOpen(true);
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          <User className="h-5 w-5" />
          基本信息
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-6">
        <div className="flex items-center gap-4">
          {/* 头像 - 悬浮显示编辑提示 */}
          <div className="relative group">
            <Avatar
              className="h-16 w-16 cursor-pointer transition-all duration-200 group-hover:ring-2 group-hover:ring-primary/20"
              onClick={openAvatarDialog}
            >
              {user.avatar && (
                <AvatarImage src={user.avatar} alt={user.nickname} />
              )}
              <AvatarFallback className="text-lg font-semibold">
                {user.nickname.charAt(0).toUpperCase()}
              </AvatarFallback>
            </Avatar>
            {/* 悬浮编辑图标 */}
            <div className="absolute inset-0 bg-black/40 rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-200 flex items-center justify-center cursor-pointer">
              <Camera className="h-5 w-5 text-white" />
            </div>
          </div>

          <div className="flex-1">
            {/* 昵称 - 悬浮显示编辑图标 */}
            <div className="flex items-center gap-2 group">
              <h3 className="font-semibold text-lg">{user.nickname}</h3>
              <Button
                variant="ghost"
                size="sm"
                className="h-6 w-6 p-0 opacity-0 group-hover:opacity-100 transition-opacity duration-200"
                onClick={openNicknameDialog}
              >
                <Edit2 className="h-3 w-3" />
              </Button>
            </div>
            <p className="text-sm text-muted-foreground">{user.email}</p>
            <Badge variant="secondary" className="text-xs mt-2">
              {USER_ROLE_LABELS[
                user.roles[0] as keyof typeof USER_ROLE_LABELS
              ] || user.roles[0]}
            </Badge>
          </div>
        </div>

        {/* 头像编辑对话框 */}
        <Dialog open={isAvatarDialogOpen} onOpenChange={setIsAvatarDialogOpen}>
          <DialogContent className="sm:max-w-md">
            <DialogHeader>
              <DialogTitle>修改头像</DialogTitle>
              <DialogDescription>请输入新的头像链接地址</DialogDescription>
            </DialogHeader>
            <form onSubmit={handleUpdateAvatar} className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="avatar">头像链接</Label>
                <Input
                  id="avatar"
                  placeholder="请输入头像链接"
                  value={avatarValue}
                  onChange={(e) => setAvatarValue(e.target.value)}
                />
              </div>
              <div className="flex flex-col sm:flex-row gap-3">
                <Button
                  type="submit"
                  disabled={updateProfileMutation.isPending}
                  className="flex-1"
                >
                  {updateProfileMutation.isPending ? "保存中..." : "保存"}
                </Button>
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => setIsAvatarDialogOpen(false)}
                  disabled={updateProfileMutation.isPending}
                  className="flex-1"
                >
                  取消
                </Button>
              </div>
            </form>
          </DialogContent>
        </Dialog>

        {/* 昵称编辑对话框 */}
        <Dialog
          open={isNicknameDialogOpen}
          onOpenChange={setIsNicknameDialogOpen}
        >
          <DialogContent className="sm:max-w-md">
            <DialogHeader>
              <DialogTitle>修改昵称</DialogTitle>
              <DialogDescription>请输入新的昵称</DialogDescription>
            </DialogHeader>
            <form onSubmit={handleUpdateNickname} className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="nickname">昵称</Label>
                <Input
                  id="nickname"
                  placeholder="请输入昵称"
                  value={nicknameValue}
                  onChange={(e) => setNicknameValue(e.target.value)}
                />
              </div>
              <div className="flex flex-col sm:flex-row gap-3">
                <Button
                  type="submit"
                  disabled={updateProfileMutation.isPending}
                  className="flex-1"
                >
                  {updateProfileMutation.isPending ? "保存中..." : "保存"}
                </Button>
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => setIsNicknameDialogOpen(false)}
                  disabled={updateProfileMutation.isPending}
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
