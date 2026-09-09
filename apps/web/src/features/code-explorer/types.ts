export interface CodeFile {
  id: number;
  runId: number;
  path: string;
  language: string;
  lineCount: number;
}

export interface Symbol {
  id: number;
  fileId: number;
  runId: number;
  name: string;
  kind: "function" | "class" | "interface" | "variable";
  line: number;
  exported: boolean;
}
