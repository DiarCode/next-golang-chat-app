import axios from "axios";
import { getApiUrl } from "../api";
import { CreateChatDto } from "@/shared/types/chat/chat.dto";
import { Chat } from "@/shared/types/chat/chat.type";

export class ChatApiService {
  public static createChat(dto: CreateChatDto) {
    return axios.post<Chat>(getApiUrl("chats"), dto);
  }

  public static getAllChats() {
    return axios.get<Chat[]>(getApiUrl("chats"));
  }

  public static getChatById(id: number) {
    return axios.get<Chat>(getApiUrl(`chats/${id}`));
  }
}
