import { ChatExcerptScreen } from "@/screens/chat/screens/chat-excerpt.screen";
import { NextPageAuth } from "@/shared/types/page/page.type";

const ChatExcerptPage: NextPageAuth = () => {
  return <ChatExcerptScreen />;
};

ChatExcerptPage.onlyUser = true;

export default ChatExcerptPage;
