import { Blog } from "@/shared/types/blog/blog.type";
import React, { FC } from "react";
import { BlogItem } from "./blog-item";

interface BlogsListProps {
  blogs: Blog[];
}

export const BlogsList: FC<BlogsListProps> = ({ blogs }) => {
  console.log(blogs)

  const renderedBlogs = blogs.map((blog, index) => (
    <div key={blog.id}>
      {index > 0 && <hr />}
      <div
        style={{
          paddingTop: index > 0 ? "26px" : "0",
          paddingBottom: 26,
        }}
      >
        <BlogItem blog={blog} />
      </div>
      <hr />
    </div>
  ));
  return <div className="flex flex-col">{renderedBlogs}</div>;
};
