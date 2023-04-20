import radius as radius

param magpieimage string

param environment string

param location string = resourceGroup().location

param namespaceresourceid string

resource app 'Applications.Core/applications@2023-04-15-preview' = {
  name: 'corerp-resources-dapr-pubsub-servicebus'
  location: location
  properties: {
    environment: environment
  }
}

resource publisher 'Applications.Core/containers@2023-04-15-preview' = {
  name: 'sb-publisher'
  location: location
  properties: {
    application: app.id
    connections: {
      daprpubsub: {
        source: pubsub.id
      }
    }
    container: {
      image: magpieimage
      env: {
        BINDING_DAPRPUBSUB_NAME: pubsub.name
        BINDING_DAPRPUBSUB_TOPIC: pubsub.properties.topic
      }
      readinessProbe:{
        kind: 'httpGet'
        containerPort: 3000
        path: '/healthz'
      }
    }
    extensions: [
      {
        kind: 'daprSidecar'
        appId: 'sb-pubsub'
        appPort: 3000
      }
    ]
  }
}

resource pubsub 'Applications.Link/daprPubSubBrokers@2023-04-15-preview' = {
  name: 'sb-pubsub'
  location: location
  properties: {
    environment: environment
    application: app.id
    mode: 'resource'
    resource: namespaceresourceid
  }
}
