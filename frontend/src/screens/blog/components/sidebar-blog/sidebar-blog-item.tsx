import { PAGES_LINKS } from "@/shared/config/links.config";
import { Blog } from "@/shared/types/blog/blog.type";
import Link from "next/link";
import React, { FC } from "react";

interface SidebarBlogItemProps {
  blog: Blog;
}

const MAX_DESCRIPTION_SIZE = 150;

export const SidebarBlogItem: FC<SidebarBlogItemProps> = ({ blog }) => {
  const blogLink = PAGES_LINKS.Blog.subs.BlogExcerpt.link(blog.id);
  const slicedDescription = blog.body.slice(0, MAX_DESCRIPTION_SIZE);

  return (
    <Link href={blogLink} className="rounded-lg bg-gray-100 p-4 cursor-pointer">
      <div>
        <h2 className="font-bold text-sm">{blog.title}</h2>

        <p className="mt-2 text-xs text-gray-500">{slicedDescription}...</p>
      </div>
    </Link>
  );
};
