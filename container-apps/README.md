# Ecommerce Project on Azure ContainerApps

## Overview

This project implements Microservice Pattern with Azure ContainerApps. It uses:

- Azure ContainerApps
- DAPR (for services mesh)
- Go (backend implementation)
- GraphQL (for ApiGateway)
- ReactJS (frontend)
- MongoDB Cloud (persistent storage)

## Considerations

Each microservice own its application database.

- [Application database pattern](https://martinfowler.com/bliki/ApplicationDatabase.html)
- [Serverless database](https://www.mongodb.com/cloud/atlas/serverless)

It's not a production-ready project

## How it works

- Frontal graphql Api-Gateway for backend services
- DAPR to enable inter-service communication
- Docker compose is used in development mode, with two meaning:

  - Dev simplicity
  - Replicate as much as possible production env

## Setup

- Local or remote installation of Azure CLI
- Setup 3 MongoDB databases (by service) : (MongoDB cloud serverless db can be used)
- Start project:
  - Backend: `cd backend && docker-compose up`
  - Frontend: `cd frontend && npm start`
  - Deployment: `cd modules && chmod +x deployment.sh && ./deployment.sh`
