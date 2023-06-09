import { PAGES_LINKS } from "@/shared/config/links.config";
import { useUser } from "@/shared/hooks/useUser";
import { Blog } from "@/shared/types/blog/blog.type";
import { formatDateFromUnix } from "@/shared/utils/date-formatter";
import Link from "next/link";
import React, { FC } from "react";

interface BlogItemProps {
  blog: Blog;
}

export const BlogItem: FC<BlogItemProps> = ({ blog }) => {
  const blogLink = PAGES_LINKS.Blog.subs.BlogExcerpt.link(blog.id);
  const { data: author } = useUser(blog.authorId);

  return (
    <div>
      <div>
        <Link href={blogLink}>
          <h2
            className="font-bold text-base md:text-xl cursor-pointer hover:text-primary/90
            transition-all duration-200"
          >
            {blog.title}
          </h2>
        </Link>

        <p className="mt-2 text-xs md:text-sm text-gray-500">{blog.body}</p>
      </div>

      <div className="flex items-center gap-x-3 mt-6">
        <p className="font-semibold text-xs md:text-sm text-primary/80">
          {author?.data.username}
        </p>
        <p className="text-xs md:text-sm text-gray-500">
          {formatDateFromUnix(blog.publishedAt, "DD MMM, YYYY")}
        </p>
      </div>
    </div>
  );
};
