import { useQuery } from "@tanstack/react-query";
import { dependencyGraphApi } from "../api";
import type { GraphData } from "../types";

export function useDependencyGraph(runId: number) {
  return useQuery({
    queryKey: ["graph", runId],
    queryFn: async (): Promise<GraphData> => {
      const edges = await dependencyGraphApi.getEdges(runId);
      const nodeIds = new Set<number>();
      edges.forEach((e) => { nodeIds.add(e.fromId); nodeIds.add(e.toId); });
      return {
        nodes: Array.from(nodeIds).map((id) => ({ id: String(id), label: String(id) })),
        edges: edges.map((e) => ({
          id: String(e.id),
          source: String(e.fromId),
          target: String(e.toId),
          label: e.edgeType,
        })),
      };
    },
    enabled: runId > 0,
  });
}
