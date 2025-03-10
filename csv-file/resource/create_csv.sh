#!/bin/bash

# File size (MB)
file_size=5

# Output file
output_file="/tmp/large_users.csv"

# Target size in bytes
target_size=$(($file_size * 1024 * 1024))

echo "Creating csv file of $target_size MB"

# Write the header
echo "uuid,email,name,phone" > $output_file

# Create rows until the file reaches the target size
row_count=0
while [[ $(stat -c%s "$output_file") -lt $target_size ]]; do
  printf "uuid-%08d,user%d@example.com,User %d,123-456-%04d\n" \
    $row_count $row_count $row_count $((7890 + row_count % 10000)) >> $output_file
  ((row_count++))
done

echo "Created $output_file with size $(stat -c%s "$output_file") bytes and $row_count rows."
