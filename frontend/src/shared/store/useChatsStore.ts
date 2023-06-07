import { create } from "zustand";
import { Chat } from "../types/chat/chat.type";

interface ChatsState {
  chats: Chat[];
  addChat: (chat: Chat) => void;
  initChats: (chats: Chat[]) => void;
}

export const useChatsStore = create<ChatsState>()(set => ({
  chats: [],
  addChat: (chat: Chat) => set(state => ({ chats: [...state.chats, chat] })),
  initChats: (chats: Chat[]) => set(state => ({ chats: chats })),
}));
