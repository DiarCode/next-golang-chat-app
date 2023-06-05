export interface ChatMessage {
  id: number;
  content: string;
  userId: number;
  roomId: number;
  sendedAt: number;
}
