"use client";

import { useRef, type ChangeEvent } from "react";
import { useUploadRepository } from "../hooks/useRepositories";
import { Button } from "@/components/ui/button";

interface Props { projectId: number }

export function UploadZip({ projectId }: Props) {
  const inputRef = useRef<HTMLInputElement>(null);
  const { mutate, isPending, error, isSuccess } = useUploadRepository(projectId);

  function handleChange(e: ChangeEvent<HTMLInputElement>) {
    const file = e.target.files?.[0];
    if (file) mutate(file);
  }

  return (
    <div className="space-y-2">
      <input ref={inputRef} type="file" accept=".zip" className="hidden" onChange={handleChange} />
      <Button variant="outline" onClick={() => inputRef.current?.click()} disabled={isPending}>
        {isPending ? "Uploading…" : "Upload ZIP"}
      </Button>
      {isSuccess && <p className="text-sm text-green-600">Upload complete.</p>}
      {error && <p className="text-sm text-red-600">{error.message}</p>}
    </div>
  );
}
