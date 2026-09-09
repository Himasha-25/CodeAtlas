// Shared types between apps/web and any future consumers.
// Only add here if a type is used on both sides of the API boundary
// (e.g. generated from OpenAPI). Hand-write feature types in features/*/types.ts instead.

export interface ApiError {
  error: string;
}

export type AnalysisStatus = "pending" | "running" | "completed" | "failed";

export type EdgeType = "file" | "symbol" | "api";

export type SymbolKind = "function" | "class" | "interface" | "variable";
