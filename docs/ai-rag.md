# AI / RAG

## Flow

```
question + projectId
  → retriever: Postgres full-text search over symbols/files
  → context_builder: assembles bounded prompt (question + top-k snippets + paths)
  → groq_client: Groq cloud API inference (llama-3.3-70b-versatile)
  → response + cited file/symbol references saved to messages
```

## Interfaces

- `Retriever` — v1: `ILIKE`/full-text over `symbols`/`files`. Future: pgvector embeddings.
- `LLMClient` — v1: `groq_client.go` (Groq cloud API). The interface is intentionally thin (`Complete(prompt string) (string, error)`) so the provider can be swapped again (e.g. back to a local model) without touching `service.go` or `handler.go`.

Both are swappable without touching `assistant/service.go` or `assistant/handler.go`.

## Configuration

Set `GROQ_API_KEY` in your environment (or `.env`). There is no insecure default — the application will return an auth error from Groq if the key is absent when assistant features are used.

## Constraints

- Context window is bounded — top-k snippets only, never the whole repo.
- Responses are saved to `conversations`/`messages` with cited source references.
