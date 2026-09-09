export default function RegisterPage() {
  return (
    <div className="rounded-lg border bg-white p-8 shadow-sm">
      <h1 className="mb-6 text-2xl font-semibold">Create account</h1>
      <form className="space-y-4">
        <div>
          <label className="mb-1 block text-sm font-medium">Email</label>
          <input type="email" className="w-full rounded border px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="mb-1 block text-sm font-medium">Password</label>
          <input type="password" className="w-full rounded border px-3 py-2 text-sm" />
        </div>
        <button type="submit" className="w-full rounded bg-black py-2 text-sm text-white">
          Create account
        </button>
      </form>
    </div>
  );
}
