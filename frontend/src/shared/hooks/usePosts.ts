import { useQuery } from "@tanstack/react-query";
import { PostsApiService } from "../api/posts/posts.api";
import { usePostsStore } from "../store/usePostsStore";
import { useLayoutEffect } from "react";

export const usePosts = () => {
  const { posts, initPosts, addPost } = usePostsStore();

  const { data } = useQuery({
    queryKey: ["posts"],
    queryFn: PostsApiService.getAllPosts,
  });

  useLayoutEffect(() => {
    initPosts(data?.data ?? []);
  }, [data?.data, initPosts]);

  return { posts, addPost };
};

export const usePost = (id: number) =>
  useQuery({
    queryKey: ["posts", id],
    queryFn: () => PostsApiService.getPostById(id),
    enabled: Boolean(id),
  });
