import { Chat } from "@/shared/types/chat/chat.type";
import React, { FC } from "react";
import { ChatItem } from "./chat-item";

interface ChatsListProps {
  chats: Chat[];
}

export const ChatsList: FC<ChatsListProps> = ({ chats }) => {
  const renderedChats = chats.map(chat => (
    <ChatItem key={chat.id} chat={chat} />
  ));

  return <div className="grid grid-cols-fluid gap-4">{renderedChats}</div>;
};
