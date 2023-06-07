import axios from "axios";
import { getApiUrl } from "../api";
import { CreateChatDto } from "@/shared/types/chat/chat.dto";
import { Chat } from "@/shared/types/chat/chat.type";
import { ChatMessage } from "@/shared/types/chat/message.type";
import { SendMessageDto } from "@/shared/types/chat/message.dto";

export class ChatApiService {
  public static createChatRoom(dto: CreateChatDto) {
    return axios.post<Chat>(getApiUrl("chat/rooms"), dto);
  }

  public static getAllChatRooms() {
    return axios.get<Chat[]>(getApiUrl("chat/rooms"));
  }

  public static getChatRoomById(id: number) {
    return axios.get<Chat>(getApiUrl(`chat/rooms/${id}`));
  }

  public static getMessagesByRoomId(id: number) {
    return axios.get<ChatMessage[]>(getApiUrl(`chat/rooms/${id}/messages`));
  }

  public static sendMessage(dto: SendMessageDto) {
    return axios.post(getApiUrl(`chat/rooms/send`), dto);
  }

  public static getChatRoomWebSocketClient(id: number) {
    const url = `ws://localhost:8080/api/v1/chat/rooms/${id}/ws`;
    const client = new WebSocket(url);
    return client;
  }
}
