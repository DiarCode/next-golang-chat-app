import React from "react";
import { Send } from "react-feather";

export const ChatInput = () => {
  return (
    <div className="w-full bg-slate-200 mt-3 rounded-xl h-16 p-3 px-4">
      <div className="h-full flex items-center gap-x-3 ">
        <input
          type="text"
          className="h-full w-full rounded-md border border-gray-300 p-3"
          placeholder="Type message..."
        />

        <button className="px-4 py-[6px] bg-primary text-white rounded-md flex items-center gap-2">
          <p>Send</p>
          <Send size={16} />
        </button>
      </div>
    </div>
  );
};
