param location string
param containerRegistry string
param containerRegistryUsername string
param appContainerEnvironmentName string

var logAnalyticsWorkspaceName = 'logs-${appContainerEnvironmentName}'
var appInsightsName = 'appins-${appContainerEnvironmentName}'

resource logAnalyticsWorkspace 'Microsoft.OperationalInsights/workspaces@2020-03-01-preview' = {
  name: logAnalyticsWorkspaceName
  location: location
  properties: any({
    retentionInDays: 30
    features: {
      searchVersion: 1
    }
    sku: {
      name: 'PerGB2018'
    }
  })
}

resource appInsights 'Microsoft.Insights/components@2020-02-02' = {
  name: appInsightsName
  location: location
  kind: 'web'
  properties: {
    Application_Type: 'web'
    WorkspaceResourceId: logAnalyticsWorkspace.id
  }
}

resource appContainerEnvironment 'Microsoft.App/managedEnvironments@2022-01-01-preview' = {
  name: appContainerEnvironmentName
  location: location
  properties: {
    daprAIInstrumentationKey: reference(appInsights.id, '2020-02-02').InstrumentationKey
    appLogsConfiguration: {
      destination: 'log-analytics'
      logAnalyticsConfiguration: {
        customerId: reference(logAnalyticsWorkspace.id, '2020-03-01-preview').customerId
        sharedKey: listKeys(logAnalyticsWorkspace.id, '2020-03-01-preview').primarySharedKey
      }
    }
  }
}

resource apiGateway 'Microsoft.App/containerApps@2022-01-01-preview' = {
  name: 'api-gateway'
  location: location
  properties: {
    managedEnvironmentId: appContainerEnvironment.id
    configuration: {
      ingress: {
        external: true
        targetPort: 3005
        transport: 'auto'
      }
      dapr: {
        enabled: true
        appPort: 3005
        appId: 'api-gateway'
      }
    }
    template: {
      containers: [
        {
          image: '${containerRegistry}/${containerRegistryUsername}/api-gateway:latest'
          name: 'api-gateway'
        }
      ]
      scale: {
        minReplicas: 1
      }
    }
  }
}

resource productSercice 'Microsoft.App/containerApps@2022-01-01-preview' = {
  name: 'product-service'
  location: location
  properties: {
    managedEnvironmentId: appContainerEnvironment.id
    configuration: {
      dapr: {
        enabled: true
        appPort: 3000
        appProtocol: 'http'
        appId: 'product-service'
      }
    }
    template: {
      containers: [
        {
          image: '${containerRegistry}/${containerRegistryUsername}/product-service:latest'
          name: 'product-service'
        }
      ]
      scale: {
        minReplicas: 1
      }
    }
  }
}

resource orderService 'Microsoft.App/containerApps@2022-01-01-preview' = {
  name: 'order-service'
  location: location
  properties: {
    managedEnvironmentId: appContainerEnvironment.id
    configuration: {
      dapr: {
        enabled: true
        appPort: 3001
        appId: 'order-service'
      }
    }
    template: {
      containers: [
        {
          image: '${containerRegistry}/${containerRegistryUsername}/order-service:latest'
          name: 'order-service'
        }
      ]
      scale: {
        minReplicas: 1
      }
    }
  }
}

resource userService 'Microsoft.App/containerApps@2022-01-01-preview' = {
  name: 'user-service'
  location: location
  properties: {
    managedEnvironmentId: appContainerEnvironment.id
    configuration: {
      dapr: {
        enabled: true
        appPort: 3002
        appId: 'user-service'
      }
    }
    template: {
      containers: [
        {
          image: '${containerRegistry}/${containerRegistryUsername}/user-service:latest'
          name: 'user-service'
        }
      ]
      scale: {
        minReplicas: 1
      }
    }
  }
}
