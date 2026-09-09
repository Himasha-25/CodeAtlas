import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { repositoriesApi } from "../api";

export const repoKeys = {
  list: (projectId: number) => ["repositories", projectId] as const,
};

export function useRepositories(projectId: number) {
  return useQuery({
    queryKey: repoKeys.list(projectId),
    queryFn: () => repositoriesApi.list(projectId),
  });
}

export function useUploadRepository(projectId: number) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (file: File) => repositoriesApi.upload(projectId, file),
    onSuccess: () => qc.invalidateQueries({ queryKey: repoKeys.list(projectId) }),
  });
}
