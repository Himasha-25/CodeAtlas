import * as React from "react";
import { cn } from "@/lib/utils";

const variants = {
  default: "bg-black text-white",
  secondary: "bg-gray-100 text-gray-800",
  outline: "border border-gray-300 text-gray-700",
  success: "bg-green-100 text-green-800",
  destructive: "bg-red-100 text-red-800",
  warning: "bg-yellow-100 text-yellow-800",
} as const;

export interface BadgeProps extends React.HTMLAttributes<HTMLSpanElement> {
  variant?: keyof typeof variants;
}

export function Badge({ className, variant = "default", ...props }: BadgeProps) {
  return (
    <span
      className={cn(
        "inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium",
        variants[variant],
        className,
      )}
      {...props}
    />
  );
}
