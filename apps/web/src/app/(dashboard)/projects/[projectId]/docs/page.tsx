type Props = { params: Promise<{ projectId: string }> };

export default async function DocsPage({ params }: Props) {
  const { projectId } = await params;
  return (
    <div>
      <h1 className="mb-4 text-2xl font-semibold">Documentation</h1>
      <p className="text-sm text-gray-500">Project {projectId} — AI-generated codebase documentation.</p>
    </div>
  );
}
