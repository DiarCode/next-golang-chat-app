import React, { MutableRefObject, useEffect, useRef } from "react";
import { ChatMessage, MessageFromType } from "./chat-message";
import { useChat } from "@/shared/hooks/useChat";
import { useAuth } from "@/shared/hooks/useAuth";

export const ChatMessagesList = () => {
  const { messages } = useChat();
  const { auth } = useAuth();

  const messagesEndRef = useRef(
    null
  ) as MutableRefObject<HTMLDivElement | null>;

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  const renderedMessages = messages.map(message => {
    const messageType: MessageFromType =
      auth?.id === message.userId ? "me" : "sender";

    return (
      <ChatMessage
        message={message}
        key={message.id}
        messageFromType={messageType}
      />
    );
  });

  useEffect(() => {
    if (renderedMessages.length > 0) {
      scrollToBottom();
    }
  }, [renderedMessages.length]);

  return (
    <div className="bg-slate-200 overflow-y-auto h-[55vh] md:h-[65vh] rounded-xl border-2 p-4 relative">
      <div className="flex flex-col">{renderedMessages}</div>

      <div ref={messagesEndRef} />
    </div>
  );
};
