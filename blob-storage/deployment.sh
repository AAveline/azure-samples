#!/bin/bash

# Get your subscriptionID
subscription_id=$(az account show --query id --output tsv)
# Create RBAC role for your subscription
az ad sp create-for-rbac -n "rbac-for-storage" --role Contributor --scopes /subscriptions/$subscription_id

# Create resources group
rg_name="storage-rg"
location="westeurope"

az group create -g $rg_name -l $location

# Create deployment group
az deployment group create -f ./main.bicep -g $rg_name 