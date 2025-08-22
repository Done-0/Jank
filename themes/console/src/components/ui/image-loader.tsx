import React, { useState, useEffect } from 'react';
import { cn } from '@/lib/utils';

interface ImageLoaderProps {
  src?: string | null;
  alt: string;
  aspectRatio?: 'square' | 'video' | 'portrait';
  fallbackIcon?: React.ReactNode;
  className?: string;
}

export const ImageLoader: React.FC<ImageLoaderProps> = ({
  src,
  alt,
  className,
  fallbackIcon,
  aspectRatio = 'video'
}) => {
  const [isLoading, setIsLoading] = useState(!!src);
  const [hasError, setHasError] = useState(false);

  const aspectClasses = {
    square: 'aspect-square',
    video: 'aspect-video',
    portrait: 'aspect-[3/4]'
  };

  // 当 src 变化时重置状态
  useEffect(() => {
    if (src) {
      setIsLoading(true);
      setHasError(false);
    }
  }, [src]);

  // 如果没有src，直接显示fallback
  if (!src) {
    return (
      <div className={cn('relative overflow-hidden', aspectClasses[aspectRatio], className)}>
        <div className="absolute inset-0 flex items-center justify-center bg-muted text-muted-foreground">
          {fallbackIcon}
        </div>
      </div>
    );
  }

  const handleLoad = () => {
    setIsLoading(false);
  };

  const handleError = () => {
    setIsLoading(false);
    setHasError(true);
  };

  return (
    <div className={cn('relative overflow-hidden', aspectClasses[aspectRatio], className)}>
      {/* 波纹加载效果 */}
      {isLoading && <div className="absolute inset-0 skeleton-loading" />}
      
      {/* 图片 */}
      {!hasError && (
        <img
          src={src}
          alt={alt}
          onLoad={handleLoad}
          onError={handleError}
          className={cn(
            'w-full h-full object-cover transition-opacity duration-500 ease-out',
            isLoading ? 'opacity-0' : 'opacity-100'
          )}
          loading="lazy"
        />
      )}
      
      {/* 错误回退 */}
      {hasError && (
        <div className="absolute inset-0 flex items-center justify-center bg-muted text-muted-foreground">
          {fallbackIcon}
        </div>
      )}
    </div>
  );
};
