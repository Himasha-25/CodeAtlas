import { apiClient } from "@/lib/api-client";
import type { Documentation } from "./types";

export const documentationApi = {
  get: (projectId: number) =>
    apiClient.get<Documentation>(`/api/v1/projects/${projectId}/docs`),

  generate: (projectId: number, runId: number) =>
    apiClient.post<Documentation>(`/api/v1/projects/${projectId}/docs/generate?runId=${runId}`),
};
