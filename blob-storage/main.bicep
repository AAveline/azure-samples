param uuid string = newGuid()
param location string = resourceGroup().location

var storageAccountName = uniqueString(uuid)

resource storageAccount 'Microsoft.Storage/storageAccounts@2021-08-01' = {
  name: storageAccountName
  kind: 'StorageV2'
  location: location
  sku: {
    name: 'Standard_RAGRS'
  }
}

resource blobService 'Microsoft.Storage/storageAccounts/blobServices@2021-08-01' = {
  name: '${storageAccountName}/default'
  dependsOn: [
    storageAccount
  ]
}

resource blobContainer 'Microsoft.Storage/storageAccounts/blobServices/containers@2021-08-01' = {
  name: '${storageAccountName}/default/$web'
  dependsOn: [
    storageAccount
    blobService
  ]
}

resource contributorRoleDefinition 'Microsoft.Authorization/roleDefinitions@2018-01-01-preview' existing = {
  scope: subscription()
  name: '17d1049b-9a84-46fb-8f53-869881c3d3ab'
}

resource managedIdentity 'Microsoft.ManagedIdentity/userAssignedIdentities@2018-11-30' = {
  name: 'DeploymentScript'
  location: location
}

resource roleAssignment 'Microsoft.Authorization/roleAssignments@2020-04-01-preview' = {
  scope: storageAccount
  name: guid(resourceGroup().id, managedIdentity.id, contributorRoleDefinition.id)
  properties: {
    roleDefinitionId: contributorRoleDefinition.id
    principalId: managedIdentity.properties.principalId
    principalType: 'ServicePrincipal'
  }
}

resource deployment 'Microsoft.Resources/deploymentScripts@2020-10-01' = {
  name: 'deploymentScript'
  location: location
  kind: 'AzureCLI'

  dependsOn: [
    storageAccount
    blobContainer
  ]

  properties: {
    azCliVersion: '2.35.0'
    retentionInterval: 'PT1H'
    environmentVariables: [
      {
        name: 'storageAccountName'
        value: storageAccountName
      }
    ]
    scriptContent: '''
      az login --service-principal -u <sp-name> -p <sp-password> --tenant <sp-tenant>
      az storage blob service-properties update --account-name \${Env:storageAccountName} --static-website --index-document index.html
      az storage blob upload --file index.html --container-name \$web --account-name \${Env:storageAccountName}
    '''
  }
}

output websiteAddress string = storageAccount.properties.primaryEndpoints.web
