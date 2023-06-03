import { AppLayout } from "@/shared/layouts/app-layout";
import { Button } from "@/shared/ui/button";
import { PageTitle } from "@/shared/ui/title";
import React, { useContext } from "react";
import { Plus } from "react-feather";
import { ChatsList } from "../components/chats-list";
import { chats } from "@/shared/mocks/chats";
import { CreateChatModal } from "../components/create-chat-modal";
import { CreateChatModalContext } from "../context/create-chat-modal.context";
import { useChats } from "@/shared/hooks/useChat";

export const ChatScreen = () => {
  const { data } = useChats();
  const { setVisible } = useContext(CreateChatModalContext);

  return (
    <AppLayout title="Chat">
      <div className="flex justify-between items-center">
        <PageTitle content="Chats list" />
        <Button
          style={{ gap: 4 }}
          onClick={() => {
            setVisible(true);
          }}
        >
          <Plus size={15} />
          <p>Create chat</p>
        </Button>
      </div>

      <div className="mt-6">
        <ChatsList chats={data?.data ?? chats} />
      </div>

      <CreateChatModal />
    </AppLayout>
  );
};
