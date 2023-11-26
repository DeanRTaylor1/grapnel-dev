#!/bin/bash

# Define the API endpoint you want to test
API_ENDPOINT="http://localhost:8080/"

# Define the number of requests you want to send
NUM_REQUESTS=100

# Define the interval between requests in seconds
REQUEST_INTERVAL=1

# Loop to send requests
for ((i=1; i<=$NUM_REQUESTS; i++)); do
  # Send a GET request to the API endpoint using curl
  response=$(curl -si -X GET $API_ENDPOINT)
  
  # Extract and display the status code from the response headers
  status_code=$(echo "$response" | grep "HTTP/" | awk '{print $2}')
  echo "Request $i: Status Code - $status_code"
  
  # Sleep for the specified interval before sending the next request
#   sleep $REQUEST_INTERVAL
done
