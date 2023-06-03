import { PAGES_LINKS } from "@/shared/config/links.config";
import { Chat } from "@/shared/types/chat/chat.type";
import Link from "next/link";
import React, { FC } from "react";

interface ChatItemProps {
  chat: Chat;
}

export const ChatItem: FC<ChatItemProps> = ({ chat }) => {
  const chatLink = PAGES_LINKS.Chat.subs.ChatExcerpt.link(chat.id);

  return (
    <Link href={chatLink}>
      <div
        className="bg-gray-100 border text-white rounded-lg p-4 text-sm cursor-pointer
        hover:bg-gray-200"
      >
        <p className="text-gray-600">Chat #{chat.id}</p>
        <p className="text-primary mt-1 font-semibold text-base">
          {chat.name}
        </p>
      </div>
    </Link>
  );
};
