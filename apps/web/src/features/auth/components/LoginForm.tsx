"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { loginSchema, type LoginSchema } from "../schemas";
import { useLogin } from "../hooks/useAuth";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Form, FormField, FormLabel, FormMessage } from "@/components/ui/form";

export function LoginForm() {
  const { mutate, isPending, error } = useLogin();
  const { register, handleSubmit, formState: { errors } } = useForm<LoginSchema>({
    resolver: zodResolver(loginSchema),
  });

  return (
    <Form onSubmit={handleSubmit((data) => mutate(data))}>
      <FormField>
        <FormLabel htmlFor="email">Email</FormLabel>
        <Input id="email" type="email" {...register("email")} />
        {errors.email && <FormMessage>{errors.email.message}</FormMessage>}
      </FormField>
      <FormField>
        <FormLabel htmlFor="password">Password</FormLabel>
        <Input id="password" type="password" {...register("password")} />
        {errors.password && <FormMessage>{errors.password.message}</FormMessage>}
      </FormField>
      {error && <FormMessage>{error.message}</FormMessage>}
      <Button type="submit" className="w-full" disabled={isPending}>
        {isPending ? "Signing in…" : "Sign in"}
      </Button>
    </Form>
  );
}
