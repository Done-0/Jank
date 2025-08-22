import { ReactNode, useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogClose,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";

interface ConfirmDialogProps {
  children: ReactNode;
  title: string;
  description: string;
  onConfirm: () => void;
  confirmText?: string;
  cancelText?: string;
  destructive?: boolean;
}

export function ConfirmDialog({
  children,
  title,
  description,
  onConfirm,
  confirmText = "确认",
  cancelText = "取消",
  destructive = false,
}: ConfirmDialogProps) {
  const [open, setOpen] = useState(false);

  const handleConfirm = () => {
    onConfirm();
    setOpen(false);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger 
        asChild 
        onClick={(e: React.MouseEvent) => e.stopPropagation()}
      >
        {children}
      </DialogTrigger>
      <DialogContent onClick={(e) => e.stopPropagation()}>
        <DialogHeader>
          <DialogTitle>{title}</DialogTitle>
          <DialogDescription className="whitespace-pre-line">
            {description}
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">{cancelText}</Button>
          </DialogClose>
          <Button
            variant={destructive ? "destructive" : "default"}
            onClick={handleConfirm}
          >
            {confirmText}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}