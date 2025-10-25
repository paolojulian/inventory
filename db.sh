#!/bin/bash

# Check if an argument is provided
if [ "$1" = "test" ]; then
    # Connect to test database
    docker exec -it server-db-test-1 psql -U postgres -d inventory_test
else
    # Connect to main database
    docker exec -it server-db-1 psql -U postgres -d inventory
fi
