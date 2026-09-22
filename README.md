# CodeAtlas

AI-Powered Codebase Intelligence & Developer Onboarding Platform.

## Stack

- **Frontend:** Next.js (App Router), TypeScript, Tailwind CSS, shadcn/ui, TanStack Query
- **Backend:** Go, Gin, GORM
- **Database:** PostgreSQL
- **AI:** Groq (cloud LLM API), RAG over code symbols/files
- **Analyzer:** Tree-sitter AST parsing (Go package, no separate service)

## Quick Start

```bash
cp .env.example .env
docker compose -f infra/docker/docker-compose.dev.yml up -d
```

## Development

```bash
make dev        # start all services
make migrate    # run DB migrations
make test       # run all tests
```

See `docs/development-setup.md` for full setup instructions.

## CI Status

![CI API](https://github.com/Himasha-25/CodeAtlas/actions/workflows/ci-api.yml/badge.svg)
![CI Web](https://github.com/Himasha-25/CodeAtlas/actions/workflows/ci-web.yml/badge.svg)

## Phase 0 — Setup & Tooling

| Task | What was done |
|------|---------------|
| Monorepo structure | Top-level folders: `apps/web`, `apps/api`, `packages/shared-types`, `infra/docker`, `infra/scripts`, `docs`, `.github/workflows` |
| WSL2 + Docker | Ubuntu on WSL2, Docker installed inside WSL (no Docker Desktop), user added to docker group |
| Project on WSL filesystem | Repo lives at `~/projects/codeatlas` for fast I/O and reliable Docker volume mounts |
| Docker Compose (dev) | `infra/docker/docker-compose.dev.yml` runs Postgres on port 5432; Groq API used instead of Ollama |
| Yarn migration | Removed `package-lock.json`, switched `apps/web` to Yarn — `yarn.lock` committed, Dockerfile and docs updated |
| CI pipeline | `ci-api.yml` runs `go build/test/vet` with Go 1.23 + `GOTOOLCHAIN=local`; `ci-web.yml` caches via `yarn.lock` and runs lint/test/build |
| Branching workflow | Feature branches (`feat/`, `fix/`, `chore/`), PR per branch, CI must pass before merge; branch protection on `main` |
