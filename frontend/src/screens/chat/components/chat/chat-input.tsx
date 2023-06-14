import { useChat } from "@/shared/hooks/useChat";
import React, { useState } from "react";
import { Send } from "react-feather";

export const ChatInput = () => {
  const { sendMessage } = useChat();
  const [message, setMessage] = useState("");

  const onSendSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!message || message.length === 0) return;

    sendMessage(message);
    setMessage("");
  }

  return (
    <div className="w-full bg-slate-200 mt-3 rounded-xl h-14 p-2">
      <form className="h-full flex items-center gap-x-3" onSubmit={onSendSubmit}>
        <input
          type="text"
          className="h-full w-full focus:outline-none  rounded-md border border-gray-300 p-3"
          placeholder="Type message..."
          value={message}
          onChange={e => setMessage(e.target.value)}
        />

        <button
          type="submit"
          className="px-4 h-full bg-primary text-white rounded-md flex items-center gap-2"
        >
          <p>Send</p>
          <Send size={16} />
        </button>
      </form>
    </div>
  );
};
