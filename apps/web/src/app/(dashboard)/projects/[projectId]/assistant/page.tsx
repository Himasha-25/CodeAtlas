type Props = { params: Promise<{ projectId: string }> };

export default async function AssistantPage({ params }: Props) {
  const { projectId } = await params;
  return (
    <div className="flex h-full flex-col">
      <h1 className="mb-4 text-2xl font-semibold">AI Assistant</h1>
      <p className="text-sm text-gray-500">Project {projectId} — ask questions about your codebase.</p>
    </div>
  );
}
