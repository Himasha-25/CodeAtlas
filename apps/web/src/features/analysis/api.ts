import { apiClient } from "@/lib/api-client";
import type { AnalysisRun } from "./types";

export const analysisApi = {
  start: (repositoryId: number) =>
    apiClient.post<AnalysisRun>(`/api/v1/repositories/${repositoryId}/analysis`),

  status: (repositoryId: number) =>
    apiClient.get<AnalysisRun>(`/api/v1/repositories/${repositoryId}/analysis`),
};
