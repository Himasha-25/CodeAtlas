# Database

PostgreSQL via GORM. Migrations in `apps/api/migrations/`.

## Schema Overview

```
users → projects → repositories → analysis_runs
                                       ├── files → symbols
                                       ├── dependencies
                                       └── api_endpoints

projects → documentation
projects → conversations → messages
```

## Key Tables

- `analysis_runs` — tracks status (`pending → running → completed/failed`); re-analysis creates a new run, preserving history.
- `dependencies` — single edge table with `from_id`, `to_id`, `edge_type`, `run_id`. Covers file→file, symbol→symbol, component→api.
- `symbols` — functions, classes, interfaces extracted per file per run.
