#!/bin/bash
set -e
echo "Seeding database..."
psql "$DATABASE_URL" -f "$(dirname "$0")/../../apps/api/migrations/0001_init.sql"
echo "Done."
