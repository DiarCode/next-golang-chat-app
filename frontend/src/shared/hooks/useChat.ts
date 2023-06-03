import { useQuery } from "@tanstack/react-query";
import { ChatApiService } from "../api/chat/chat.api";

export const useChats = () =>
  useQuery({ queryKey: ["chats"], queryFn: ChatApiService.getAllChats });

export const useChat = (id: number) =>
  useQuery({
    queryKey: ["posts", id],
    queryFn: () => ChatApiService.getChatById(id),
    enabled: Boolean(id),
  });
