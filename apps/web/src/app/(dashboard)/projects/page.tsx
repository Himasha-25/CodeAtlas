export default function ProjectsPage() {
  return (
    <div>
      <div className="mb-6 flex items-center justify-between">
        <h1 className="text-2xl font-semibold">Projects</h1>
        <button className="rounded bg-black px-4 py-2 text-sm text-white">New project</button>
      </div>
      <p className="text-sm text-gray-500">No projects yet.</p>
    </div>
  );
}
