#!/bin/bash
set -e
echo "Resetting database..."
psql "$DATABASE_URL" -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
echo "Done. Run seed.sh to re-seed."
