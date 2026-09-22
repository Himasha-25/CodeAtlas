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
