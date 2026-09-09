"use client";

import { useDocumentation, useGenerateDocumentation } from "../hooks/useDocumentation";
import { Button } from "@/components/ui/button";

interface Props { projectId: number; runId: number }

export function DocViewer({ projectId, runId }: Props) {
  const { data: doc, isLoading } = useDocumentation(projectId);
  const { mutate: generate, isPending } = useGenerateDocumentation(projectId);

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-lg font-semibold">Documentation</h2>
        <Button size="sm" onClick={() => generate(runId)} disabled={isPending}>
          {isPending ? "Generating…" : "Regenerate"}
        </Button>
      </div>
      {isLoading && <p className="text-sm text-gray-400">Loading…</p>}
      {doc ? (
        <pre className="whitespace-pre-wrap rounded border bg-gray-50 p-4 text-sm">{doc.content}</pre>
      ) : (
        !isLoading && <p className="text-sm text-gray-400">No documentation yet. Click Regenerate.</p>
      )}
    </div>
  );
}
