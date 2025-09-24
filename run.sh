#!/bin/bash

set -e

echo “🛠️ Building images and launching services...”
docker-compose up --build -d

echo “⏳ Waiting for containers to start...”
sleep 5

echo “✅ Services up:”
docker ps --format “table {{.Names}}\t{{.Status}}\t{{.Ports}}”

echo “📡 RabbitMQ UI: http://localhost:15672 (guest/guest)”
echo “🎉 Consumer and Publisher are now running inside containers.”
