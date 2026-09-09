type Props = { params: Promise<{ projectId: string }> };

export default async function ImpactPage({ params }: Props) {
  const { projectId } = await params;
  return (
    <div>
      <h1 className="mb-4 text-2xl font-semibold">Impact Analysis</h1>
      <p className="text-sm text-gray-500">Project {projectId} — trace what a change would affect.</p>
    </div>
  );
}
