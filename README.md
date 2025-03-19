# Flexcity TSO

## Introduction

This micro-service is responsible for managing the TSO (Transmission System Operator) data. It is responsible to manage the activation of the assets.

## How to run

### Environment Variables

The following environment variables are required to run the service:
- `HTTP_PORT`: The port where the service will be listening to HTTP requests.

## Local

1. Install go 1.24, link [here](https://go.dev/dl/)
2. Clone the project and open a terminal in the folder
3. Copy the file `.env.example` to the `.env`
4. Run the application with `go run main.go`

### Curl example requests
Activation
```bash
curl --request POST \
  --url http://localhost:3000/v1/assets/activation \
  --header 'content-type: application/json' \
  --data '{
  "date": "2025-03-24T10:00:00Z",
  "volume": 3001
}'
```

Heartbeat
```bash
curl --request GET --url http://localhost:3000/ping
```

## Kubernetes

Use the helm chart (**TODO, theorical**) to deploy the service on a Kubernetes cluster.  
Define the environment variables in the `values.yaml` file.

## Planned Features

- Access control
- Better error management
- External storage like a database
- E2E Test
- Tracing
- Monitoring
- Helm chart
- Update CI: code linter, vulnerability check, publish image
