/**
 * Checkbox component
 */
import * as React from "react";
import { cn } from "@/lib/utils";
import { CheckIcon } from "lucide-react";

export interface CheckboxProps
  extends Omit<React.InputHTMLAttributes<HTMLInputElement>, "type"> {
  onCheckedChange?: (checked: boolean) => void;
}

const Checkbox = React.forwardRef<HTMLInputElement, CheckboxProps>(
  ({ className, onCheckedChange, onChange, ...props }, ref) => {
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      onCheckedChange?.(e.target.checked);
      onChange?.(e);
    };

    return (
      <div className="relative inline-flex items-center">
        <input
          type="checkbox"
          className="sr-only"
          ref={ref}
          onChange={handleChange}
          {...props}
        />
        <div
          className={cn(
            "peer h-4 w-4 shrink-0 rounded-sm border border-primary ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
            "data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground",
            props.checked && "bg-primary text-primary-foreground",
            className
          )}
        >
          {props.checked && (
            <CheckIcon className="h-4 w-4 text-current" />
          )}
        </div>
      </div>
    );
  }
);

Checkbox.displayName = "Checkbox";

export { Checkbox };