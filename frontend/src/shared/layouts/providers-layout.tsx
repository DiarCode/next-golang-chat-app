import { NavbarContextProvider } from "@/widgets/navbar/context/navbar.context";
import React, { FC, PropsWithChildren } from "react";
import { ChatContextProvider } from "@/screens/chat/context/chat.context";

export const ProvidersLayout: FC<PropsWithChildren> = ({ children }) => {
  return (
    <NavbarContextProvider>
      <ChatContextProvider>{children}</ChatContextProvider>
    </NavbarContextProvider>
  );
};
