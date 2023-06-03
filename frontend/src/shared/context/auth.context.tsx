import {
  FC,
  PropsWithChildren,
  createContext,
  useEffect,
  useMemo,
  useState,
} from "react";
import { AuthState } from "../types/user/user.type";
import { NextComponentAuth } from "../types/page/page.type";
import { Router, useRouter } from "next/router";
import { PAGES_LINKS } from "../config/links.config";

interface AuthContextState {
  isAuth: boolean;
  error: string;
  auth: AuthState | null;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextState>(
  {} as AuthContextState
);

export const AuthContextProvider: FC<PropsWithChildren<NextComponentAuth>> = ({
  children,
  Component,
}) => {
  const router = useRouter();
  const [auth, setAuth] = useState<AuthState | null>(null);

  const [isAuth, setIsAuth] = useState(false);
  const [onMountLoading, setOnMountLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (Component.onlyUser && !isAuth) {
      router.push(PAGES_LINKS.Signup.link).catch(err => setError(err));
      return;
    }
  }, [Component.onlyUser, isAuth, router]);

  useEffect(() => {
    //TODO: fetch user and set

    new Promise(resolve => {
      setTimeout(() => {
        resolve(setAuth({ id: 1, username: "Diar", email: "", token: "" }));
      }, 2000);
    }).then(() => {
      setIsAuth(true);
      setOnMountLoading(false);
    });
  }, []);

  const logout = () => {
    setAuth(null);
    setIsAuth(false);
    setError("");
  };

  const value: AuthContextState = useMemo(() => {
    return {
      isAuth: isAuth,
      error: error,
      auth: auth,
      logout: logout,
    };
  }, [error, isAuth, auth]);

  const renderedChildren = !onMountLoading && children;

  return (
    <AuthContext.Provider value={value}>
      {renderedChildren}
    </AuthContext.Provider>
  );
};
