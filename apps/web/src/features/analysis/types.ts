export type AnalysisStatus = "pending" | "running" | "completed" | "failed";

export interface AnalysisRun {
  id: number;
  repositoryId: number;
  status: AnalysisStatus;
  error?: string;
  startedAt?: string;
  completedAt?: string;
  createdAt: string;
}
