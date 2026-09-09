import { apiClient } from "@/lib/api-client";
import type { Project, CreateProjectRequest, UpdateProjectRequest } from "./types";

export const projectsApi = {
  list: () =>
    apiClient.get<Project[]>("/api/v1/projects"),

  get: (id: number) =>
    apiClient.get<Project>(`/api/v1/projects/${id}`),

  create: (body: CreateProjectRequest) =>
    apiClient.post<Project>("/api/v1/projects", body),

  update: (id: number, body: UpdateProjectRequest) =>
    apiClient.patch<Project>(`/api/v1/projects/${id}`, body),

  delete: (id: number) =>
    apiClient.delete<void>(`/api/v1/projects/${id}`),
};
