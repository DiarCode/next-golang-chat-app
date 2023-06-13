import { Chat } from "@/shared/types/chat/chat.type";
import { ChatMessage } from "@/shared/types/chat/message.type";
import {
  Dispatch,
  FC,
  PropsWithChildren,
  SetStateAction,
  createContext,
  useCallback,
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

  const pushMessage = useCallback((message: ChatMessage) => {
    if(messages.find(msg => msg.id === message.id)) return
    setMessages(prev => [...prev, message]);
  }, [messages]);

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
    [chat, messages, pushMessage, socket]
  );

  return (
    <ChatContext.Provider value={memoValue}>{children}</ChatContext.Provider>
  );
};
