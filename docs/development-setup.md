# Development Setup

## Prerequisites

- Go 1.23+
- Node.js 20+
- Docker & Docker Compose
- Ollama (optional for AI features)

## Start

```bash
# 1. Clone and configure
cp .env.example .env

# 2. Start Postgres (and Ollama if needed)
docker-compose -f infra/docker/docker-compose.dev.yml up -d

# 3. Run migrations
cd apps/api && go run cmd/server/main.go migrate

# 4. Start API
cd apps/api && go run cmd/server/main.go

# 5. Start frontend
cd apps/web && npm install && npm run dev
```

## Build Order

See `docs/architecture.md` — build auth + projects end-to-end before touching the analyzer.
