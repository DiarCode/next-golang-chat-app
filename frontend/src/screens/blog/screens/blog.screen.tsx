import { AppLayout } from "@/shared/layouts/app-layout";
import React, { useContext, useState } from "react";
import { BlogsList } from "../components/blog-list";
import { blogs } from "@/shared/mocks/blogs";
import { SidebarBlog } from "../components/sidebar-blog/sidebar-blog";
import { Edit2 } from "react-feather";
import { PageTitle } from "@/shared/ui/title";
import { Button } from "@/shared/ui/button";
import { CreatePostModal } from "../components/create-post-modal";
import { CreatePostModalContext } from "../context/create-post-modal.context";
import { usePosts } from "@/shared/hooks/usePosts";

export const BlogScreen = () => {
  const { setVisible } = useContext(CreatePostModalContext);
  const { data } = usePosts();

  return (
    <AppLayout title="Blog">
      <main className="grid grid-cols-3">
        <div className="col-span-3 md:col-span-2 h-full md:pr-7">
          <div className="flex justify-between items-center mb-10">
            <PageTitle content="Blog" />
            <Button style={{ gap: 12 }} onClick={() => setVisible(true)}>
              <Edit2 size={15} />
              <p>Write</p>
            </Button>
          </div>
          <BlogsList blogs={data?.data || []} />
        </div>

        <div className="hidden md:block border-l col-span-1 h-full pl-7">
          <SidebarBlog />
        </div>
      </main>

      <CreatePostModal />
    </AppLayout>
  );
};
