import { AppLayout } from "@/shared/layouts/app-layout";
import { PageTitle } from "@/shared/ui/title";
import React, { useEffect, useState } from "react";
import { SidebarChatDetails } from "../components/sidebar-chat-details";
import { Chat } from "../components/chat/chat";
import { chats } from "@/shared/mocks/chats";
import { ChevronLeft } from "react-feather";
import { useRouter } from "next/router";
import Link from "next/link";
import { PAGES_LINKS } from "@/shared/config/links.config";
import { useChat } from "@/shared/hooks/useChat";

export const ChatExcerptScreen = () => {
  const router = useRouter();
  const { id } = router.query;

  const { chat: data } = useChat(Number(id));

  const chat = data;

  return (
    <AppLayout title={chat?.name ?? "Unknown"}>
      <main className="grid grid-cols-3">
        <div className="col-span-3 md:col-span-2 h-full md:pr-7">
          <Link href={PAGES_LINKS.Chat.link}>
            <button className="flex items-center text-gray-500 mb-2">
              <ChevronLeft size={20} />
              <p>Back</p>
            </button>
          </Link>

          <PageTitle content={chat?.name ?? "Unknown"} />

          <div className="mt-4">
            <Chat />
          </div>
        </div>

        <div className="hidden md:block border-l col-span-1 h-full pl-7">
          <SidebarChatDetails />
        </div>
      </main>
    </AppLayout>
  );
};
