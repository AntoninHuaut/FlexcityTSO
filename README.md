# Flexcity TSO

## Introduction

This micro-service is responsible for managing the TSO (Transmission System Operator) data. It is responsible to manage the activation of the assets.

## How to run

### Environment Variables

The following environment variables are required to run the service:
- `HTTP_PORT`: The port where the service will be listening to HTTP requests.

## Charts

Use the helm chart (**TODO**) to deploy the service on a Kubernetes cluster.  
Define the environment variables in the `values.yaml` file.

## Planned Features

- Access control
- Better error management
- External storage like a database
- E2E Test
- Tracing
- Monitoring
- Helm chart
- Update CI: Added code linter, vulnerability check, publish image
