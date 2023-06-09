export const PAGES_LINKS = {
  Home: { name: "Home", link: "/" },
  Blog: {
    name: "Blog",
    link: "/blog",
    subs: {
      BlogExcerpt: {
        name: "Blog Excerpt",
        link: (id: unknown) => `/blog/${id}`,
      },
    },
  },
  Chat: {
    name: "Chat",
    link: "/chat",
    subs: {
      ChatExcerpt: {
        name: "Chat Excerpt",
        link: (id: unknown) => `/chat/${id}`,
      },
    },
  },
  Login: { name: "Login", link: "/login" },
  Signup: { name: "Signup", link: "/signup" },
};

export const NAVBAR_LINKS = {
  Home: PAGES_LINKS.Home,
  Blog: PAGES_LINKS.Blog,
  Chat: PAGES_LINKS.Chat,
};
