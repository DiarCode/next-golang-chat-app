import { useQuery } from "@tanstack/react-query";
import { ChatApiService } from "../api/chat/chat.api";
import { useContext, useEffect } from "react";
import { ChatContext } from "@/screens/chat/context/chat.context";
import { ChatMessage } from "../types/chat/message.type";
import { useAuth } from "./useAuth";
import { Chat } from "../types/chat/chat.type";
import { SendMessageDto } from "../types/chat/message.dto";
import { getApiUrl } from "../api/api";

export const useChats = () =>
  useQuery({ queryKey: ["chats"], queryFn: ChatApiService.getAllChatRooms });

export const useChat = (id?: number) => {
  const { auth } = useAuth();
  const chatContext = useContext(ChatContext);

  const roomResponse = useQuery({
    queryKey: ["rooms", id],
    queryFn: () => ChatApiService.getChatRoomById(id!),
    enabled: !chatContext.chat && Boolean(id),
  });

  const messagesResponse = useQuery({
    queryKey: ["messages", id],
    queryFn: () => ChatApiService.getMessagesByRoomId(id!),
    enabled: !chatContext.chat && Boolean(id),
  });

  useEffect(() => {
    console.log("MEEESAGES", messagesResponse.data?.data);
    if (!chatContext.chat) {
      chatContext.setChat(roomResponse.data?.data ?? null);
      chatContext.setMessages(messagesResponse.data?.data ?? []);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [chatContext.chat, messagesResponse.data?.data, roomResponse.data?.data]);

  useEffect(() => {
    if (!id) {
      return;
    }

    const client = ChatApiService.getChatRoomWebSocketClient(id);

    client.onopen = () => {
      console.log("WebSocket connection established");
      chatContext.setSocket(client);
    };

    client.onmessage = message => {
      const newMessage = message.data;
      console.log(newMessage);
      chatContext.pushMessage(newMessage);
    };

    client.onclose = () => {
      console.log("WebSocket connection closed");
      chatContext.setSocket(null);
    };

    return () => {
      if (client) {
        client.close();
        chatContext.setMessages([]);
        chatContext.setChat(null);
        chatContext.setSocket(null);
      }
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [id]);

  const sendMessage = async (msg: string) => {
    if (
      !chatContext.chat ||
      !auth ||
      !chatContext.socket ||
      chatContext.socket.readyState !== chatContext.socket.OPEN
    ) {
      return;
    }

    const chatMessage: SendMessageDto = {
      content: msg,
      userId: auth.id,
      roomId: chatContext.chat.id,
    };

    const res = await ChatApiService.sendMessage(chatMessage);
    if (res.status !== 200) {
      console.log(res.data);
    }
  };

  return { ...chatContext, sendMessage };
};
