import { create } from "zustand";
import { Blog } from "../types/blog/blog.type";

interface PostsState {
  posts: Blog[];
  addPost: (post: Blog) => void;
  initPosts: (posts: Blog[]) => void;
}

export const usePostsStore = create<PostsState>()(set => ({
  posts: [],
  addPost: (post: Blog) => set(state => ({ posts: [...state.posts, post] })),
  initPosts: (posts: Blog[]) => set(state => ({ posts: posts })),
}));
