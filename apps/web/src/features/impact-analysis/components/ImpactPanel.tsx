"use client";

import { useState } from "react";
import { useImpactAnalysis } from "../hooks/useImpact";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { FormField, FormLabel } from "@/components/ui/form";

interface Props { runId: number }

export function ImpactPanel({ runId }: Props) {
  const [symbolId, setSymbolId] = useState(0);
  const [submitted, setSubmitted] = useState(0);
  const { data, isLoading } = useImpactAnalysis(runId, submitted);

  return (
    <div className="space-y-4">
      <div className="flex items-end gap-2">
        <FormField>
          <FormLabel>Symbol ID</FormLabel>
          <Input
            type="number"
            value={symbolId || ""}
            onChange={(e) => setSymbolId(Number(e.target.value))}
            className="w-32"
          />
        </FormField>
        <Button onClick={() => setSubmitted(symbolId)} disabled={!symbolId}>
          Analyse
        </Button>
      </div>
      {isLoading && <p className="text-sm text-gray-400">Analysing…</p>}
      {data && (
        <div className="text-sm">
          <p className="font-medium">Origin: {data.originId}</p>
          <p className="mt-1 text-gray-600">
            {data.affectedIds.length} affected node{data.affectedIds.length !== 1 ? "s" : ""}:
          </p>
          <ul className="mt-1 list-inside list-disc text-gray-500">
            {data.affectedIds.map((id) => <li key={id}>{id}</li>)}
          </ul>
        </div>
      )}
    </div>
  );
}
