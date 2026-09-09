type Props = { params: Promise<{ projectId: string }> };

export default async function ProjectPage({ params }: Props) {
  const { projectId } = await params;
  return (
    <div>
      <h1 className="mb-4 text-2xl font-semibold">Project {projectId}</h1>
      <p className="text-sm text-gray-500">Overview of analysis runs and project status.</p>
    </div>
  );
}
