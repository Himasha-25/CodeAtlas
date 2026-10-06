import { LoginForm } from "@/features/auth/components/LoginForm";
import Link from "next/link";

export default function LoginPage() {
  return (
    <div className="rounded-lg border bg-white p-8 shadow-sm">
      <h1 className="mb-6 text-2xl font-semibold">Sign in</h1>
      <LoginForm />
      <p className="mt-4 text-center text-sm text-gray-500">
        No account?{" "}
        <Link href="/register" className="font-medium underline">
          Register
        </Link>
      </p>
    </div>
  );
}
