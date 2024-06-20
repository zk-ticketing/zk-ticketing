# ZK Ticketing

## Overview

This is the monorepo for the ZK Ticketing system. This project aim to:

1. Bridge gap of email-based ticketing and ZK ticket verification.
2. Allow users to check-in to events without revealing any personal information.

## Project Structure

### Frontend

The frontend is a Next.js application that provides the user interface for the ZK Ticketing system.

### Backend

The backend is a Go-based HTTP server that handles the server-side logic for the application.

### OpenAPI

The openapi folder contains the OpenAPI specification for the ZK Ticketing system.

### Infrastructure

The infra folder contains Kubernetes manifests for deploying the application.

## Development

### Prerequisites

- docker
- kubectl
- minikube
- [skaffold](http://skaffold.dev)
- [openapi-generator](https://github.com/OpenAPITools/openapi-generator)

## Running the Application Locally

### Use minikube to start local cluster

```
minikube start --profile zk-ticketing
skaffold config set --global local-cluster true
eval $(minikube -p zk-ticketing docker-env)
```

### Run skaffold in local development mode

```
skaffold dev
```

### Access the application

Port-forward backend and frontend (since they are using LoadBalancer type service):

```
minikube tunnel -p zk-ticketing
```

Once the tunnel is running, you can access the application at:

- Frontend: http://localhost:8080
- Backend: http://localhost:3000

Port-forward postgres db:

```
kubectl port-forward svc/postgres 5432:5432
```

Once the tunnel is running, you can access db using any postgres client (e.g. pgAdmin, DBeaver, etc.) with the following credentials (for local development only:

- Host: localhost
- Port: 5432
- Username: postgres
- Password: password

### Open Swagger UI to test the API

In a new terminal, run the following command:

```
start-swagger-ui
```

Then open the Swagger UI at http://localhost:8081 to test the API against our forwarded backend port.
