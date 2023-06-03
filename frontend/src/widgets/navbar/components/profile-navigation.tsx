import { useAuth } from "@/shared/hooks/useAuth";
import { useNavbar } from "../hooks/useNavbar";
import { User } from "react-feather";

export const ProfileNavigation = () => {
  const { auth, logout } = useAuth();

  const {
    profileMenu: { open, setOpen },
  } = useNavbar();

  const onEmptyAreaClick = () => {
    setOpen(false);
  };

  return (
    <div>
      <button onClick={() => setOpen(!open)}>
        <div className="flex items-center gap-x-2">
          <User size={18} />
          <p className="font-medium text-sm md:text-base">{auth?.username}</p>
        </div>
      </button>

      <div
        className={`z-50 flex absolute top-[74px] left-0 w-full h-screen
            flex-1 justify-end transition-all duration-300 ease`}
        style={{ visibility: open ? "visible" : "hidden" }}
        onClick={onEmptyAreaClick}
      >
        <div
          onClick={e => e.stopPropagation()}
          className={`flex w-full md:w-fit h-fit py-4 px-10 justify-center rounded-xl
            transition-transform duration-300 ease-in-out bg-primary ${
              open ? " translate-x-0" : "translate-x-full"
            }`}
        >
          <div className="flex flex-col items-center gap-y-4">
            <button onClick={logout}>
              <p className="cursor-pointer text-white text-base">Logout</p>
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};
