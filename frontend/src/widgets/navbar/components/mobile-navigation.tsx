import Link from "next/link";
import { NAVBAR_LINKS } from "@/shared/config/links.config";
import { Menu } from "react-feather";
import { useNavbar } from "../hooks/useNavbar";

export const MobileNavigation = () => {
  const {
    mobileMenu: { open, setOpen },
  } = useNavbar();

  const renderedLinks = Object.entries(NAVBAR_LINKS).map(([_, link]) => (
    <li key={link.name}>
      <Link href={link.link}>
        <p className="text-white text-base">{link.name}</p>
      </Link>
    </li>
  ));

  const onEmptyAreaClick = () => {
    setOpen(false);
  };

  return (
    <div className="sm:hidden">
      <button onClick={() => setOpen(!open)}>
        <Menu />
      </button>

      <div
        className={`z-50 flex absolute top-[74px] left-0 w-full h-screen
            flex-1 justify-end transition-all duration-300 ease`}
        style={{ visibility: open ? "visible" : "hidden" }}
        onClick={onEmptyAreaClick}
      >
        <div
          onClick={e => e.stopPropagation()}
          className={`flex w-full h-fit py-7 px-10 justify-center rounded-xl
            transition-transform duration-300 ease-in-out bg-primary ${
              open ? " translate-x-0" : "translate-x-full"
            }`}
        >
          <ul className="flex flex-col items-center gap-y-4">
            {renderedLinks}
          </ul>
        </div>
      </div>
    </div>
  );
};
