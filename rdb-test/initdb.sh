#!/bin/bash

# Create table
for f in $(find /var/db/dump/ -name "*.sql")
do
   psql "$POSTGRES_DB" -U "$POSTGRES_USER" -f $f
done
