param location string = resourceGroup().location
param containerRegistry string
param containerRegistryUsername string
param appContainerEnvironmentName string = 'daprappregistry'

module containerModule 'templates/containerapps.bicep' = {
  name: 'containerapps'
  params: {
    location: location
    containerRegistry: containerRegistry
    containerRegistryUsername: containerRegistryUsername
    appContainerEnvironmentName: appContainerEnvironmentName
  }
}
