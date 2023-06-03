import { NavbarContextProvider } from "@/widgets/navbar/context/navbar.context";
import React, { FC, PropsWithChildren } from "react";
import { AuthContextProvider } from "../context/auth.context";

export const ProvidersLayout: FC<PropsWithChildren> = ({ children }) => {
  return <NavbarContextProvider>{children}</NavbarContextProvider>;
};
