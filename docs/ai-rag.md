# AI / RAG

## Flow

```
question + projectId
  → retriever: Postgres full-text search over symbols/files
  → context_builder: assembles bounded prompt (question + top-k snippets + paths)
  → ollama_client: local Ollama inference
  → response + cited file/symbol references saved to messages
```

## Interfaces

- `Retriever` — v1: `ILIKE`/full-text over `symbols`/`files`. Future: pgvector embeddings.
- `LLMClient` — v1: `ollama_client.go`. Future: `openai_client.go` with same interface.

Both are swappable without touching `assistant/service.go` or `assistant/handler.go`.

## Constraints

- Context window is bounded — top-k snippets only, never the whole repo.
- Responses are saved to `conversations`/`messages` with cited source references.
