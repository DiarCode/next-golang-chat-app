import React from "react";
import { ChatInput } from "./chat-input";
import { ChatMessagesList } from "./chat-messages-list";

export const Chat = () => {
  return (
    <div>
      <ChatMessagesList />
      <ChatInput />
    </div>
  );
};
