import React from "react";
import { ChatItem } from "./chat-item";
import { chats } from "@/shared/mocks/chats";
import { useChats } from "@/shared/hooks/useChat";

export const SidebarChatDetails = () => {
  const { data } = useChats();
  const renderedChats = data?.data.map(chat => (
    <ChatItem key={chat.id} chat={chat} />
  ));
  
  return (
    <>
      <h4 className="font-semibold text-base mb-3">Chats</h4>
      <div className="flex flex-col gap-3">{renderedChats}</div>
    </>
  );
};
