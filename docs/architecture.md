# Architecture

CodeAtlas is a modular monolith: one frontend, one backend, one database, one local LLM.

## Units

| Unit | Tech | Role |
|------|------|------|
| `apps/web` | Next.js, TypeScript | UI |
| `apps/api` | Go, Gin, GORM | REST API + analyzer |
| PostgreSQL | — | Persistence |
| Ollama | — | Local LLM |

## Key Decisions

- **Analyzer is a Go package** inside `apps/api`, not a separate service. Isolated by code boundaries, not network boundaries.
- **Background jobs** use goroutines + `analysis_runs.status` column. No queue needed at this scale.
- **RAG retriever** is behind an interface — v1 uses Postgres full-text search, swappable to pgvector later.
- **LLM client** is behind an interface — `ollama_client.go` today, swappable to OpenAI without touching service logic.

## What Would Change at Scale

- Analyzer → separate worker service with a real job queue (e.g., SQS)
- Retriever → pgvector or dedicated search engine
- Auth → dedicated identity provider
