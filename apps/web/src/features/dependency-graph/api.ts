import { apiClient } from "@/lib/api-client";
import type { Dependency } from "./types";

export const dependencyGraphApi = {
  getEdges: (runId: number) =>
    apiClient.get<Dependency[]>(`/api/v1/graph?runId=${runId}`),
};
