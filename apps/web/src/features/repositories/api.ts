import { apiClient } from "@/lib/api-client";
import type { Repository } from "./types";

export const repositoriesApi = {
  list: (projectId: number) =>
    apiClient.get<Repository[]>(`/api/v1/projects/${projectId}/repositories`),

  upload: (projectId: number, file: File) => {
    const form = new FormData();
    form.append("file", file);
    return fetch(
      `${process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080"}/api/v1/projects/${projectId}/repositories`,
      {
        method: "POST",
        headers: {
          Authorization: `Bearer ${localStorage.getItem("codeatlas_token") ?? ""}`,
        },
        body: form,
      },
    ).then((r) => {
      if (!r.ok) throw new Error("Upload failed");
      return r.json() as Promise<Repository>;
    });
  },
};
