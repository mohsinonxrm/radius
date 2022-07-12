// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package handlers

import (
	"context"
	"testing"

	"github.com/project-radius/radius/pkg/radrp/outputresource"
	"github.com/project-radius/radius/pkg/resourcekinds"
	"github.com/stretchr/testify/require"
	controller_runtime "sigs.k8s.io/controller-runtime/pkg/client/fake"

	k8s "github.com/project-radius/radius/pkg/kubernetes"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

var (
	applicationName = "testApplication"
	resourceName    = "testResource"
	objectMeta      = metav1.ObjectMeta{
		Name:      resourceName,
		Namespace: applicationName,
	}
	deployment = appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      k8s.MakeResourceName(applicationName, resourceName),
			Namespace: applicationName,
			Labels:    k8s.MakeDescriptiveLabels(applicationName, resourceName),
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: k8s.MakeSelectorLabels(applicationName, resourceName),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
				Spec:       corev1.PodSpec{},
			},
		},
	}
)

func Test_Deployment_Readiness_Success(t *testing.T) {
	// Arrange
	clientSet := fake.NewSimpleClientset()
	client := controller_runtime.NewFakeClient()

	handler := NewKubernetesHandler(client, clientSet)

	testOutputResource := outputresource.NewKubernetesOutputResource(resourcekinds.Deployment, outputresource.LocalIDDeployment, &deployment, deployment.ObjectMeta)

	ctx := context.Background()
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: applicationName,
		},
	}

	err := client.Create(ctx, ns)
	require.Equal(t, err, nil)
	err = client.Create(ctx, &deployment)
	require.Equal(t, err, nil)

	success_condition := appsv1.DeploymentCondition{
		Status: "True",
		Type:   appsv1.DeploymentProgressing,
		Reason: "NewReplicaSetAvailable",
	}
	deployment.Status.Conditions = append(deployment.Status.Conditions, success_condition)
	_, err = clientSet.AppsV1().Deployments(applicationName).Create(ctx, &deployment, metav1.CreateOptions{})
	require.Equal(t, err, nil)

	errorCh := make(chan error, 1)
	go func() {
		// Action - Test the handler
		errorCh <- handler.Put(ctx, &testOutputResource)
	}()

	err = <-errorCh
	require.Equal(t, err, nil)
}

func Test_Deployment_Readiness_Failed(t *testing.T) {
	// Arrange
	clientSet := fake.NewSimpleClientset()
	client := controller_runtime.NewFakeClient()

	handler := NewKubernetesHandler(client, clientSet)

	testOutputResource := outputresource.NewKubernetesOutputResource(resourcekinds.Deployment, outputresource.LocalIDDeployment, &deployment, deployment.ObjectMeta)
	TestHook = true
	ctx := context.Background()
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: applicationName,
		},
	}

	err := client.Create(ctx, ns)
	require.Equal(t, err, nil)
	err = client.Create(ctx, &deployment)
	require.Equal(t, err, nil)

	failed_condition := appsv1.DeploymentCondition{
		Status: "False",
		Type:   appsv1.DeploymentReplicaFailure,
		Reason: "FailedCreate",
	}
	deployment.Status.Conditions = append(deployment.Status.Conditions, failed_condition)
	_, err = clientSet.AppsV1().Deployments(applicationName).Create(ctx, &deployment, metav1.CreateOptions{})
	require.Equal(t, err, nil)

	errorCh := make(chan error, 1)
	go func() {
		// Action - Test the handler
		errorCh <- handler.Put(ctx, &testOutputResource)
	}()

	err = <-errorCh
	require.NotEqual(t, err, nil)
	require.Error(t, err, "deployment timed out, name: testApplication-testResource, namespace testApplication, status: , reason: FailedCreate")
}
