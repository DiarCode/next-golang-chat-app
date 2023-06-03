import { PAGES_LINKS } from "@/shared/config/links.config";
import { Blog } from "@/shared/types/blog/blog.type";
import Link from "next/link";
import React, { FC } from "react";

interface BlogItemProps {
  blog: Blog;
}

export const BlogItem: FC<BlogItemProps> = ({ blog }) => {
  const blogLink = PAGES_LINKS.Blog.subs.BlogExcerpt.link(blog.id);

  return (
    <div>
      <div>
        <Link href={blogLink}>
          <h2
            className="font-bold text-base md:text-xl cursor-pointer hover:text-primary/90
            transition-all duration-200"
          >
            Zero UI: The end of the screen-based interfaces and what it means
            forf the screen-based interfaces and what it means for the business
          </h2>
        </Link>

        <p className="mt-2 text-xs md:text-sm text-gray-500">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequuntur
          reprehenderit praesentium animi tenetur, deserunt vel sapiente?
          Incidunt ipsum dolorum ducimus. Necessitatibus cumque sequi laborum
          illum natus ipsa consequuntur. Fuga, illum?
        </p>
      </div>

      <div className="flex items-center gap-x-3 mt-6">
        <p className="font-semibold text-xs md:text-sm text-primary/80">Chrisitan Merffy</p>
        <p className="text-xs md:text-sm text-gray-500">Apr 16, 2023</p>
      </div>
    </div>
  );
};
