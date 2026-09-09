import { useQuery } from "@tanstack/react-query";
import { impactApi } from "../api";

export function useImpactAnalysis(runId: number, symbolId: number) {
  return useQuery({
    queryKey: ["impact", runId, symbolId],
    queryFn: () => impactApi.analyze(runId, symbolId),
    enabled: runId > 0 && symbolId > 0,
  });
}
