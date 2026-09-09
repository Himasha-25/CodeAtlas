import { apiClient } from "@/lib/api-client";
import type { CodeFile, Symbol } from "./types";

export const codeExplorerApi = {
  listFiles: (runId: number) =>
    apiClient.get<CodeFile[]>(`/api/v1/explorer/files?runId=${runId}`),

  listSymbols: (fileId: number) =>
    apiClient.get<Symbol[]>(`/api/v1/explorer/files/${fileId}/symbols`),

  search: (runId: number, q: string) =>
    apiClient.get<Symbol[]>(`/api/v1/explorer/search?runId=${runId}&q=${encodeURIComponent(q)}`),
};
