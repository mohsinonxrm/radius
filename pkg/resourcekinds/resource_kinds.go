// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package resourcekinds

// ResourceKinds supported. The RP determines how these are created/deleted and the HealthService determines how
// health checks are handled for these
const (
	Azure                            = "azure"
	Kubernetes                       = "kubernetes"
	Deployment                       = "Deployment"
	Service                          = "Service"
	Secret                           = "Secret"
	Gateway                          = "Gateway"
	KubernetesHTTPRoute              = "HTTPRoute" // For httproutes.networking.x-k8s.io
	RadiusHttpRoute                  = "HttpRoute" // For httproutes.radius.dev
	StatefulSet                      = "StatefulSet"
	SecretProviderClass              = "SecretProviderClass"
	DaprStateStoreAzureStorage       = "dapr.statestore.azurestorage"
	DaprGeneric                      = "generic"
	DaprComponent                    = "dapr.io.component"
	DaprPubSubTopicAzureServiceBus   = "pubsub.azure.servicebus"
	AzureCosmosAccount               = "azure.cosmosdb.account"
	AzureCosmosDBMongo               = "azure.cosmosdb.mongo"
	AzureCosmosDBSQL                 = "azure.cosmosdb.sql"
	AzureServiceBusQueue             = "azure.servicebus.queue"
	AzureSqlServer                   = "azure.sql"
	AzureSqlServerDatabase           = "azure.sql.database"
	AzureKeyVault                    = "azure.keyvault"
	AzureKeyVaultSecret              = "azure.keyvault.secret"
	AzurePodIdentity                 = "azure.aadpodidentity"
	AzureUserAssignedManagedIdentity = "azure.userassignedmanagedidentity"
	AzureRoleAssignment              = "azure.roleassignment"
	AzureRedis                       = "azure.redis"
	AzureFileShare                   = "azure.fileshare"
	AzureFileShareStorageAccount     = "azure.fileshare.account"
)
