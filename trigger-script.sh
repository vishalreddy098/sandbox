#!/bin/bash
set -e

DOCKER_USER="vishalreddy11"

echo "Starting Minikube..."
minikube start

echo "Pulling latest images..."
# Optional for offline cache
docker pull $DOCKER_USER/goapp:latest
docker pull $DOCKER_USER/postgres-custom:latest

echo "Deploying to Minikube..."
kubectl delete deployment postgres --ignore-not-found
kubectl apply -f landlord/postgres.yaml
kubectl apply -f landlord/postgres-pv.yaml
kubectl apply -f landlord/pgadmin.yaml
kubectl apply -f landlord/deployment.yaml

echo "Application deployed! Access at:"
minikube service goapp-service
