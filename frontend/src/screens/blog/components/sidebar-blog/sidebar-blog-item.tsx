import { PAGES_LINKS } from "@/shared/config/links.config";
import { Blog } from "@/shared/types/blog/blog.type";
import Link from "next/link";
import React, { FC } from "react";

interface SidebarBlogItemProps {
  blog: Blog;
}

export const SidebarBlogItem: FC<SidebarBlogItemProps> = ({ blog }) => {
  const blogLink = PAGES_LINKS.Blog.subs.BlogExcerpt.link(blog.id);

  return (
    <Link href={blogLink} className="rounded-lg bg-gray-100 p-4 cursor-pointer">
      <div>
        <h2 className="font-bold text-sm">{blog.title}</h2>

        <p className="mt-2 text-xs text-gray-500">{blog.body}</p>
      </div>
    </Link>
  );
};
