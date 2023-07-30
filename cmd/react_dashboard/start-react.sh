#!/bin/bash

if [[ -f "./keys/cert.pem" && -f "./keys/key.pem" ]]; then
  echo "Starting React app with SSL certificates..."
  HTTPS=true SSL_CRT_FILE=./keys/cert.pem SSL_KEY_FILE=./keys/key.pem npm run start-safe
else
  echo "Starting React app without SSL certificates..."
  npm run start
fi
