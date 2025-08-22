/**
 * 文章编辑器组件
 */

import { useState, useEffect, useRef, useCallback } from "react";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { ArrowLeft, Loader2 } from "lucide-react";
import { useTheme } from "@/components/theme/theme-provider";
import Vditor from "vditor";
import "vditor/dist/index.css";
import type { PostStatus } from "@/constants/post";
import { POST_STATUS } from "@/constants/post";
import type { CreatePostRequest, UpdatePostRequest } from "@/types/post";
import type { CategoryItem } from "@/types/category";

interface PostEditorProps {
  postId?: string;
  postData?: any;
  categories: CategoryItem[];
  isLoading: boolean;
  onSave: (data: CreatePostRequest | UpdatePostRequest) => void;
  onCancel: () => void;
}

export function PostEditor({
  postId,
  postData,
  categories,
  isLoading,
  onSave,
  onCancel,
}: PostEditorProps) {
  // ===== 状态管理 =====
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [description, setDescription] = useState("");
  const [image, setImage] = useState("");
  const [categoryId, setCategoryId] = useState("");
  const [status, setStatus] = useState<PostStatus>(POST_STATUS.DRAFT);
  const [isDirty, setIsDirty] = useState(false);
  const [showSaveDialog, setShowSaveDialog] = useState(false);
  const [showUnsavedDialog, setShowUnsavedDialog] = useState(false);
  const [editorReady, setEditorReady] = useState(false);

  const editorRef = useRef<HTMLDivElement>(null);
  const vditorRef = useRef<Vditor>();
  const initialContentRef = useRef<string>("");
  const { theme } = useTheme();
  const isEditMode = Boolean(postId && postData);

  // ===== 事件处理 =====
  const handleSave = useCallback(() => {
    if (!title.trim()) return toast.error("请输入文章标题");
    setShowSaveDialog(true);
  }, [title]);

  const handleSaveConfirm = useCallback(() => {
    const markdown = vditorRef.current?.getValue() || content;
    const data = isEditMode
      ? ({
          id: postId!,
          title: title.trim(),
          markdown: markdown.trim(),
          description: description.trim(),
          image: image.trim(),
          category_id: categoryId,
          status,
        } as UpdatePostRequest)
      : ({
          title: title.trim(),
          markdown: markdown.trim(),
          description: description.trim(),
          image: image.trim(),
          category_id: categoryId,
          status,
        } as CreatePostRequest);

    onSave(data);
    setShowSaveDialog(false);
    setIsDirty(false);
  }, [
    isEditMode,
    postId,
    title,
    content,
    description,
    image,
    categoryId,
    status,
    onSave,
  ]);

  // ===== 副作用 =====
  // 同步数据
  useEffect(() => {
    if (postData) {
      setTitle(postData.title || "");
      const newContent = postData.markdown || "";
      setContent(newContent);
      initialContentRef.current = newContent;
      setDescription(postData.description || "");
      setImage(postData.image || "");
      setCategoryId(postData.category_id || "");
      setStatus(postData.status || POST_STATUS.DRAFT);
      setIsDirty(false);

      if (vditorRef.current && editorReady) {
        vditorRef.current.setValue(newContent);
      }
    }
  }, [postData, editorReady]);

  // 编辑器初始化
  useEffect(() => {
    if (!editorRef.current) return;

    const timer = setTimeout(() => {
      if (!editorRef.current || vditorRef.current) return;

      vditorRef.current = new Vditor(editorRef.current, {
        height: "100%",
        mode: "ir",
        theme: "classic",
        placeholder: "开始编写你的内容...",
        toolbar: [
          "headings",
          "bold",
          "italic",
          "strike",
          "|",
          "line",
          "quote",
          "list",
          "ordered-list",
          "check",
          "|",
          "code",
          "inline-code",
          "table",
          "|",
          "link",
          "|",
          "edit-mode",
          "|",
          "export",
          "|",
          "undo",
          "redo",
          "|",
          "preview",
          "fullscreen",
        ],
        outline: { enable: true, position: "left" },
        cache: { enable: false },
        after: () => {
          setEditorReady(true);
          if (initialContentRef.current && vditorRef.current) {
            vditorRef.current.setValue(initialContentRef.current);
          }
        },
        input: (value) => {
          setContent(value);
          setIsDirty(true);
        },
      });
    }, 50);

    return () => {
      clearTimeout(timer);
      vditorRef.current?.destroy();
      vditorRef.current = undefined;
      setEditorReady(false);
    };
  }, [theme]);

  // 快捷键
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if ((e.ctrlKey || e.metaKey) && e.key === "s") {
        e.preventDefault();
        handleSave();
      }
    };
    document.addEventListener("keydown", handleKeyDown);
    return () => document.removeEventListener("keydown", handleKeyDown);
  }, [handleSave]);

  // ===== 渲染 =====
  return (
    <div className="flex-1 flex flex-col h-full overflow-hidden w-full min-w-0">
      {/* 顶部导航 */}
      <div className="flex items-center justify-between px-4 py-3 border-b bg-background">
        <div className="flex items-center gap-3 flex-1">
          <Button
            variant="outline"
            onClick={() => (isDirty ? setShowUnsavedDialog(true) : onCancel())}
            className="h-10 px-4 shrink-0 rounded-full"
          >
            <ArrowLeft className="w-4 h-4 mr-2" />
            <span className="hidden sm:inline">返回</span>
          </Button>
          <div className="flex-1 relative flex flex-col items-center">
            <Input
              value={title}
              onChange={(e) => {
                setTitle(e.target.value);
                setIsDirty(true);
              }}
              placeholder="标题"
              className="!text-2xl !font-bold border-none shadow-none focus-visible:ring-0 px-0 bg-transparent placeholder:text-muted-foreground/50 text-center w-full max-w-xl min-w-0 !leading-tight"
            />
            <div
              className="h-px bg-border mt-3"
              style={{
                width: title
                  ? `${Math.min(title.length * 24 + 48, 640)}px`
                  : "140px",
              }}
            ></div>
          </div>
        </div>

        <Button
          onClick={handleSave}
          disabled={isLoading}
          className="h-10 px-4 shrink-0 rounded-full ml-3"
        >
          {isLoading ? <Loader2 className="w-4 h-4 mr-2 animate-spin" /> : null}
          <span className="hidden sm:inline">
            {isLoading ? "保存中" : "保存"}
          </span>
          <span className="sm:hidden">{isLoading ? "保存中" : "保存"}</span>
        </Button>
      </div>

      {/* 编辑器区域 */}
      <div className="flex-1 flex flex-col overflow-hidden w-full min-w-0">
        <div className="flex-1 overflow-y-auto overflow-x-hidden">
          <div className="h-full relative w-full min-w-0">
            <div
              ref={editorRef}
              className="w-full h-full min-w-0 [&_.vditor]:!rounded-lg [&_.vditor]:!border-border"
            />
            {!editorReady && (
              <div className="absolute inset-0 flex items-center justify-center bg-background/90">
                <div className="flex items-center gap-2">
                  <Loader2 className="w-4 h-4 animate-spin" />
                  <span className="text-sm text-muted-foreground">
                    加载编辑器...
                  </span>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>

      {/* 保存对话框 */}
      <Dialog open={showSaveDialog} onOpenChange={setShowSaveDialog}>
        <DialogContent className="max-w-lg mx-4">
          <DialogHeader>
            <DialogTitle>保存文章</DialogTitle>
          </DialogHeader>
          <div className="space-y-4">
            <div>
              <label className="text-sm font-medium">标题</label>
              <Input
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="输入文章标题"
              />
            </div>
            <div>
              <label className="text-sm font-medium">描述</label>
              <Textarea
                value={description}
                onChange={(e) => {
                  setDescription(e.target.value);
                  setIsDirty(true);
                }}
                placeholder="输入文章描述"
                rows={3}
              />
            </div>
            <div>
              <label className="text-sm font-medium">封面图片</label>
              <Input
                value={image}
                onChange={(e) => {
                  setImage(e.target.value);
                  setIsDirty(true);
                }}
                placeholder="输入图片 URL"
                type="url"
              />
            </div>
            <div>
              <label className="text-sm font-medium">分类</label>
              <Select value={categoryId} onValueChange={setCategoryId}>
                <SelectTrigger>
                  <SelectValue placeholder="选择分类" />
                </SelectTrigger>
                <SelectContent>
                  {categories.map((category) => (
                    <SelectItem key={category.id} value={category.id}>
                      {category.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
            <div>
              <label className="text-sm font-medium">状态</label>
              <Select
                value={status}
                onValueChange={(value: PostStatus) => setStatus(value)}
              >
                <SelectTrigger>
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value={POST_STATUS.DRAFT}>草稿</SelectItem>
                  <SelectItem value={POST_STATUS.PUBLISHED}>已发布</SelectItem>
                  <SelectItem value={POST_STATUS.ARCHIVED}>已归档</SelectItem>
                  <SelectItem value={POST_STATUS.PRIVATE}>私有</SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setShowSaveDialog(false)}>
              取消
            </Button>
            <Button onClick={handleSaveConfirm} disabled={isLoading}>
              {isLoading ? (
                <>
                  <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                  保存中...
                </>
              ) : (
                "保存文章"
              )}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 未保存提示 */}
      <Dialog open={showUnsavedDialog} onOpenChange={setShowUnsavedDialog}>
        <DialogContent className="mx-4">
          <DialogHeader>
            <DialogTitle>未保存的更改</DialogTitle>
          </DialogHeader>
          <p className="text-sm text-muted-foreground">
            您有未保存的更改，确定要离开吗？
          </p>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setShowUnsavedDialog(false)}
            >
              继续编辑
            </Button>
            <Button
              variant="destructive"
              onClick={() => {
                setShowUnsavedDialog(false);
                onCancel();
              }}
            >
              丢弃更改
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
