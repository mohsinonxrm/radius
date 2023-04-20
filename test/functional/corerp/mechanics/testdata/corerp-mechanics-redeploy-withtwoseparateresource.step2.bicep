import radius as radius

@description('Specifies the location for resources.')
param location string = 'global'

@description('Specifies the environment for resources.')
param environment string = 'test'

@description('Specifies the image to be deployed.')
param magpieimage string

resource app 'Applications.Core/applications@2023-04-15-preview' = {
  name: 'corerp-mechanics-redeploy-withtwoseparateresource'
  location: location
  properties: {
    environment: environment
  }
}

resource mechanicsf 'Applications.Core/containers@2023-04-15-preview' = {
  name: 'mechanicsf'
  location: location
  properties: {
    application: app.id
    container: {
      image: magpieimage
    }
  }
}
