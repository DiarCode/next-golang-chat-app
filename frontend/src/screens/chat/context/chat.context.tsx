import { Chat } from "@/shared/types/chat/chat.type";
import { ChatMessage } from "@/shared/types/chat/message.type";
import {
  Dispatch,
  FC,
  PropsWithChildren,
  SetStateAction,
  createContext,
  useMemo,
  useState,
} from "react";

interface ChatContextState {
  chat: Chat | null;
  setChat: Dispatch<SetStateAction<Chat | null>>;

  messages: ChatMessage[];
  setMessages: Dispatch<SetStateAction<ChatMessage[]>>;
  pushMessage: (message: ChatMessage) => void;

  socket: WebSocket | null;
  setSocket: Dispatch<SetStateAction<WebSocket | null>>;
}

export const ChatContext = createContext<ChatContextState>(
  {} as ChatContextState
);

export const ChatContextProvider: FC<PropsWithChildren> = ({ children }) => {
  const [chat, setChat] = useState<Chat | null>(null);
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const [socket, setSocket] = useState<WebSocket | null>(null);

  const pushMessage = (message: ChatMessage) => {
    setMessages(prev => [...prev, message]);
  };

  const memoValue: ChatContextState = useMemo(
    () => ({
      chat,
      setChat,
      messages,
      socket,
      setSocket,
      pushMessage,
      setMessages,
    }),
    [chat, messages, socket]
  );

  return (
    <ChatContext.Provider value={memoValue}>{children}</ChatContext.Provider>
  );
};
