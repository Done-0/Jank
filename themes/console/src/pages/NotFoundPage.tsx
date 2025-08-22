import { Home, ArrowLeft } from "lucide-react";

export const NotFoundPage = () => {
  return (
    <div className="min-h-screen bg-background flex items-center justify-center p-4">
      <div className="text-center max-w-lg">
        {/* 404数字  */}
        <div className="mb-8">
          <h1 className="text-9xl font-bold text-muted-foreground/20 mb-4 leading-none select-none">
            404
          </h1>
          <div className="w-16 h-1 bg-primary rounded-full mx-auto mb-6" />
        </div>

        {/* 错误信息 */}
        <div className="mb-8 space-y-3">
          <h2 className="text-2xl font-semibold text-foreground">页面未找到</h2>
          <p className="text-muted-foreground leading-relaxed">
            抱歉，您访问的页面不存在或已被移除
          </p>
        </div>

        {/* 按钮 */}
        <div className="flex gap-3 justify-center flex-wrap">
          <button
            onClick={() => (window.location.href = "/")}
            className="inline-flex items-center gap-2 px-6 py-3 bg-primary text-primary-foreground rounded-full font-medium hover:bg-primary/90 transition-colors"
          >
            <Home className="h-4 w-4" />
            返回首页
          </button>
          <button
            onClick={() => window.history.back()}
            className="inline-flex items-center gap-2 px-6 py-3 bg-transparent text-foreground border border-border rounded-full font-medium hover:bg-accent transition-colors"
          >
            <ArrowLeft className="h-4 w-4" />
            返回上页
          </button>
        </div>
      </div>
    </div>
  );
};
