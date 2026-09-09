import { apiClient } from "@/lib/api-client";
import type { AskRequest, AskResponse, Message } from "./types";

export const assistantApi = {
  ask: (projectId: number, body: AskRequest) =>
    apiClient.post<AskResponse>(`/api/v1/projects/${projectId}/assistant/ask`, body),

  listMessages: (projectId: number, conversationId: number) =>
    apiClient.get<Message[]>(`/api/v1/projects/${projectId}/assistant/conversations/${conversationId}/messages`),
};
