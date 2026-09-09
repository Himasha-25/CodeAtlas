"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { registerSchema, type RegisterSchema } from "../schemas";
import { useRegister } from "../hooks/useAuth";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Form, FormField, FormLabel, FormMessage } from "@/components/ui/form";

export function RegisterForm() {
  const { mutate, isPending, error } = useRegister();
  const { register, handleSubmit, formState: { errors } } = useForm<RegisterSchema>({
    resolver: zodResolver(registerSchema),
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
        {isPending ? "Creating account…" : "Create account"}
      </Button>
    </Form>
  );
}
