import { useCallback, useState } from "react";

let _collapsed = false;
const _listeners = new Set<(v: boolean) => void>();

export function useSidebarStore() {
  const [collapsed, setCollapsed] = useState(_collapsed);

  useState(() => {
    _listeners.add(setCollapsed);
    return () => { _listeners.delete(setCollapsed); };
  });

  const toggle = useCallback(() => {
    _collapsed = !_collapsed;
    _listeners.forEach((fn) => fn(_collapsed));
  }, []);

  return { collapsed, toggle };
}
