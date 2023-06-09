import {
  Dispatch,
  FC,
  PropsWithChildren,
  SetStateAction,
  createContext,
  useCallback,
  useMemo,
  useState,
} from "react";

interface NavbarContextState {
  mobileMenu: {
    open: boolean;
    setOpen: Dispatch<SetStateAction<boolean>>;
  };
  profileMenu: {
    open: boolean;
    setOpen: Dispatch<SetStateAction<boolean>>;
  };
}

export const NavbarContext = createContext<NavbarContextState>({} as NavbarContextState);

export const NavbarContextProvider: FC<PropsWithChildren> = ({ children }) => {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  const [profileMenuOpen, setProfileMenuOpen] = useState(false);

  const closeAllModals = () => {
    setMobileMenuOpen(false);
    setProfileMenuOpen(false);
  };

  const setMobileMenuOpenWrapper = useCallback(
    (state: SetStateAction<boolean>) => {
      closeAllModals();
      setMobileMenuOpen(state);
    },
    []
  );

  const setProfileMenuOpenWrapper = useCallback(
    (state: SetStateAction<boolean>) => {
      closeAllModals();
      setProfileMenuOpen(state);
    },
    []
  );

  const value: NavbarContextState = useMemo(() => {
    return {
      mobileMenu: {
        open: mobileMenuOpen,
        setOpen: setMobileMenuOpenWrapper,
      },
      profileMenu: {
        open: profileMenuOpen,
        setOpen: setProfileMenuOpenWrapper,
      },
    };
  }, [
    mobileMenuOpen,
    profileMenuOpen,
    setMobileMenuOpenWrapper,
    setProfileMenuOpenWrapper,
  ]);

  return (
    <NavbarContext.Provider value={value}>{children}</NavbarContext.Provider>
  );
};
