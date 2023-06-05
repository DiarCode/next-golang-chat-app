import { ButtonHTMLAttributes, PropsWithChildren } from "react";
import { COLORS } from "../config/colors.config";

interface ButtonProps
  extends PropsWithChildren,
    ButtonHTMLAttributes<HTMLButtonElement> {
  solid?: boolean;
}

export const Button = ({ solid = false, children, ...props }: ButtonProps) => {
  const solidStyle = solid
    ? { backgroundColor: COLORS.primary, color: "white" }
    : {};

  const buttonClassName = `px-5 py-2 border-solid border rounded-xl text-sm hover:bg-primary 
      hover:text-white transition-all duration-300 flex items-center justify-center`;

  return (
    <button
      {...props}
      className={buttonClassName}
      style={{
        borderColor: solid ? COLORS.primary : COLORS.light_gray,
        ...solidStyle,
        ...props.style,
      }}
    >
      {children}
    </button>
  );
};
