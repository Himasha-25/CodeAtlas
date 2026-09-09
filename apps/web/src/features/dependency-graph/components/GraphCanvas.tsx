"use client";

import { useDependencyGraph } from "../hooks/useDependencyGraph";

interface Props { runId: number }

export function GraphCanvas({ runId }: Props) {
  const { data, isLoading, error } = useDependencyGraph(runId);

  if (isLoading) return <p className="text-sm text-gray-400">Loading graph…</p>;
  if (error) return <p className="text-sm text-red-500">Failed to load graph.</p>;
  if (!data?.nodes.length) return <p className="text-sm text-gray-400">No dependency data.</p>;

  return (
    <div className="rounded border bg-gray-50 p-4 text-sm text-gray-500">
      <p>{data.nodes.length} nodes · {data.edges.length} edges</p>
      <p className="mt-1 text-xs">React Flow canvas goes here.</p>
    </div>
  );
}
