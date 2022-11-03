// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package v20220315privatepreview

import (
	azto "github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/project-radius/radius/pkg/armrpc/api/conv"
	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/corerp/datamodel"
	"github.com/project-radius/radius/pkg/rp"
)

// ConvertTo converts from the versioned Container resource to version-agnostic datamodel.
func (src *ContainerResource) ConvertTo() (conv.DataModelInterface, error) {
	// Note: SystemData conversion isn't required since this property comes ARM and datastore.

	connections := make(map[string]datamodel.ConnectionProperties)
	for key, val := range src.Properties.Connections {
		if val != nil {
			roles := []string{}
			var kind datamodel.IAMKind

			if val.Iam != nil {
				for _, r := range val.Iam.Roles {
					roles = append(roles, to.String(r))
				}
				kind = toKindDataModel(val.Iam.Kind)
			}

			var disableDefaultEnvVars bool
			if val.DisableDefaultEnvVars != nil {
				disableDefaultEnvVars = to.Bool(val.DisableDefaultEnvVars)
			}

			connections[key] = datamodel.ConnectionProperties{
				Source:                to.String(val.Source),
				DisableDefaultEnvVars: &disableDefaultEnvVars,
				IAM: datamodel.IAMProperties{
					Kind:  kind,
					Roles: roles,
				},
			}
		}
	}

	var livenessProbe datamodel.HealthProbeProperties
	if src.Properties.Container.LivenessProbe != nil {
		livenessProbe = toHealthProbePropertiesDataModel(src.Properties.Container.LivenessProbe)
	}

	var readinessProbe datamodel.HealthProbeProperties
	if src.Properties.Container.ReadinessProbe != nil {
		readinessProbe = toHealthProbePropertiesDataModel(src.Properties.Container.ReadinessProbe)
	}

	ports := make(map[string]datamodel.ContainerPort)
	for key, val := range src.Properties.Container.Ports {
		ports[key] = datamodel.ContainerPort{
			ContainerPort: to.Int32(val.ContainerPort),
			Protocol:      toProtocolDataModel(val.Protocol),
			Provides:      to.String(val.Provides),
		}
	}

	var volumes map[string]datamodel.VolumeProperties
	if src.Properties.Container.Volumes != nil {
		volumes = make(map[string]datamodel.VolumeProperties)
		for key, val := range src.Properties.Container.Volumes {
			volumes[key] = toVolumePropertiesDataModel(val)
		}
	}

	var extensions []datamodel.Extension
	if src.Properties.Extensions != nil {
		for _, e := range src.Properties.Extensions {
			extensions = append(extensions, toExtensionDataModel(e))
		}
	}

	converted := &datamodel.ContainerResource{
		BaseResource: v1.BaseResource{
			TrackedResource: v1.TrackedResource{
				ID:       to.String(src.ID),
				Name:     to.String(src.Name),
				Type:     to.String(src.Type),
				Location: to.String(src.Location),
				Tags:     to.StringMap(src.Tags),
			},
			InternalMetadata: v1.InternalMetadata{
				UpdatedAPIVersion:      Version,
				AsyncProvisioningState: toProvisioningStateDataModel(src.Properties.ProvisioningState),
			},
		},
		Properties: datamodel.ContainerProperties{
			BasicResourceProperties: rp.BasicResourceProperties{
				Application: to.String(src.Properties.Application),
			},
			Connections: connections,
			Container: datamodel.Container{
				Image:          to.String(src.Properties.Container.Image),
				Env:            to.StringMap(src.Properties.Container.Env),
				LivenessProbe:  livenessProbe,
				Ports:          ports,
				ReadinessProbe: readinessProbe,
				Volumes:        volumes,
				Command:        stringSlice(src.Properties.Container.Command),
				Args:           stringSlice(src.Properties.Container.Args),
				WorkingDir:     to.String(src.Properties.Container.WorkingDir),
			},
			Extensions: extensions,
		},
	}

	if src.Properties.Identity != nil {
		converted.Properties.Identity = &rp.IdentitySettings{
			Kind:       toIdentityKind(src.Properties.Identity.Kind),
			OIDCIssuer: to.String(src.Properties.Identity.OidcIssuer),
			Resource:   to.String(src.Properties.Identity.Resource),
		}
	}

	return converted, nil
}

// ConvertFrom converts from version-agnostic datamodel to the versioned Container resource.
func (dst *ContainerResource) ConvertFrom(src conv.DataModelInterface) error {
	c, ok := src.(*datamodel.ContainerResource)
	if !ok {
		return conv.ErrInvalidModelConversion
	}

	connections := make(map[string]*ConnectionProperties)
	for key, val := range c.Properties.Connections {
		roles := []*string{}
		var kind *Kind

		for _, r := range val.IAM.Roles {
			roles = append(roles, to.StringPtr(r))
		}

		kind = fromKindDataModel(val.IAM.Kind)

		var disableDefaultEnvVars bool
		if val.DisableDefaultEnvVars != nil {
			disableDefaultEnvVars = to.Bool(val.DisableDefaultEnvVars)
		}

		connections[key] = &ConnectionProperties{
			Source:                to.StringPtr(val.Source),
			DisableDefaultEnvVars: &disableDefaultEnvVars,
			Iam: &IamProperties{
				Kind:  kind,
				Roles: roles,
			},
		}
	}

	var livenessProbe HealthProbePropertiesClassification
	if !c.Properties.Container.LivenessProbe.IsEmpty() {
		livenessProbe = fromHealthProbePropertiesDataModel(c.Properties.Container.LivenessProbe)
	}

	var readinessProbe HealthProbePropertiesClassification
	if !c.Properties.Container.ReadinessProbe.IsEmpty() {
		readinessProbe = fromHealthProbePropertiesDataModel(c.Properties.Container.ReadinessProbe)
	}

	ports := make(map[string]*ContainerPort)
	for key, val := range c.Properties.Container.Ports {
		ports[key] = &ContainerPort{
			ContainerPort: to.Int32Ptr(val.ContainerPort),
			Protocol:      fromProtocolDataModel(val.Protocol),
			Provides:      to.StringPtr(val.Provides),
		}
	}

	var volumes map[string]VolumeClassification
	if c.Properties.Container.Volumes != nil {
		volumes = make(map[string]VolumeClassification)
		for key, val := range c.Properties.Container.Volumes {
			volumes[key] = fromVolumePropertiesDataModel(val)
		}
	}

	var extensions []ContainerExtensionClassification
	if c.Properties.Extensions != nil {
		for _, e := range c.Properties.Extensions {
			extensions = append(extensions, fromExtensionClassificationDataModel(e))
		}
	}

	var identity *IdentitySettings
	if c.Properties.Identity != nil {
		identity = &IdentitySettings{
			Kind:       fromIdentityKind(c.Properties.Identity.Kind),
			Resource:   azto.Ptr(c.Properties.Identity.Resource),
			OidcIssuer: azto.Ptr(c.Properties.Identity.OIDCIssuer),
		}
	}

	dst.ID = to.StringPtr(c.ID)
	dst.Name = to.StringPtr(c.Name)
	dst.Type = to.StringPtr(c.Type)
	dst.SystemData = fromSystemDataModel(c.SystemData)
	dst.Location = to.StringPtr(c.Location)
	dst.Tags = *to.StringMapPtr(c.Tags)
	dst.Properties = &ContainerProperties{
		Status: &ResourceStatus{
			OutputResources: rp.BuildExternalOutputResources(c.Properties.Status.OutputResources),
		},
		ProvisioningState: fromProvisioningStateDataModel(c.InternalMetadata.AsyncProvisioningState),
		Application:       to.StringPtr(c.Properties.Application),
		Connections:       connections,
		Container: &Container{
			Image:          to.StringPtr(c.Properties.Container.Image),
			Env:            *to.StringMapPtr(c.Properties.Container.Env),
			LivenessProbe:  livenessProbe,
			Ports:          ports,
			ReadinessProbe: readinessProbe,
			Volumes:        volumes,
			Command:        azto.SliceOfPtrs(c.Properties.Container.Command...),
			Args:           azto.SliceOfPtrs(c.Properties.Container.Args...),
			WorkingDir:     to.StringPtr(c.Properties.Container.WorkingDir),
		},
		Extensions: extensions,
		Identity:   identity,
	}

	return nil
}

func toHealthProbePropertiesDataModel(h HealthProbePropertiesClassification) datamodel.HealthProbeProperties {
	switch c := h.(type) {
	case *ExecHealthProbeProperties:
		converted := &datamodel.HealthProbeProperties{
			Kind: datamodel.ExecHealthProbe,
			Exec: &datamodel.ExecHealthProbeProperties{
				HealthProbeBase: toHealthProbeBase(*c.GetHealthProbeProperties()),
				Command:         to.String(c.Command),
			},
		}
		return *converted
	case *HTTPGetHealthProbeProperties:
		converted := &datamodel.HealthProbeProperties{
			Kind: datamodel.HTTPGetHealthProbe,
			HTTPGet: &datamodel.HTTPGetHealthProbeProperties{
				HealthProbeBase: toHealthProbeBase(*c.GetHealthProbeProperties()),
				ContainerPort:   to.Int32(c.ContainerPort),
				Path:            to.String(c.Path),
				Headers:         to.StringMap(c.Headers),
			},
		}
		return *converted
	case *TCPHealthProbeProperties:
		converted := &datamodel.HealthProbeProperties{
			Kind: datamodel.TCPHealthProbe,
			TCP: &datamodel.TCPHealthProbeProperties{
				HealthProbeBase: toHealthProbeBase(*c.GetHealthProbeProperties()),
				ContainerPort:   to.Int32(c.ContainerPort),
			},
		}
		return *converted
	}

	return datamodel.HealthProbeProperties{}
}

func fromHealthProbePropertiesDataModel(h datamodel.HealthProbeProperties) HealthProbePropertiesClassification {
	switch h.Kind {
	case datamodel.ExecHealthProbe:
		converted := ExecHealthProbeProperties{
			Kind:                (*string)(&h.Kind),
			FailureThreshold:    h.Exec.FailureThreshold,
			InitialDelaySeconds: h.Exec.InitialDelaySeconds,
			PeriodSeconds:       h.Exec.PeriodSeconds,
			TimeoutSeconds:      h.Exec.TimeoutSeconds,
			Command:             to.StringPtr(h.Exec.Command),
		}
		return &converted
	case datamodel.HTTPGetHealthProbe:
		converted := HTTPGetHealthProbeProperties{
			Kind:                (*string)(&h.Kind),
			FailureThreshold:    h.HTTPGet.FailureThreshold,
			InitialDelaySeconds: h.HTTPGet.InitialDelaySeconds,
			PeriodSeconds:       h.HTTPGet.PeriodSeconds,
			TimeoutSeconds:      h.HTTPGet.TimeoutSeconds,
			ContainerPort:       to.Int32Ptr(h.HTTPGet.ContainerPort),
			Path:                to.StringPtr(h.HTTPGet.Path),
			Headers:             *to.StringMapPtr(h.HTTPGet.Headers),
		}
		return &converted
	case datamodel.TCPHealthProbe:
		converted := TCPHealthProbeProperties{
			Kind:                (*string)(&h.Kind),
			FailureThreshold:    h.TCP.FailureThreshold,
			InitialDelaySeconds: h.TCP.InitialDelaySeconds,
			PeriodSeconds:       h.TCP.PeriodSeconds,
			TimeoutSeconds:      h.TCP.TimeoutSeconds,
			ContainerPort:       to.Int32Ptr(h.TCP.ContainerPort),
		}
		return &converted
	}

	return nil
}

func toKindDataModel(kind *Kind) datamodel.IAMKind {
	// TODO: This always returns datamodel.KindAzure. Why?
	switch *kind {
	case KindAzure:
		return datamodel.KindAzure
	default:
		return datamodel.KindAzure
	}
}

func fromKindDataModel(kind datamodel.IAMKind) *Kind {
	var k Kind
	switch kind {
	case datamodel.KindAzure:
		k = KindAzure
	default:
		k = KindAzure
	}
	return &k
}

func toProtocolDataModel(protocol *Protocol) datamodel.Protocol {
	if protocol == nil {
		return datamodel.ProtocolHTTP
	}
	switch *protocol {
	case ProtocolHTTP:
		return datamodel.ProtocolHTTP
	case ProtocolGrpc:
		return datamodel.ProtocolGrpc
	case ProtocolTCP:
		return datamodel.ProtocolTCP
	case ProtocolUDP:
		return datamodel.ProtocolUDP
	default:
		return datamodel.ProtocolHTTP
	}
}

func fromProtocolDataModel(protocol datamodel.Protocol) *Protocol {
	var p Protocol
	switch protocol {
	case datamodel.ProtocolHTTP:
		p = ProtocolHTTP
	case datamodel.ProtocolGrpc:
		p = ProtocolGrpc
	case datamodel.ProtocolTCP:
		p = ProtocolTCP
	case datamodel.ProtocolUDP:
		p = ProtocolUDP
	default:
		p = ProtocolHTTP
	}
	return &p
}

func toVolumePropertiesDataModel(h VolumeClassification) datamodel.VolumeProperties {
	switch c := h.(type) {
	case *EphemeralVolume:
		converted := &datamodel.VolumeProperties{
			Kind: datamodel.Ephemeral,
			Ephemeral: &datamodel.EphemeralVolume{
				VolumeBase:   toVolumeBaseDataModel(*c.GetVolume()),
				ManagedStore: toManagedStoreDataModel(c.ManagedStore),
			},
		}
		return *converted
	case *PersistentVolume:
		converted := &datamodel.VolumeProperties{
			Kind: datamodel.Persistent,
			Persistent: &datamodel.PersistentVolume{
				VolumeBase: toVolumeBaseDataModel(*c.GetVolume()),
				Source:     to.String(c.Source),
				Permission: toPermissionDataModel(c.Permission),
			},
		}
		return *converted
	}

	return datamodel.VolumeProperties{}
}

func fromVolumePropertiesDataModel(v datamodel.VolumeProperties) VolumeClassification {
	switch v.Kind {
	case datamodel.Ephemeral:
		converted := EphemeralVolume{
			Kind:         (*string)(&v.Kind),
			MountPath:    &v.Ephemeral.MountPath,
			ManagedStore: fromManagedStoreDataModel(v.Ephemeral.ManagedStore),
		}
		return converted.GetVolume()
	case datamodel.Persistent:
		converted := PersistentVolume{
			Kind:       (*string)(&v.Kind),
			MountPath:  &v.Persistent.MountPath,
			Source:     &v.Persistent.Source,
			Permission: fromPermissionDataModel(v.Persistent.Permission),
		}
		return converted.GetVolume()
	}

	return nil
}

func toManagedStoreDataModel(ms *ManagedStore) datamodel.ManagedStore {
	switch *ms {
	case ManagedStoreDisk:
		return datamodel.ManagedStoreDisk
	case ManagedStoreMemory:
		return datamodel.ManagedStoreMemory
	default:
		return datamodel.ManagedStoreDisk
	}
}

func fromManagedStoreDataModel(managedStore datamodel.ManagedStore) *ManagedStore {
	var m ManagedStore
	switch managedStore {
	case datamodel.ManagedStoreDisk:
		m = ManagedStoreDisk
	case datamodel.ManagedStoreMemory:
		m = ManagedStoreMemory
	default:
		m = ManagedStoreDisk
	}
	return &m
}

func toPermissionDataModel(rbac *VolumePermission) datamodel.VolumePermission {
	if rbac == nil {
		return datamodel.VolumePermissionRead
	}

	switch *rbac {
	case VolumePermissionRead:
		return datamodel.VolumePermissionRead
	case VolumePermissionWrite:
		return datamodel.VolumePermissionWrite
	default:
		return datamodel.VolumePermissionRead
	}
}

func fromPermissionDataModel(rbac datamodel.VolumePermission) *VolumePermission {
	var r VolumePermission
	switch rbac {
	case datamodel.VolumePermissionRead:
		r = VolumePermissionRead
	case datamodel.VolumePermissionWrite:
		r = VolumePermissionWrite
	default:
		r = VolumePermissionRead
	}
	return &r
}

func toExtensionDataModel(e ExtensionClassification) datamodel.Extension {
	switch c := e.(type) {
	case *ManualScalingExtension:
		converted := &datamodel.Extension{
			Kind: datamodel.ManualScaling,
			ManualScaling: &datamodel.ManualScalingExtension{
				Replicas: c.Replicas,
			},
		}
		return *converted
	case *DaprSidecarExtension:
		converted := &datamodel.Extension{
			Kind: datamodel.DaprSidecar,
			DaprSidecar: &datamodel.DaprSidecarExtension{
				AppID:    to.String(c.AppID),
				AppPort:  to.Int32(c.AppPort),
				Config:   to.String(c.Config),
				Protocol: toProtocolDataModel(c.Protocol),
				Provides: to.String(c.Provides),
			},
		}
		return *converted
	}

	return datamodel.Extension{}
}

func fromExtensionClassificationDataModel(e datamodel.Extension) ContainerExtensionClassification {
	switch e.Kind {
	case datamodel.ManualScaling:
		converted := ManualScalingExtension{
			Kind:     to.StringPtr(string(e.Kind)),
			Replicas: e.ManualScaling.Replicas,
		}
		return converted.GetContainerExtension()
	case datamodel.DaprSidecar:
		converted := DaprSidecarExtension{
			Kind:     to.StringPtr(string(e.Kind)),
			AppID:    to.StringPtr(e.DaprSidecar.AppID),
			AppPort:  to.Int32Ptr(e.DaprSidecar.AppPort),
			Config:   to.StringPtr(e.DaprSidecar.Config),
			Protocol: fromProtocolDataModel(e.DaprSidecar.Protocol),
			Provides: to.StringPtr(e.DaprSidecar.Provides),
		}
		return converted.GetContainerExtension()
	}

	return nil
}

func toHealthProbeBase(h HealthProbeProperties) datamodel.HealthProbeBase {
	return datamodel.HealthProbeBase{
		FailureThreshold:    h.FailureThreshold,
		InitialDelaySeconds: h.InitialDelaySeconds,
		PeriodSeconds:       h.PeriodSeconds,
		TimeoutSeconds:      h.TimeoutSeconds,
	}
}

func toVolumeBaseDataModel(v Volume) datamodel.VolumeBase {
	return datamodel.VolumeBase{
		MountPath: *v.MountPath,
	}
}
