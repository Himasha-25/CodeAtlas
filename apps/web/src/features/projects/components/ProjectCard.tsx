import Link from "next/link";
import { Card, CardHeader, CardTitle, CardDescription, CardFooter } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import type { Project } from "../types";

interface Props {
  project: Project;
  onDelete: (id: number) => void;
}

export function ProjectCard({ project, onDelete }: Props) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>
          <Link href={`/projects/${project.id}`} className="hover:underline">
            {project.name}
          </Link>
        </CardTitle>
        {project.description && (
          <CardDescription>{project.description}</CardDescription>
        )}
      </CardHeader>
      <CardFooter className="gap-2">
        <Button variant="outline" size="sm" asChild>
          <Link href={`/projects/${project.id}`}>Open</Link>
        </Button>
        <Button variant="destructive" size="sm" onClick={() => onDelete(project.id)}>
          Delete
        </Button>
      </CardFooter>
    </Card>
  );
}
