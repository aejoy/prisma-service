#!/bin/sh

databases=(db0 db1 db2 db3)

for db in "${databases[@]}"; do
  pg-isready -U "$POSTGRES_USER" -d "$db" -h 127.0.0.1 -p 5432
  if [ $? -ne 0 ]; then
    echo "Database $db is not ready"
    exit 1
  fi
done

echo "All databases are ready"
exit 0