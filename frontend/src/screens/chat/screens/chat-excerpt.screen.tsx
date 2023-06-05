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
import { ChatApiService } from "@/shared/api/chat/chat.api";

export const ChatExcerptScreen = () => {
  const router = useRouter();
  const { id } = router.query;
  const [socket, setSocket] = useState<WebSocket | null>(null);

  const { chat: data } = useChat(Number(id));

  // useEffect(() => {
  //   if (!id) return;

  //   const client = ChatApiService.getChatRoomWebSocketClient(Number(id));

  //   client.onopen = () => {
  //     console.log("WebSocket connection established");
  //     setSocket(client);
  //   };

  //   client.onmessage = message => {
  //     // Handle incoming messages
  //     const newMessage = message.data;
  //     console.log(newMessage);
  //   };

  //   client.onclose = event => {
  //     console.log("WebSocket connection closed");
  //     setSocket(null);
  //   };

  //   return () => {
  //     // Close the WebSocket connection when the component unmounts
  //     if (client) {
  //       client.close();
  //       setSocket(null);
  //     }
  //   };
  // }, [id]);

  const chat = data ?? chats[0];

  return (
    <AppLayout title={chat.name}>
      <main className="grid grid-cols-3">
        <div className="col-span-3 md:col-span-2 h-full md:pr-7">
          <Link href={PAGES_LINKS.Chat.link}>
            <button className="flex items-center text-gray-500 mb-2">
              <ChevronLeft size={20} />
              <p>Back</p>
            </button>
          </Link>

          <PageTitle content={chat.name} />

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
