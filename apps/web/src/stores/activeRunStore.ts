import { useCallback, useState } from "react";

const _map: Record<number, number> = {};
const _listeners = new Set<() => void>();

export function useActiveRunStore(projectId: number) {
  const [runId, setRunId] = useState<number>(_map[projectId] ?? 0);

  useState(() => {
    const sync = () => setRunId(_map[projectId] ?? 0);
    _listeners.add(sync);
    return () => { _listeners.delete(sync); };
  });

  const setActiveRun = useCallback((id: number) => {
    _map[projectId] = id;
    _listeners.forEach((fn) => fn());
  }, [projectId]);

  return { runId, setActiveRun };
}
