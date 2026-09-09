import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { analysisApi } from "../api";
import type { AnalysisStatus } from "../types";

const ACTIVE: AnalysisStatus[] = ["pending", "running"];

export const analysisKeys = {
  status: (repositoryId: number) => ["analysis", repositoryId] as const,
};

export function useAnalysisStatus(repositoryId: number) {
  return useQuery({
    queryKey: analysisKeys.status(repositoryId),
    queryFn: () => analysisApi.status(repositoryId),
    refetchInterval: (query) =>
      query.state.data && ACTIVE.includes(query.state.data.status) ? 2000 : false,
  });
}

export function useStartAnalysis(repositoryId: number) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: () => analysisApi.start(repositoryId),
    onSuccess: () => qc.invalidateQueries({ queryKey: analysisKeys.status(repositoryId) }),
  });
}
