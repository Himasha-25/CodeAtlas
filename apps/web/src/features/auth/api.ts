import { apiClient } from "@/lib/api-client";
import type { AuthResponse, LoginRequest, RegisterRequest } from "./types";

export const authApi = {
  register: (body: RegisterRequest) =>
    apiClient.post<AuthResponse>("/api/v1/auth/register", body),

  login: (body: LoginRequest) =>
    apiClient.post<AuthResponse>("/api/v1/auth/login", body),
};
