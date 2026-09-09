import type { CodeFile } from "../types";

interface Props { file: CodeFile | null }

export function CodeViewer({ file }: Props) {
  if (!file) {
    return (
      <div className="flex h-full items-center justify-center text-sm text-gray-400">
        Select a file to view
      </div>
    );
  }

  return (
    <div className="h-full overflow-auto rounded border bg-gray-50 p-4">
      <p className="mb-2 text-xs text-gray-500">{file.path}</p>
      <p className="text-xs text-gray-400">
        {file.language} · {file.lineCount} lines
      </p>
    </div>
  );
}
