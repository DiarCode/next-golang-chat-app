import {
  FC,
  PropsWithChildren,
  createContext,
  useCallback,
  useEffect,
  useMemo,
  useState,
} from "react";
import { NextComponentAuth } from "../types/page/page.type";
import { useRouter } from "next/router";
import { PAGES_LINKS } from "../config/links.config";
import { AuthState, LoginDto, SignupDto } from "../types/auth/auth.dto";
import { AuthApiService } from "../api/auth/auth.api";
import { useCookies } from "react-cookie";

interface AuthContextState {
  isAuth: boolean;
  error: string;
  auth: AuthState | null;

  logout: () => void;
  login: (dto: LoginDto) => void;
  signup: (dto: SignupDto) => void;
}

export const AuthContext = createContext<AuthContextState>(
  {} as AuthContextState
);

const AUTH_COOKIES_NAME = "auth";

export const AuthContextProvider: FC<PropsWithChildren<NextComponentAuth>> = ({
  children,
  Component,
}) => {
  const router = useRouter();
  const [cookies, setCookie, removeCookie] = useCookies([AUTH_COOKIES_NAME]);

  const [auth, setAuth] = useState<AuthState | null>(null);
  const [isAuth, setIsAuth] = useState(false);
  const [onMountLoading, setOnMountLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (Component.onlyUser && !isAuth) {
      router.push(PAGES_LINKS.Signup.link).catch(err => setError(err));
    }
  }, [Component.onlyUser, isAuth, router]);

  useEffect(() => {
    if (cookies.auth) {
      const authStateCookies = cookies.auth as AuthState;
      setAuth(authStateCookies);
      setIsAuth(true);
      setOnMountLoading(false);
      return;
    }

    setAuth(null);
    setIsAuth(false);
    setOnMountLoading(false);
  }, [cookies.auth]);

  const logout = useCallback(() => {
    setAuth(null);
    setIsAuth(false);
    setError("");
    removeCookie("auth");
  }, [removeCookie]);

  const login = useCallback(
    async (dto: LoginDto) => {
      try {
        const { data, status } = await AuthApiService.login(dto);
        if (status !== 200) {
          setError(data?.message ?? "Failed to login");
          return;
        }

        setAuth(data);
        setIsAuth(true);
        setCookie("auth", data);
      } catch (error) {
        if (typeof error === "string") {
          setError(error);
          return;
        }
      }
    },
    [setCookie]
  );

  const signup = useCallback(async (dto: SignupDto) => {
    try {
      const { data, status } = await AuthApiService.signup(dto);
      if (status !== 200) {
        setError(data?.message ?? "Failed to login");
        return;
      }

      setAuth(data);
      setIsAuth(true);
      setCookie("auth", data);
    } catch (error) {
      if (typeof error === "string") {
        setError(error);
        return;
      }
    }
  }, [setCookie]);

  const value: AuthContextState = useMemo(() => {
    return {
      isAuth: isAuth,
      error: error,
      auth: auth,
      logout: logout,
      login: login,
      signup: signup,
    };
  }, [isAuth, error, auth, logout, login, signup]);

  const renderedChildren = !onMountLoading && children;

  return (
    <AuthContext.Provider value={value}>
      {renderedChildren}
    </AuthContext.Provider>
  );
};
