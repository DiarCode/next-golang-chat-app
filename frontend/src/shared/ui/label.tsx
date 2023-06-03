import { FC } from "react";

interface TextProps {
  required?: boolean;
  className?: string;
  text: string;
}

export const Label: FC<TextProps> = ({ className, required = false, text }) => {
  return (
    <div className="flex items-start gap-x-1">
      <p className={`text-sm ${className} font-medium`}>{text}</p>
      {required && <span className="text-red-500 text-base">*</span>}
    </div>
  );
};
