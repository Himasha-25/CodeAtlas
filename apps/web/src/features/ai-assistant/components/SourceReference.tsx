import type { Snippet } from "../types";

interface Props { snippet: Snippet }

export function SourceReference({ snippet }: Props) {
  return (
    <div className="rounded border border-gray-200 bg-gray-50 px-3 py-2 text-xs">
      <span className="font-medium text-gray-700">{snippet.filePath}</span>
      {snippet.symbol && (
        <span className="ml-2 text-gray-500">→ {snippet.symbol}</span>
      )}
    </div>
  );
}
