#!/bin/bash

# Create resources group
rg_name="ecommerce-rg"
location="westeurope"
container_registry="index.docker.io"
container_registry_username="aaveline"

az group create -g $rg_name -l $location

# Create deployment group
az deployment group create -f ./main.bicep -g $rg_name \
  -p \
    containerRegistry=$container_registry \
    containerRegistryUsername=$container_registry_username 