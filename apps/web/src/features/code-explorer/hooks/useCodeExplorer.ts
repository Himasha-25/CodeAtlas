import { useQuery } from "@tanstack/react-query";
import { codeExplorerApi } from "../api";

export const explorerKeys = {
  files: (runId: number) => ["explorer", "files", runId] as const,
  symbols: (fileId: number) => ["explorer", "symbols", fileId] as const,
  search: (runId: number, q: string) => ["explorer", "search", runId, q] as const,
};

export function useFiles(runId: number) {
  return useQuery({
    queryKey: explorerKeys.files(runId),
    queryFn: () => codeExplorerApi.listFiles(runId),
    enabled: runId > 0,
  });
}

export function useSymbols(fileId: number) {
  return useQuery({
    queryKey: explorerKeys.symbols(fileId),
    queryFn: () => codeExplorerApi.listSymbols(fileId),
    enabled: fileId > 0,
  });
}

export function useSymbolSearch(runId: number, q: string) {
  return useQuery({
    queryKey: explorerKeys.search(runId, q),
    queryFn: () => codeExplorerApi.search(runId, q),
    enabled: runId > 0 && q.length > 1,
  });
}
