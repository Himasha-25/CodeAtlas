import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { documentationApi } from "../api";

export const docsKeys = {
  get: (projectId: number) => ["docs", projectId] as const,
};

export function useDocumentation(projectId: number) {
  return useQuery({
    queryKey: docsKeys.get(projectId),
    queryFn: () => documentationApi.get(projectId),
  });
}

export function useGenerateDocumentation(projectId: number) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (runId: number) => documentationApi.generate(projectId, runId),
    onSuccess: () => qc.invalidateQueries({ queryKey: docsKeys.get(projectId) }),
  });
}
