#!/bin/bash

# Check if version argument is provided
if [ -z "$1" ]; then
    echo "Usage: ./fix-migration.sh <version_number>"
    echo "Example: ./fix-migration.sh 7"
    exit 1
fi

VERSION=$1

# Fix main database
echo "Fixing main database..."
docker exec -i server-db-1 psql -U postgres -d inventory -c "UPDATE schema_migrations SET version=$VERSION, dirty=false;"

# Fix test database
echo "Fixing test database..."
docker exec -i server-db-test-1 psql -U postgres -d inventory_test -c "UPDATE schema_migrations SET version=$VERSION, dirty=false;"

echo "Migration version set to $VERSION and dirty flag cleared."
