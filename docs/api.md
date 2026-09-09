# API

REST API served by Go/Gin on port 8080.

## Auth

All routes except `/api/v1/auth/*` require `Authorization: Bearer <token>`.

## Route Groups

| Prefix | Domain |
|--------|--------|
| `/api/v1/auth` | register, login |
| `/api/v1/projects` | project CRUD |
| `/api/v1/projects/:id/repositories` | repo upload, status |
| `/api/v1/projects/:id/analysis` | trigger, poll status |
| `/api/v1/projects/:id/explorer` | files, symbols |
| `/api/v1/projects/:id/graph` | dependency edges |
| `/api/v1/projects/:id/docs` | generated documentation |
| `/api/v1/projects/:id/assistant` | AI chat |
| `/api/v1/projects/:id/impact` | impact analysis |
