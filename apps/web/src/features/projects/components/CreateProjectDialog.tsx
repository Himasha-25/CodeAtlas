"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { createProjectSchema, type CreateProjectSchema } from "../schemas";
import { useCreateProject } from "../hooks/useProjects";
import { Dialog, DialogHeader, DialogTitle, DialogContent, DialogFooter } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Form, FormField, FormLabel, FormMessage } from "@/components/ui/form";

interface Props {
  open: boolean;
  onClose: () => void;
}

export function CreateProjectDialog({ open, onClose }: Props) {
  const { mutate, isPending } = useCreateProject();
  const { register, handleSubmit, reset, formState: { errors } } = useForm<CreateProjectSchema>({
    resolver: zodResolver(createProjectSchema),
  });

  function onSubmit(data: CreateProjectSchema) {
    mutate(data, { onSuccess: () => { reset(); onClose(); } });
  }

  return (
    <Dialog open={open} onClose={onClose}>
      <DialogHeader>
        <DialogTitle>New project</DialogTitle>
      </DialogHeader>
      <DialogContent>
        <Form onSubmit={handleSubmit(onSubmit)} className="space-y-3">
          <FormField>
            <FormLabel htmlFor="name">Name</FormLabel>
            <Input id="name" {...register("name")} />
            {errors.name && <FormMessage>{errors.name.message}</FormMessage>}
          </FormField>
          <FormField>
            <FormLabel htmlFor="description">Description</FormLabel>
            <Input id="description" {...register("description")} />
          </FormField>
        </Form>
      </DialogContent>
      <DialogFooter>
        <Button variant="outline" onClick={onClose}>Cancel</Button>
        <Button disabled={isPending} onClick={handleSubmit(onSubmit)}>
          {isPending ? "Creating…" : "Create"}
        </Button>
      </DialogFooter>
    </Dialog>
  );
}
