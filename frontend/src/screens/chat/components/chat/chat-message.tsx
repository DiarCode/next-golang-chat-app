import React, { FC } from "react";

type MessageFromType = "me" | "sender";

interface ChatMessageProps {
  messageFromType: MessageFromType;
}

export const ChatMessage: FC<ChatMessageProps> = ({ messageFromType }) => {
  switch (messageFromType) {
    case "me":
      return <ChatMessageMe />;
    case "sender":
      return <ChatMessageSender />;
    default:
      return null;
  }
};

const ChatMessageMe = () => {
  return (
    <div className="place-self-end bg-primary_light rounded-lg p-2 px-3 w-fit mb-3">
      <p className="text-white text-sm md:text-base">
        Hi all! How are you doing?
      </p>

      <div className="flex justify-end mt-1">
        <p className="text-xs text-gray-300">12:45 pm</p>
      </div>
    </div>
  );
};

const ChatMessageSender = () => {
  return (
    <div className="place-self-start bg-white rounded-lg p-2 px-3 w-fit mb-3">
      <p className="text-xs md:text-sm text-primary">Harrison Ford</p>

      <p className="text-black text-sm md:text-base">
        Hi all! How are you doing?
      </p>

      <div className="flex justify-end mt-1">
        <p className="text-xs text-gray-500">12:45 pm</p>
      </div>
    </div>
  );
};
