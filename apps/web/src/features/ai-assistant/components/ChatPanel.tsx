"use client";

import { useState } from "react";
import { useAsk } from "../hooks/useAssistant";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { SourceReference } from "./SourceReference";
import type { AskResponse } from "../types";

interface Props { projectId: number; runId: number }

export function ChatPanel({ projectId, runId }: Props) {
  const [question, setQuestion] = useState("");
  const [history, setHistory] = useState<{ question: string; response: AskResponse }[]>([]);
  const { mutate, isPending } = useAsk(projectId);

  function submit() {
    if (!question.trim()) return;
    mutate(
      { runId, question },
      {
        onSuccess: (response) => {
          setHistory((h) => [...h, { question, response }]);
          setQuestion("");
        },
      },
    );
  }

  return (
    <div className="flex h-full flex-col gap-4">
      <div className="flex-1 space-y-4 overflow-y-auto">
        {history.map((item, i) => (
          <div key={i} className="space-y-2">
            <p className="text-sm font-medium">{item.question}</p>
            <p className="text-sm text-gray-700">{item.response.answer}</p>
            {item.response.sources.map((s, j) => (
              <SourceReference key={j} snippet={s} />
            ))}
          </div>
        ))}
        {!history.length && (
          <p className="text-sm text-gray-400">Ask a question about your codebase.</p>
        )}
      </div>
      <div className="flex gap-2">
        <Input
          value={question}
          onChange={(e) => setQuestion(e.target.value)}
          onKeyDown={(e) => e.key === "Enter" && submit()}
          placeholder="How does authentication work?"
          disabled={isPending}
        />
        <Button onClick={submit} disabled={isPending || !question.trim()}>
          {isPending ? "…" : "Ask"}
        </Button>
      </div>
    </div>
  );
}
