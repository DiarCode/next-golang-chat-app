import {
  CSSProperties,
  InputHTMLAttributes,
  forwardRef,
} from "react";
import { Label } from "./label";

interface InputFieldProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  required?: boolean;
  error?: string;
}

const InputField = forwardRef<HTMLInputElement, InputFieldProps>(
  ({ label, required, error, ...props }, ref) => {
    const isError = Boolean(error);

    return (
      <div>
        {label && <Label text={label} required={required} className="mb-2" />}

        <input
          ref={ref}
          {...props}
          className="w-full min-h-[34px] bg-transparent rounded-lg px-4 py-1 
              border-solid border-2 border-gray-200 placeholder:text-sm"
        />
        {isError && <p className="text-xs text-red-500 mt-[6px]">{error}</p>}
      </div>
    );
  }
);

InputField.displayName = "InputField";
export { InputField };
