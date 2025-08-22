/**
 * 通用时间选择器组件
 */

import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Calendar, Clock } from "lucide-react";

interface DateTimePickerProps {
  value?: string;
  onChange: (value: string) => void;
  placeholder?: string;
}

export function DateTimePicker({
  value,
  onChange,
  placeholder = "点击选择出生时间",
}: DateTimePickerProps) {
  const [open, setOpen] = useState(false);
  const [selectedYear, setSelectedYear] = useState(2000);
  const [selectedMonth, setSelectedMonth] = useState(1);
  const [selectedDay, setSelectedDay] = useState(1);
  const [selectedHour, setSelectedHour] = useState(12);
  const [selectedMinute, setSelectedMinute] = useState(0);

  // 生成年份选项 (1900-2030)
  const years = Array.from({ length: 131 }, (_, i) => 1900 + i);
  const months = Array.from({ length: 12 }, (_, i) => i + 1);
  const days = Array.from({ length: 31 }, (_, i) => i + 1);
  const hours = Array.from({ length: 24 }, (_, i) => i);
  const minutes = Array.from({ length: 60 }, (_, i) => i);

  const handleConfirm = () => {
    const dateTime = `${selectedYear}-${String(selectedMonth).padStart(
      2,
      "0"
    )}-${String(selectedDay).padStart(2, "0")}T${String(selectedHour).padStart(
      2,
      "0"
    )}:${String(selectedMinute).padStart(2, "0")}`;
    onChange(dateTime);
    setOpen(false);
  };

  const formatDisplayValue = (value: string) => {
    if (!value) return placeholder;
    const date = new Date(value);
    return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(
      2,
      "0"
    )}-${String(date.getDate()).padStart(2, "0")} ${String(
      date.getHours()
    ).padStart(2, "0")}:${String(date.getMinutes()).padStart(2, "0")}`;
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button
          variant="outline"
          className="w-full justify-start text-left font-normal"
        >
          <Calendar className="mr-2 h-4 w-4" />
          {formatDisplayValue(value || "")}
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[500px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle className="flex items-center gap-2">
            <Clock className="h-5 w-5" />
            选择出生时间
          </DialogTitle>
          <p className="text-sm text-muted-foreground">
            请准确选择您的出生年月日和时辰
          </p>
        </DialogHeader>

        <div className="space-y-6">
          {/* 时间选择器 */}
          <div className="grid grid-cols-5 gap-2 sm:gap-4 text-center">
            <div>
              <div className="text-xs sm:text-sm font-medium mb-2">年</div>
              <select
                value={selectedYear}
                onChange={(e) => setSelectedYear(Number(e.target.value))}
                className="w-full p-1 sm:p-2 border rounded-md text-center text-xs sm:text-sm"
              >
                {years.map((year) => (
                  <option key={year} value={year}>
                    {year}年
                  </option>
                ))}
              </select>
            </div>

            <div>
              <div className="text-xs sm:text-sm font-medium mb-2">月</div>
              <select
                value={selectedMonth}
                onChange={(e) => setSelectedMonth(Number(e.target.value))}
                className="w-full p-1 sm:p-2 border rounded-md text-center text-xs sm:text-sm"
              >
                {months.map((month) => (
                  <option key={month} value={month}>
                    {month}月
                  </option>
                ))}
              </select>
            </div>

            <div>
              <div className="text-xs sm:text-sm font-medium mb-2">日</div>
              <select
                value={selectedDay}
                onChange={(e) => setSelectedDay(Number(e.target.value))}
                className="w-full p-1 sm:p-2 border rounded-md text-center text-xs sm:text-sm"
              >
                {days.map((day) => (
                  <option key={day} value={day}>
                    {day}日
                  </option>
                ))}
              </select>
            </div>

            <div>
              <div className="text-xs sm:text-sm font-medium mb-2">时</div>
              <select
                value={selectedHour}
                onChange={(e) => setSelectedHour(Number(e.target.value))}
                className="w-full p-1 sm:p-2 border rounded-md text-center text-xs sm:text-sm"
              >
                {hours.map((hour) => (
                  <option key={hour} value={hour}>
                    {hour}时
                  </option>
                ))}
              </select>
            </div>

            <div>
              <div className="text-sm font-medium mb-2">分</div>
              <select
                value={selectedMinute}
                onChange={(e) => setSelectedMinute(Number(e.target.value))}
                className="w-full p-2 border rounded-md text-center"
              >
                {minutes.map((minute) => (
                  <option key={minute} value={minute}>
                    {String(minute).padStart(2, "0")}分
                  </option>
                ))}
              </select>
            </div>
          </div>

          {/* 按钮 */}
          <div className="flex gap-3 justify-end">
            <Button variant="outline" onClick={() => setOpen(false)}>
              取消
            </Button>
            <Button
              onClick={handleConfirm}
              className="bg-black text-white hover:bg-gray-800"
            >
              确定
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
