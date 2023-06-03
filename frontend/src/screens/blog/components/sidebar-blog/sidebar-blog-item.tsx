import { Blog } from "@/shared/types/blog/blog.type";
import React, { FC } from "react";

interface SidebarBlogItemProps {
  blog: Blog;
}

export const SidebarBlogItem: FC<SidebarBlogItemProps> = ({ blog }) => {
  return (
    <div className="rounded-lg bg-gray-100 p-4 cursor-pointer">
      <div>
        <h2 className="font-bold text-sm">
          Zero UI: The end of the screen-based interfaces..
        </h2>

        <p className="mt-2 text-xs text-gray-500">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequuntur
          reprehenderit praesentium animi tenetur, deserunt vel sapiente?
        </p>
      </div>
    </div>
  );
};
