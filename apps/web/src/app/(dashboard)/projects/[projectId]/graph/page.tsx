type Props = { params: Promise<{ projectId: string }> };

export default async function GraphPage({ params }: Props) {
  const { projectId } = await params;
  return (
    <div>
      <h1 className="mb-4 text-2xl font-semibold">Dependency Graph</h1>
      <p className="text-sm text-gray-500">Project {projectId} — visualise file and symbol dependencies.</p>
    </div>
  );
}
