import { useQuery } from "@tanstack/react-query";
import { ChatApiService } from "../api/chat/chat.api";
import { useContext, useEffect, useLayoutEffect } from "react";
import { ChatContext } from "@/screens/chat/context/chat.context";
import { ChatMessage } from "../types/chat/message.type";
import { useAuth } from "./useAuth";
import { SendMessageDto } from "../types/chat/message.dto";
import { useChatsStore } from "../store/useChatsStore";

export const useChats = () => {
  const { chats, addChat, initChats } = useChatsStore();
  const { data } = useQuery({
    queryKey: ["chats"],
    queryFn: ChatApiService.getAllChatRooms,
  });

  useLayoutEffect(() => {
    initChats(data?.data ?? []);
  }, [data?.data, initChats]);

  return { chats, addChat };
};

export const useChat = (id?: number) => {
  const { auth } = useAuth();
  const {
    messages,
    setMessages,
    socket,
    setSocket,
    chat,
    setChat,
    pushMessage,
  } = useContext(ChatContext);

  const roomResponse = useQuery({
    queryKey: ["rooms", id],
    queryFn: () => ChatApiService.getChatRoomById(id!),
    enabled: Boolean(id),
  });

  const messagesResponse = useQuery({
    queryKey: ["messages", id],
    queryFn: () => ChatApiService.getMessagesByRoomId(id!),
    enabled: Boolean(id),
  });

  useEffect(() => {
    setChat(roomResponse.data?.data ?? null);
    setMessages(messagesResponse.data?.data ?? []);

    return () => {
      setMessages([]);
      setChat(null);
    };
  }, [
    messagesResponse.data?.data,
    roomResponse.data?.data,
    setChat,
    setMessages,
  ]);

  useEffect(() => {
    if (!id) {
      return;
    }

    const client = ChatApiService.getChatRoomWebSocketClient(id);

    client.onopen = () => {
      console.log("WebSocket Connected!");
      setSocket(client);
    };

    client.onmessage = message => {
      const newMessageStr = message?.data ?? null;
      if (!newMessageStr) return;
      const newMessageObj = JSON.parse(newMessageStr) as ChatMessage;
      pushMessage(newMessageObj);
    };

    client.onclose = () => {
      console.log("WebSocket Closed");
      setSocket(null);
    };

    return () => {
      if (client) {
        client.close();
        setSocket(null);
      }
    };
  }, [id, pushMessage, setSocket]);

  const sendMessage = async (msg: string) => {
    if (!chat || !auth || !socket || socket.readyState !== socket.OPEN) {
      return;
    }

    const chatMessage: SendMessageDto = {
      content: msg,
      userId: auth.id,
      roomId: chat.id,
    };

    const res = await ChatApiService.sendMessage(chatMessage);
    if (res.status !== 200) {
      console.log(res.data);
    }
  };

  return { messages, chat, socket, sendMessage };
};
