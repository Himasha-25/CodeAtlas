import { useMutation } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { authApi } from "../api";
import { setToken, clearToken } from "@/lib/auth";
import type { LoginRequest, RegisterRequest } from "../types";

export function useLogin() {
  const router = useRouter();
  return useMutation({
    mutationFn: (body: LoginRequest) => authApi.login(body),
    onSuccess: ({ token }) => {
      setToken(token);
      router.push("/dashboard");
    },
  });
}

export function useRegister() {
  const router = useRouter();
  return useMutation({
    mutationFn: (body: RegisterRequest) => authApi.register(body),
    onSuccess: ({ token }) => {
      setToken(token);
      router.push("/dashboard");
    },
  });
}

export function useLogout() {
  const router = useRouter();
  return () => {
    clearToken();
    router.push("/login");
  };
}
