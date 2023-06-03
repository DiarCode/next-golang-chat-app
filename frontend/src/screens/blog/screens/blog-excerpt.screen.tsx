import { usePost } from "@/shared/hooks/usePosts";
import { AppLayout } from "@/shared/layouts/app-layout";
import { blogs } from "@/shared/mocks/blogs";
import { useRouter } from "next/router";
import React from "react";
import { ChevronLeft } from "react-feather";

export const BlogExcerptScreen = () => {
  const router = useRouter();
  const { id } = router.query;

  const { data } = usePost(Number(id));

  // TODO: fetch post by id and get title
  const blog = data?.data ?? blogs[0];

  const title = blog.title;

  return (
    <AppLayout title={title}>
      <button
        onClick={() => router.back()}
        className="flex items-center gap-x-1 cursor-pointer"
      >
        <ChevronLeft size={18} color={"rgb(107 114 128)"} />
        <p className="text-gray-500">Back</p>
      </button>

      <div className="mt-6">
        <h1 className="max-w-2xl text-xl md:text-2xl font-bold">
          {blog.title}
        </h1>
        <p className="text-sm md:text-base mt-3">{blog.body}</p>
      </div>

      <div className="flex items-center gap-x-3 mt-6">
        <p className="font-semibold text-xs md:text-sm">Chrisitan Merffy</p>
        <p className="text-xs md:text-sm text-gray-500">Apr 16, 2023</p>
      </div>
    </AppLayout>
  );
};
