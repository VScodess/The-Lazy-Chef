#!/bin/sh

# Run the data loading script
echo "Loading initial data into MongoDB..."
go run backend/scripts/loadData.go

# After loading data, start the main application
echo "Starting the main application..."
./main
