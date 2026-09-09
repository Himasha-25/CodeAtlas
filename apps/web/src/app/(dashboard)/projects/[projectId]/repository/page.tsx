type Props = { params: Promise<{ projectId: string }> };

export default async function RepositoryPage({ params }: Props) {
  const { projectId } = await params;
  return (
    <div>
      <h1 className="mb-4 text-2xl font-semibold">Repository</h1>
      <p className="mb-4 text-sm text-gray-500">Project {projectId} — upload a ZIP to analyse.</p>
      <form className="space-y-4">
        <input type="file" accept=".zip" className="block text-sm" />
        <button type="submit" className="rounded bg-black px-4 py-2 text-sm text-white">
          Upload & analyse
        </button>
      </form>
    </div>
  );
}
