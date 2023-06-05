import React, { CSSProperties, FC } from "react";

interface PageTitleProps {
  content: string;
  style?: CSSProperties;
}

export const PageTitle: FC<PageTitleProps> = ({ content, style }) => {
  return (
    <h1 className="font-semibold text-base md:text-2xl" style={style}>
      {content}
    </h1>
  );
};
