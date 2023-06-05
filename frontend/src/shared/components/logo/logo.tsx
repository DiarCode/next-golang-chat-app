import { PAGES_LINKS } from "@/shared/config/links.config";
import Link from "next/link";

export const Logo = () => {
  return (
    <Link href={PAGES_LINKS.Home.link}>
      <p
        className={`font-bold text-lg text-transparent bg-clip-text bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500`}
      >
        Meowchat
      </p>
    </Link>
  );
};
