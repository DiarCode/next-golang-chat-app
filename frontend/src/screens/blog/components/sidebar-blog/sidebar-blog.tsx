import React from "react";
import { SidebarBlogList } from "./sidebar-blog-list";
import { usePosts } from "@/shared/hooks/usePosts";

export const SidebarBlog = () => {
  const { posts } = usePosts();

  return (
    <>
      <h4 className="font-semibold text-base mb-6">Popular posts</h4>
      <SidebarBlogList blogs={posts} />
    </>
  );
};
