/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resource_test

import (
	"context"
	"testing"

	"github.com/project-radius/radius/test/functional/shared"
	"github.com/project-radius/radius/test/step"
	"github.com/project-radius/radius/test/validation"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_SecretStore_CreateSecret(t *testing.T) {
	template := "testdata/corerp-resources-secretstore-new.bicep"
	appName := "corerp-resources-secretstore"
	appNamespace := "corerp-resources-secretstore-app"

	test := shared.NewRPTest(t, appNamespace, []shared.TestStep{
		{
			Executor: step.NewDeployExecutor(template, "@testdata/parameters/test-tls-cert.parameters.json"),
			RPResources: &validation.RPResourceSet{
				Resources: []validation.RPResource{
					{
						Name: appName,
						Type: validation.ApplicationsResource,
					},
					{
						Name: "appcert",
						Type: validation.SecretStoresResource,
						App:  appName,
					},
					{
						Name: "appsecret",
						Type: validation.SecretStoresResource,
						App:  appName,
					},
				},
			},
			K8sObjects: &validation.K8sObjectSet{
				Namespaces: map[string][]validation.K8sObject{
					appNamespace: {
						validation.NewK8sSecretForResource(appName, "appcert"),
						validation.NewK8sSecretForResource(appName, "appsecret"),
					},
				},
			},
			PostStepVerify: func(ctx context.Context, t *testing.T, test shared.RPTest) {
				secret, err := test.Options.K8sClient.CoreV1().Secrets(appNamespace).Get(ctx, "appcert", metav1.GetOptions{})
				require.NoError(t, err)

				for _, key := range []string{"tls.key", "tls.crt"} {
					_, ok := secret.Data[key]
					require.True(t, ok)
				}

				secret, err = test.Options.K8sClient.CoreV1().Secrets(appNamespace).Get(ctx, "appsecret", metav1.GetOptions{})
				require.NoError(t, err)

				for _, key := range []string{"servicePrincipalPassword", "appId", "tenantId"} {
					_, ok := secret.Data[key]
					require.True(t, ok)
				}
			},
		},
	})

	test.Test(t)
}

func Test_SecretStore_ReferenceSecret(t *testing.T) {
	template := "testdata/corerp-resources-secretstore-ref.bicep"
	appName := "corerp-resources-secretstore-ref"
	appNamespace := "corerp-resources-secretstore-ref"

	secret := shared.K8sSecretResource("default", "secret-app-existing-secret", "kubernetes.io/tls", "tls.crt", "fakecertval", "tls.key", "fakekeyval")

	test := shared.NewRPTest(t, appNamespace, []shared.TestStep{
		{
			Executor: step.NewDeployExecutor(template),
			RPResources: &validation.RPResourceSet{
				Resources: []validation.RPResource{
					{
						Name: appName,
						Type: validation.ApplicationsResource,
					},
					{
						Name: "existing-appcert",
						Type: validation.SecretStoresResource,
						App:  appName,
					},
				},
			},
			SkipObjectValidation: true,
		},
	}, secret)

	test.Test(t)
}
