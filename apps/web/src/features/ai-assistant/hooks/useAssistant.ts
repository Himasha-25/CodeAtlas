import { useMutation, useQuery } from "@tanstack/react-query";
import { assistantApi } from "../api";
import type { AskRequest } from "../types";

export function useMessages(projectId: number, conversationId: number) {
  return useQuery({
    queryKey: ["assistant", "messages", projectId, conversationId],
    queryFn: () => assistantApi.listMessages(projectId, conversationId),
    enabled: conversationId > 0,
  });
}

export function useAsk(projectId: number) {
  return useMutation({
    mutationFn: (body: AskRequest) => assistantApi.ask(projectId, body),
  });
}
