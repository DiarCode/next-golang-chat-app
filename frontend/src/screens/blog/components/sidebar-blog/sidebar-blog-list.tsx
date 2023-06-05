import { Blog } from "@/shared/types/blog/blog.type";
import React, { FC } from "react";
import { SidebarBlogItem } from "./sidebar-blog-item";

interface SidebarBlogListProps {
  blogs: Blog[];
}

export const SidebarBlogList: FC<SidebarBlogListProps> = ({ blogs }) => {
  const renderedBlogs = blogs.map(blog => (
    <SidebarBlogItem blog={blog} key={blog.id} />
  ));
  return (
    <div className="flex flex-col items-start gap-y-3">{renderedBlogs}</div>
  );
};
