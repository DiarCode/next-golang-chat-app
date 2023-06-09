import { CreatePostModalContextProvider } from "@/screens/blog/context/create-post-modal.context";
import { BlogScreen } from "@/screens/blog/screens/blog.screen";

export default function BlogPage() {
  return (
    <CreatePostModalContextProvider>
      <BlogScreen />
    </CreatePostModalContextProvider>
  );
}
