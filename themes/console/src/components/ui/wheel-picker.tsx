/**
 * 滚轮选择器组件 - 拖拽式
 */

import { useEffect, useRef, useState, useCallback } from 'react'

import { cn } from '@/lib/utils'

export interface WheelPickerItem {
  label: string
  value: string | number
}

interface WheelPickerProps {
  items: WheelPickerItem[]
  value?: string | number
  onChange?: (value: string | number) => void
  className?: string
}

export function WheelPicker({
  items,
  value,
  onChange,
  className,
}: WheelPickerProps) {
  const containerRef = useRef<HTMLDivElement>(null)
  const [selectedIndex, setSelectedIndex] = useState(0)
  const [isDragging, setIsDragging] = useState(false)
  const [startY, setStartY] = useState(0)
  const [startScrollTop, setStartScrollTop] = useState(0)

  const itemHeight = 40
  const centerIndex = 2

  // 根据value找到对应的索引
  useEffect(() => {
    if (value !== undefined) {
      const index = items.findIndex(item => item.value === value)
      if (index !== -1 && index !== selectedIndex) {
        setSelectedIndex(index)
        scrollToIndex(index, false)
      }
    }
  }, [value, items, selectedIndex])

  const scrollToIndex = (index: number, smooth = true) => {
    if (containerRef.current) {
      const scrollTop = index * itemHeight
      containerRef.current.scrollTo({
        top: scrollTop,
        behavior: smooth ? 'smooth' : 'auto',
      })
    }
  }

  const updateSelectedIndex = useCallback((scrollTop: number) => {
    const index = Math.round(scrollTop / itemHeight)
    const clampedIndex = Math.max(0, Math.min(index, items.length - 1))

    if (clampedIndex !== selectedIndex) {
      setSelectedIndex(clampedIndex)
      onChange?.(items[clampedIndex]?.value)
    }
  }, [selectedIndex, items, onChange])

  // 鼠标/触摸开始
  const handleStart = (clientY: number) => {
    setIsDragging(true)
    setStartY(clientY)
    setStartScrollTop(containerRef.current?.scrollTop || 0)
  }

  // 鼠标/触摸移动
  const handleMove = useCallback((clientY: number) => {
    if (!isDragging || !containerRef.current) return

    const deltaY = startY - clientY
    const newScrollTop = startScrollTop + deltaY
    containerRef.current.scrollTop = newScrollTop
    updateSelectedIndex(newScrollTop)
  }, [isDragging, startY, startScrollTop, updateSelectedIndex])

  // 鼠标/触摸结束
  const handleEnd = useCallback(() => {
    if (!isDragging) return
    setIsDragging(false)

    // 自动对齐到最近的项目
    if (containerRef.current) {
      const currentScrollTop = containerRef.current.scrollTop
      const targetIndex = Math.round(currentScrollTop / itemHeight)
      const clampedIndex = Math.max(0, Math.min(targetIndex, items.length - 1))
      scrollToIndex(clampedIndex, true)
    }
  }, [isDragging, items.length])

  // 鼠标事件
  const handleMouseDown = (e: React.MouseEvent) => {
    e.preventDefault()
    handleStart(e.clientY)
  }

  // 触摸事件
  const handleTouchStart = (e: React.TouchEvent) => {
    handleStart(e.touches[0].clientY)
  }

  const handleTouchMove = (e: React.TouchEvent) => {
    e.preventDefault()
    handleMove(e.touches[0].clientY)
  }

  const handleTouchEnd = () => {
    handleEnd()
  }

  // 全局鼠标事件监听
  useEffect(() => {
    if (isDragging) {
      const handleGlobalMouseMove = (e: MouseEvent) => {
        handleMove(e.clientY)
      }

      const handleGlobalMouseUp = () => {
        handleEnd()
      }

      document.addEventListener('mousemove', handleGlobalMouseMove)
      document.addEventListener('mouseup', handleGlobalMouseUp)

      return () => {
        document.removeEventListener('mousemove', handleGlobalMouseMove)
        document.removeEventListener('mouseup', handleGlobalMouseUp)
      }
    }
  }, [isDragging, startY, startScrollTop, handleEnd, handleMove])

  // 生成显示项目
  const getVisibleItems = () => {
    const result = []
    for (let i = 0; i < items.length; i++) {
      const distance = Math.abs(i - selectedIndex)
      if (distance <= centerIndex) {
        result.push({
          ...items[i],
          index: i,
          isCenter: i === selectedIndex,
          distance,
        })
      }
    }
    return result
  }

  return (
    <div
      className={cn('relative select-none', className)}
      onMouseDown={handleMouseDown}
      onTouchStart={handleTouchStart}
      onTouchMove={handleTouchMove}
      onTouchEnd={handleTouchEnd}
    >
      <div
        ref={containerRef}
        className="h-52 overflow-hidden"
        style={{
          paddingTop: centerIndex * itemHeight,
          paddingBottom: centerIndex * itemHeight,
        }}
      >
        {items.map((item, index) => (
          <div
            key={`${item.value}-${index}`}
            className="flex items-center justify-center pointer-events-none"
            style={{ height: itemHeight }}
          >
            <span className="text-sm text-transparent select-none">
              {item.label}
            </span>
          </div>
        ))}
      </div>

      {/* 可见项目覆盖层 */}
      <div className="absolute inset-0 pointer-events-none">
        {getVisibleItems().map((item) => (
          <div
            key={item.index}
            className={cn(
              'absolute inset-x-0 flex items-center justify-center',
              'transition-all duration-200 ease-out',
              item.isCenter
                ? 'text-foreground font-semibold text-base'
                : 'text-muted-foreground text-sm',
            )}
            style={{
              top: centerIndex * itemHeight + (item.index - selectedIndex) * itemHeight,
              height: itemHeight,
              opacity: item.isCenter ? 1 : Math.max(0.4, 1 - item.distance * 0.3),
              transform: item.isCenter ? 'scale(1.1)' : 'scale(0.9)',
            }}
          >
            <span>{item.label}</span>
          </div>
        ))}
      </div>

      {/* 中心选中指示器 */}
      <div
        className={cn(
          'absolute inset-x-0 border-y border-primary/30 pointer-events-none',
          'transition-all duration-200',
        )}
        style={{
          top: centerIndex * itemHeight,
          height: itemHeight,
        }}
      />
    </div>
  )
}
