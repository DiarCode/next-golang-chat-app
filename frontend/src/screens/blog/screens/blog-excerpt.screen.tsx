import { usePost } from "@/shared/hooks/usePosts";
import { useUser } from "@/shared/hooks/useUser";
import { AppLayout } from "@/shared/layouts/app-layout";
import { formatDateFromUnix } from "@/shared/utils/date-formatter";
import { useRouter } from "next/router";
import React from "react";
import { ChevronLeft } from "react-feather";

export const BlogExcerptScreen = () => {
  const router = useRouter();
  const { id } = router.query;
  const { data: post } = usePost(Number(id));

  const blog = post?.data;
  // eslint-disable-next-line react-hooks/rules-of-hooks
  const { data: author } = useUser(blog?.authorId);

  return (
    <AppLayout title={blog?.title || "Loading..."}>
      <button
        onClick={() => router.back()}
        className="flex items-center gap-x-1 cursor-pointer"
      >
        <ChevronLeft size={18} color={"rgb(107 114 128)"} />
        <p className="text-gray-500">Back</p>
      </button>

      <div className="mt-6">
        <h1 className="max-w-2xl text-xl md:text-2xl font-bold">
          {blog?.title}
        </h1>
        <p className="text-sm md:text-base mt-3">{blog?.body}</p>
      </div>

      <div className="flex items-center gap-x-3 mt-6">
        <p className="font-semibold text-xs md:text-sm">
          {author?.data.username || "Unknown"}
        </p>
        <p className="text-xs md:text-sm text-gray-500">
          {formatDateFromUnix(blog?.publishedAt || 1, "DD MMM, YYYY")}
        </p>
      </div>
    </AppLayout>
  );
};
