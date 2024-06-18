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

### Infrastructure

The infra folder contains Kubernetes manifests for deploying the application.

## Local Development

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

```
minikube tunnel -p zk-ticketing
```

Once the tunnel is running, you can access the application at:

- Frontend: http://localhost:8080
- Backend: http://localhost:3000
