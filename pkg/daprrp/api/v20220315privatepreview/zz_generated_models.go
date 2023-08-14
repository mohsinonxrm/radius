//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import "time"

// BasicDaprResourceProperties - Basic properties of a Dapr component object.
type BasicDaprResourceProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the portable resource is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the portable resource is consumed by
	Application *string `json:"application,omitempty"`

	// READ-ONLY; The name of the Dapr component object. Use this value in your code when interacting with the Dapr client to
// use the Dapr component.
	ComponentName *string `json:"componentName,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// BasicResourceProperties - Basic properties of a Radius resource.
type BasicResourceProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the portable resource is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the portable resource is consumed by
	Application *string `json:"application,omitempty"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// DaprPubSubBrokerClientBeginDeleteOptions contains the optional parameters for the DaprPubSubBrokerClient.BeginDelete method.
type DaprPubSubBrokerClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DaprPubSubBrokerClientCreateOrUpdateOptions contains the optional parameters for the DaprPubSubBrokerClient.CreateOrUpdate
// method.
type DaprPubSubBrokerClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DaprPubSubBrokerClientGetOptions contains the optional parameters for the DaprPubSubBrokerClient.Get method.
type DaprPubSubBrokerClientGetOptions struct {
	// placeholder for future optional parameters
}

// DaprPubSubBrokerClientListByRootScopeOptions contains the optional parameters for the DaprPubSubBrokerClient.ListByRootScope
// method.
type DaprPubSubBrokerClientListByRootScopeOptions struct {
	// placeholder for future optional parameters
}

// DaprPubSubBrokerProperties - Dapr PubSubBroker portable resource properties
type DaprPubSubBrokerProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the portable resource is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the portable resource is consumed by
	Application *string `json:"application,omitempty"`

	// Metadata for the Dapr PubSubBroker resource. This should match the values specified in Dapr component spec
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The recipe used to automatically deploy underlying infrastructure for the Dapr PubSubBroker portable resource
	Recipe *Recipe `json:"recipe,omitempty"`

	// Specifies how the underlying service/resource is provisioned and managed.
	ResourceProvisioning *ResourceProvisioning `json:"resourceProvisioning,omitempty"`

	// A collection of references to resources associated with the Dapr PubSubBroker
	Resources []*ResourceReference `json:"resources,omitempty"`

	// Dapr PubSubBroker type. These strings match the format used by Dapr Kubernetes configuration format.
	Type *string `json:"type,omitempty"`

	// Dapr component version
	Version *string `json:"version,omitempty"`

	// READ-ONLY; The name of the Dapr component object. Use this value in your code when interacting with the Dapr client to
// use the Dapr component.
	ComponentName *string `json:"componentName,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the Dapr PubSubBroker portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// DaprPubSubBrokerResource - Dapr PubSubBroker portable resource
type DaprPubSubBrokerResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties *DaprPubSubBrokerProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DaprPubSubBrokerResourceListResult - The response of a DaprPubSubBrokerResource list operation.
type DaprPubSubBrokerResourceListResult struct {
	// REQUIRED; The DaprPubSubBrokerResource items on this page
	Value []*DaprPubSubBrokerResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// DaprSecretStoreClientCreateOrUpdateOptions contains the optional parameters for the DaprSecretStoreClient.CreateOrUpdate
// method.
type DaprSecretStoreClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DaprSecretStoreClientDeleteOptions contains the optional parameters for the DaprSecretStoreClient.Delete method.
type DaprSecretStoreClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// DaprSecretStoreClientGetOptions contains the optional parameters for the DaprSecretStoreClient.Get method.
type DaprSecretStoreClientGetOptions struct {
	// placeholder for future optional parameters
}

// DaprSecretStoreClientListByRootScopeOptions contains the optional parameters for the DaprSecretStoreClient.ListByRootScope
// method.
type DaprSecretStoreClientListByRootScopeOptions struct {
	// placeholder for future optional parameters
}

// DaprSecretStoreProperties - Dapr SecretStore portable resource properties
type DaprSecretStoreProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the portable resource is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the portable resource is consumed by
	Application *string `json:"application,omitempty"`

	// Metadata for the Dapr SecretStore resource. This should match the values specified in Dapr component spec
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The recipe used to automatically deploy underlying infrastructure for the Dapr SecretStore portable resource
	Recipe *Recipe `json:"recipe,omitempty"`

	// Specifies how the underlying service/resource is provisioned and managed.
	ResourceProvisioning *ResourceProvisioning `json:"resourceProvisioning,omitempty"`

	// Dapr SecretStore type. These strings match the types defined in Dapr Component format: https://docs.dapr.io/reference/components-reference/supported-secret-stores/
	Type *string `json:"type,omitempty"`

	// Dapr component version
	Version *string `json:"version,omitempty"`

	// READ-ONLY; The name of the Dapr component object. Use this value in your code when interacting with the Dapr client to
// use the Dapr component.
	ComponentName *string `json:"componentName,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the dapr secret store portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// DaprSecretStoreResource - Dapr SecretStore portable resource
type DaprSecretStoreResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties *DaprSecretStoreProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DaprSecretStoreResourceListResult - The response of a DaprSecretStoreResource list operation.
type DaprSecretStoreResourceListResult struct {
	// REQUIRED; The DaprSecretStoreResource items on this page
	Value []*DaprSecretStoreResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// DaprStateStoreClientBeginDeleteOptions contains the optional parameters for the DaprStateStoreClient.BeginDelete method.
type DaprStateStoreClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DaprStateStoreClientCreateOrUpdateOptions contains the optional parameters for the DaprStateStoreClient.CreateOrUpdate
// method.
type DaprStateStoreClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DaprStateStoreClientGetOptions contains the optional parameters for the DaprStateStoreClient.Get method.
type DaprStateStoreClientGetOptions struct {
	// placeholder for future optional parameters
}

// DaprStateStoreClientListByRootScopeOptions contains the optional parameters for the DaprStateStoreClient.ListByRootScope
// method.
type DaprStateStoreClientListByRootScopeOptions struct {
	// placeholder for future optional parameters
}

// DaprStateStoreProperties - Dapr StateStore portable resource properties
type DaprStateStoreProperties struct {
	// REQUIRED; Fully qualified resource ID for the environment that the portable resource is linked to
	Environment *string `json:"environment,omitempty"`

	// Fully qualified resource ID for the application that the portable resource is consumed by
	Application *string `json:"application,omitempty"`

	// Metadata for the Dapr StateStore resource. This should match the values specified in Dapr component spec
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The recipe used to automatically deploy underlying infrastructure for the Dapr StateStore portable resource
	Recipe *Recipe `json:"recipe,omitempty"`

	// Specifies how the underlying service/resource is provisioned and managed.
	ResourceProvisioning *ResourceProvisioning `json:"resourceProvisioning,omitempty"`

	// A collection of references to resources associated with the Dapr StateStore
	Resources []*ResourceReference `json:"resources,omitempty"`

	// Dapr StateStore type. These strings match the format used by Dapr Kubernetes configuration format.
	Type *string `json:"type,omitempty"`

	// Dapr component version
	Version *string `json:"version,omitempty"`

	// READ-ONLY; The name of the Dapr component object. Use this value in your code when interacting with the Dapr client to
// use the Dapr component.
	ComponentName *string `json:"componentName,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the DaprStateStore portable resource at the time the operation was called
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Status of a resource.
	Status *ResourceStatus `json:"status,omitempty" azure:"ro"`
}

// DaprStateStoreResource - Dapr StateStore portable resource
type DaprStateStoreResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties *DaprStateStoreProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DaprStateStoreResourceListResult - The response of a DaprStateStoreResource list operation.
type DaprStateStoreResourceListResult struct {
	// REQUIRED; The DaprStateStoreResource items on this page
	Value []*DaprStateStoreResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Recipe - The recipe used to automatically deploy underlying infrastructure for a portable resource
type Recipe struct {
	// REQUIRED; The name of the recipe within the environment to use
	Name *string `json:"name,omitempty"`

	// Key/value parameters to pass into the recipe at deployment
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceReference - Describes a reference to an existing resource
type ResourceReference struct {
	// REQUIRED; Resource id of an existing resource
	ID *string `json:"id,omitempty"`
}

// ResourceStatus - Status of a resource.
type ResourceStatus struct {
	// Properties of an output resource
	OutputResources []map[string]interface{} `json:"outputResources,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

