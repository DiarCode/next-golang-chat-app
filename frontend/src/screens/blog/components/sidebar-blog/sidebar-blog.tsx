import React from "react";
import { SidebarBlogList } from "./sidebar-blog-list";
import { blogs } from "@/shared/mocks/blogs";

export const SidebarBlog = () => {
  return (
    <>
      <h4 className="font-semibold text-base mb-6">Popular posts</h4>
      <SidebarBlogList blogs={blogs} />
    </>
  );
};
