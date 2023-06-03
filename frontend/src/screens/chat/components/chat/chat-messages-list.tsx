import React, {
  MutableRefObject,
  useEffect,
  useRef,
} from "react";
import { ChatMessage } from "./chat-message";

export const ChatMessagesList = () => {
  const messagesEndRef = useRef(
    null
  ) as MutableRefObject<HTMLDivElement | null>;

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    scrollToBottom();
  }, []);

  return (
    <div className="bg-slate-200 overflow-y-auto h-[65vh] rounded-xl border-2 p-4 relative">
      <ChatMessage messageFromType="sender" />
      <ChatMessage messageFromType="sender" />

      <ChatMessage messageFromType="me" />

      <ChatMessage messageFromType="sender" />
      <ChatMessage messageFromType="sender" />

      <ChatMessage messageFromType="me" />
      <ChatMessage messageFromType="sender" />
      <ChatMessage messageFromType="sender" />

      <ChatMessage messageFromType="me" />
      <ChatMessage messageFromType="sender" />
      <ChatMessage messageFromType="sender" />

      <ChatMessage messageFromType="me" />
      <ChatMessage messageFromType="sender" />
      <ChatMessage messageFromType="sender" />

      <ChatMessage messageFromType="me" />

      <div ref={messagesEndRef} />
    </div>
  );
};
