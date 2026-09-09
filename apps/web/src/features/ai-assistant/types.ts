export interface Snippet {
  filePath: string;
  symbol: string;
  content: string;
}

export interface Message {
  id: number;
  conversationId: number;
  role: "user" | "assistant";
  content: string;
  sources?: Snippet[];
  createdAt: string;
}

export interface AskRequest {
  conversationId?: number;
  runId: number;
  question: string;
}

export interface AskResponse {
  answer: string;
  sources: Snippet[];
}
