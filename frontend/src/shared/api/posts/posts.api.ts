import axios from "axios";
import { getApiUrl } from "../api";
import { CreatePostDto } from "@/shared/types/blog/blog.dto";
import { Blog } from "@/shared/types/blog/blog.type";

export class PostsApiService {
  public static createPost(dto: CreatePostDto) {
    return axios.post<Blog>(getApiUrl("posts"), dto);
  }

  public static getAllPosts() {
    return axios.get<Blog[]>(getApiUrl("posts"));
  }

  public static getPostById(id: number) {
    return axios.get<Blog>(getApiUrl(`posts/${id}`));
  }
}
