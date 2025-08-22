/**
 * 文章编辑器页面
 */

import { useNavigate } from '@tanstack/react-router';
import { toast } from 'sonner';
import { useQueryClient } from '@tanstack/react-query';

import { PostEditor } from "@/components/console/posts/PostEditor";

import { usePost, useCreatePost, useUpdatePost, postKeys } from "@/hooks/use-posts";
import { useAllCategories } from "@/hooks/use-categories";
import { CONSOLE_ROUTES } from "@/constants/routes";
import type { CreatePostRequest, UpdatePostRequest } from "@/types/post";

export function PostEditorPage() {
  // ===== 路由参数 =====
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const urlParams = new URLSearchParams(window.location.search);
  const postId = urlParams.get('id') || undefined;
  const isEditing = !!postId;

  // ===== 数据获取 =====
  const { data: categoriesData } = useAllCategories();
  const categories = categoriesData || [];

  const { data: postData, isLoading: postLoading } = usePost(
    { id: postId! },
    { enabled: isEditing }
  );

  // ===== Mutations =====
  const createPostMutation = useCreatePost({
    onSuccess: () => {
      toast.success('文章创建成功');
      queryClient.invalidateQueries({ queryKey: postKeys.all });
      navigate({ to: CONSOLE_ROUTES.POSTS });
    },
    onError: (error) => {
      toast.error('创建失败: ' + error.message);
    },
  });

  const updatePostMutation = useUpdatePost({
    onSuccess: (_, variables) => {
      toast.success('文章更新成功');
      queryClient.invalidateQueries({ queryKey: postKeys.all });
      queryClient.invalidateQueries({ queryKey: postKeys.detail(variables.id) });
      navigate({ to: CONSOLE_ROUTES.POSTS });
    },
    onError: (error) => {
      toast.error('更新失败: ' + error.message);
    },
  });

  // ===== 事件处理 =====
  const handleSave = (data: CreatePostRequest | UpdatePostRequest) => {
    if (isEditing) {
      updatePostMutation.mutate(data as UpdatePostRequest);
    } else {
      createPostMutation.mutate(data as CreatePostRequest);
    }
  };

  const handleCancel = () => {
    navigate({ to: CONSOLE_ROUTES.POSTS });
  };

  // ===== 渲染 =====
  const isLoading = postLoading || createPostMutation.isPending || updatePostMutation.isPending;

  return (
    <PostEditor
      postId={postId}
      postData={postData}
      categories={categories}
      isLoading={isLoading}
      onSave={handleSave}
      onCancel={handleCancel}
    />
  );
}
