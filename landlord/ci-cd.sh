#!/bin/bash
set -e

DOCKER_USER="<your_dockerhub_username>"

echo "Starting Minikube..."
minikube start

echo "Pulling latest images..."
docker pull $DOCKER_USER/goapp:latest
docker pull $DOCKER_USER/postgres-custom:latest

echo "Deploying to Minikube..."
kubectl apply -f landlord/postgres.yaml
kubectl apply -f landlord/pgadmin.yaml
kubectl apply -f landlord/deployment.yaml

echo "Application deployed! Access at:"
minikube service goapp-service
