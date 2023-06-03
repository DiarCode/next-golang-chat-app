import { useQuery } from "@tanstack/react-query";
import { PostsApiService } from "../api/posts/posts.api";

export const usePosts = () =>
  useQuery({ queryKey: ["posts"], queryFn: PostsApiService.getAllPosts });

export const usePost = (id: number) =>
  useQuery({
    queryKey: ["posts", id],
    queryFn: () => PostsApiService.getPostById(id),
    enabled: Boolean(id),
  });
