import { COLORS } from "@/shared/config/colors.config";
import { NAVBAR_LINKS } from "@/shared/config/links.config";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";

const selectedRouteStyle = { color: COLORS.primary };
const defaultRouteStyle = { color: COLORS.black };

export const DesktopNavigation = () => {
  const { route } = useRouter();

  const renderedLinks = Object.entries(NAVBAR_LINKS).map(([_, link]) => {
    const routeStyle =
      route === link.link ? selectedRouteStyle : defaultRouteStyle;

    return (
      <li key={link.name}>
        <Link href={link.link}>
          <p style={{ ...routeStyle }}>{link.name}</p>
        </Link>
      </li>
    );
  });

  return (
    <div className="hidden sm:block">
      <ul className="flex items-center gap-x-7">{renderedLinks}</ul>
    </div>
  );
};
