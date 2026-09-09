export type EdgeType = "file" | "symbol" | "api";

export interface Dependency {
  id: number;
  runId: number;
  fromId: number;
  toId: number;
  edgeType: EdgeType;
}

export interface GraphData {
  nodes: { id: string; label: string }[];
  edges: { id: string; source: string; target: string; label: EdgeType }[];
}
