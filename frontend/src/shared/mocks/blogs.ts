import { Blog } from "../types/blog/blog.type";

export const blogs: Blog[] = [
  {
    id: 1,
    title:
      "Zero UI: The end of the screen-based interfaces and what it means for the business",
    body: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequuntur reprehenderit praesentium animi tenetur, deserunt vel sapiente? vIncidunt ipsum dolorum ducimus. Necessitatibus cumque sequi laborum illum natus ipsa consequuntur. Fuga, illum?",
    authorId: 1,
    createdAt: new Date(),
  },

  {
    id: 2,
    title:
      "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequuntur reprehenderit praesentium animi tenetur",
    body: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequuntur reprehenderit praesentium animi tenetur, deserunt vel sapiente? vIncidunt ipsum dolorum ducimus. Necessitatibus cumque sequi laborum illum natus ipsa consequuntur. Fuga, illum?",
    authorId: 2,
    createdAt: new Date(),
  },
];
