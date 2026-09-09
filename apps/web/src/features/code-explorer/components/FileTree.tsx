"use client";

import { useFiles } from "../hooks/useCodeExplorer";
import type { CodeFile } from "../types";

interface Props {
  runId: number;
  selectedFileId?: number;
  onSelect: (file: CodeFile) => void;
}

export function FileTree({ runId, selectedFileId, onSelect }: Props) {
  const { data: files, isLoading } = useFiles(runId);

  if (isLoading) return <p className="text-sm text-gray-400">Loading files…</p>;
  if (!files?.length) return <p className="text-sm text-gray-400">No files found.</p>;

  return (
    <ul className="space-y-0.5 text-sm">
      {files.map((file) => (
        <li key={file.id}>
          <button
            onClick={() => onSelect(file)}
            className={`w-full truncate rounded px-2 py-1 text-left hover:bg-gray-100 ${
              file.id === selectedFileId ? "bg-gray-100 font-medium" : ""
            }`}
          >
            {file.path}
          </button>
        </li>
      ))}
    </ul>
  );
}
