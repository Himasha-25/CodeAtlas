"use client";

import { useAnalysisStatus } from "../hooks/useAnalysis";
import { Badge } from "@/components/ui/badge";
import type { AnalysisStatus } from "../types";

const variantMap: Record<AnalysisStatus, "secondary" | "warning" | "success" | "destructive"> = {
  pending: "secondary",
  running: "warning",
  completed: "success",
  failed: "destructive",
};

interface Props { repositoryId: number }

export function AnalysisStatusBadge({ repositoryId }: Props) {
  const { data: run } = useAnalysisStatus(repositoryId);
  if (!run) return null;
  return <Badge variant={variantMap[run.status]}>{run.status}</Badge>;
}
