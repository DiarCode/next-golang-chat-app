import { CreateChatModalContextProvider } from "@/screens/chat/context/create-chat-modal.context";
import { ChatScreen } from "@/screens/chat/screens/chat.screen";
import { NextPageAuth } from "@/shared/types/page/page.type";

const ChatPage: NextPageAuth = () => {
  return (
    <CreateChatModalContextProvider>
      <ChatScreen />
    </CreateChatModalContextProvider>
  );
};

ChatPage.onlyUser = true;

export default ChatPage;
