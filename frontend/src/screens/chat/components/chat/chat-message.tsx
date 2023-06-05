import { useUser } from "@/shared/hooks/useUser";
import { ChatMessage as ChatMessageType } from "@/shared/types/chat/message.type";
import { formatDateFromUnix } from "@/shared/utils/date-formatter";
import React, { FC } from "react";

export type MessageFromType = "me" | "sender";

interface ChatMessageProps {
  message: ChatMessageType;
  messageFromType: MessageFromType;
}

export const ChatMessage: FC<ChatMessageProps> = ({
  messageFromType,
  message,
}) => {
  switch (messageFromType) {
    case "me":
      return <ChatMessageMe message={message} />;
    case "sender":
      return <ChatMessageSender message={message} />;
    default:
      return null;
  }
};

interface ChatMessageComponentProps {
  message: ChatMessageType;
}

const ChatMessageMe: FC<ChatMessageComponentProps> = ({ message }) => {
  return (
    <div className="place-self-end bg-primary_light rounded-lg p-2 px-3 w-fit mb-3">
      <p className="text-white text-sm md:text-base">{message.content}</p>

      <div className="flex justify-end mt-1">
        <p className="text-xs text-gray-300">
          {formatDateFromUnix(message.sendedAt)}
        </p>
      </div>
    </div>
  );
};

const ChatMessageSender: FC<ChatMessageComponentProps> = ({ message }) => {
  const { data } = useUser(message.userId);
  return (
    <div className="place-self-start bg-white rounded-lg p-2 px-3 w-fit mb-3">
      <p className="text-xs md:text-sm text-primary">
        {data?.data.username ?? "Unknown"}
      </p>

      <p className="text-black text-sm md:text-base">{message.content}</p>

      <div className="flex justify-end mt-1">
        <p className="text-xs text-gray-500">
          {formatDateFromUnix(message.sendedAt)}
        </p>
      </div>
    </div>
  );
};
