import { apiClient } from "@/lib/api-client";
import type { ImpactResult } from "./types";

export const impactApi = {
  analyze: (runId: number, symbolId: number) =>
    apiClient.get<ImpactResult>(`/api/v1/impact?runId=${runId}&symbolId=${symbolId}`),
};
