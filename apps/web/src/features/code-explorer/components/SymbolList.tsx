"use client";

import { useSymbols } from "../hooks/useCodeExplorer";
import { Badge } from "@/components/ui/badge";

interface Props { fileId: number }

export function SymbolList({ fileId }: Props) {
  const { data: symbols, isLoading } = useSymbols(fileId);

  if (isLoading) return <p className="text-sm text-gray-400">Loading symbols…</p>;
  if (!symbols?.length) return <p className="text-sm text-gray-400">No symbols.</p>;

  return (
    <ul className="space-y-1 text-sm">
      {symbols.map((s) => (
        <li key={s.id} className="flex items-center gap-2 px-2 py-1">
          <Badge variant="secondary">{s.kind}</Badge>
          <span className={s.exported ? "font-medium" : "text-gray-600"}>{s.name}</span>
          <span className="ml-auto text-xs text-gray-400">:{s.line}</span>
        </li>
      ))}
    </ul>
  );
}
