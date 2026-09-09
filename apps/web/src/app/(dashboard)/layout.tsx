export default function DashboardLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex min-h-screen">
      <nav className="w-64 border-r bg-white p-4">
        <p className="font-semibold text-lg">CodeAtlas</p>
      </nav>
      <main className="flex-1 p-6">{children}</main>
    </div>
  );
}
