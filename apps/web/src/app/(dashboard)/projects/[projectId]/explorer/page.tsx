type Props = { params: Promise<{ projectId: string }> };

export default async function ExplorerPage({ params }: Props) {
  const { projectId } = await params;
  return (
    <div>
      <h1 className="mb-4 text-2xl font-semibold">Code Explorer</h1>
      <p className="text-sm text-gray-500">Project {projectId} — browse files and symbols.</p>
    </div>
  );
}
