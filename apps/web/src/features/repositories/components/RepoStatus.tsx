"use client";

import { useAnalysisStatus, useStartAnalysis } from "@/features/analysis/hooks/useAnalysis";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import type { AnalysisStatus } from "@/features/analysis/types";

const statusVariant: Record<AnalysisStatus, "default" | "warning" | "success" | "destructive" | "secondary"> = {
  pending: "secondary",
  running: "warning",
  completed: "success",
  failed: "destructive",
};

interface Props { repositoryId: number }

export function RepoStatus({ repositoryId }: Props) {
  const { data: run, isLoading } = useAnalysisStatus(repositoryId);
  const { mutate: start, isPending } = useStartAnalysis(repositoryId);

  if (isLoading) return <p className="text-sm text-gray-400">Loading…</p>;

  return (
    <div className="flex items-center gap-3">
      {run ? (
        <Badge variant={statusVariant[run.status]}>{run.status}</Badge>
      ) : (
        <Badge variant="secondary">no analysis</Badge>
      )}
      <Button size="sm" onClick={() => start()} disabled={isPending || run?.status === "running"}>
        {isPending ? "Starting…" : "Run analysis"}
      </Button>
    </div>
  );
}
