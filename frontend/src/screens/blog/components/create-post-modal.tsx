import { Button } from "@/shared/ui/button";
import { Modal } from "@/shared/ui/modal";
import React, { useContext } from "react";
import { CreatePostModalContext } from "../context/create-post-modal.context";
import { useForm } from "react-hook-form";
import { CreatePostDto, CreatePostForm } from "@/shared/types/blog/blog.dto";
import { PostsApiService } from "@/shared/api/posts/posts.api";
import { useAuth } from "@/shared/hooks/useAuth";
import { usePosts } from "@/shared/hooks/usePosts";

export const CreatePostModal = () => {
  const { auth } = useAuth();
  const { addPost } = usePosts();
  const { register, handleSubmit, reset } = useForm<CreatePostForm>();

  const { visible, setVisible } = useContext(CreatePostModalContext);

  const onCreatePostSubmit = async (data: CreatePostForm) => {
    if (!auth) {
      return;
    }

    const dto: CreatePostDto = {
      ...data,
      authorId: auth?.id,
    };

    const res = await PostsApiService.createPost(dto);
    if (res.status !== 200) {
      console.log(res);
      setVisible(false);
      return;
    }

    reset();
    addPost(res.data);
    setVisible(false);
  };

  return (
    <Modal visible={visible} setVisibleState={setVisible}>
      <h1 className="font-semibold text-xl">Create Post</h1>

      <form className="mt-4" onSubmit={handleSubmit(onCreatePostSubmit)}>
        <div className="mt-6">
          <p className="font-semibold mb-2">Title</p>
          <input
            {...register("title", { required: true })}
            className="w-full p-3 text-base font-semibold bg-transparent focus:outline-none border-2 rounded-md"
            placeholder="Enter title"
          />
        </div>

        <div className="mt-6">
          <p className="font-semibold mb-2">Content</p>
          <textarea
            {...register("body", { required: true })}
            className="w-full h-32 p-3 text-base font-semibold bg-transparent focus:outline-none border-2 rounded-md"
            placeholder="Write post content"
          />
        </div>

        <div className="mt-6">
          <Button type="submit" style={{ width: "100%" }} solid>
            Create
          </Button>
        </div>
      </form>
    </Modal>
  );
};
