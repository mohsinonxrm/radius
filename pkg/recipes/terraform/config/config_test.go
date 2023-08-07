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

package config

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/project-radius/radius/pkg/corerp/datamodel"
	"github.com/project-radius/radius/pkg/recipes"
	"github.com/project-radius/radius/pkg/recipes/terraform/config/providers"
	"github.com/project-radius/radius/test/testcontext"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	testTemplatePath    = "Azure/redis/azurerm"
	testRecipeName      = "redis-azure"
	testTemplateVersion = "1.1.0"
)

var (
	envParams = map[string]any{
		"resource_group_name": "test-rg",
		"sku":                 "C",
	}

	resourceParams = map[string]any{
		"redis_cache_name": "redis-test",
		"sku":              "P",
	}
)

func setup(t *testing.T) (providers.MockProvider, map[string]providers.Provider) {
	ctrl := gomock.NewController(t)
	mProvider := providers.NewMockProvider(ctrl)
	providers := map[string]providers.Provider{
		providers.AWSProviderName:        mProvider,
		providers.AzureProviderName:      mProvider,
		providers.KubernetesProviderName: mProvider,
	}

	return *mProvider, providers
}

func getTestInputs() (recipes.EnvironmentDefinition, recipes.ResourceMetadata) {
	envRecipe := recipes.EnvironmentDefinition{
		Name:            testRecipeName,
		TemplatePath:    testTemplatePath,
		TemplateVersion: testTemplateVersion,
		Parameters:      envParams,
	}

	resourceRecipe := recipes.ResourceMetadata{
		Name:          testRecipeName,
		Parameters:    resourceParams,
		EnvironmentID: "/planes/radius/local/resourceGroups/test-group/providers/Applications.Environments/testEnv/env",
		ApplicationID: "/planes/radius/local/resourceGroups/test-group/providers/Applications.Applications/testApp/app",
		ResourceID:    "/planes/radius/local/resourceGroups/test-group/providers/Applications.Datastores/redisCaches/redis",
	}
	return envRecipe, resourceRecipe
}

func validateConfigIsGenerated(configFilePath string) (TerraformConfig, error) {
	// Read the JSON data from the main.tf.json file.
	jsonData, err := os.ReadFile(configFilePath)
	if err != nil {
		return TerraformConfig{}, err
	}

	// Unmarshal the JSON data into a TerraformConfig struct.
	var tfConfig TerraformConfig
	err = json.Unmarshal(jsonData, &tfConfig)
	if err != nil {
		return TerraformConfig{}, err
	}

	return tfConfig, nil
}

func TestGenerateTFConfigFile(t *testing.T) {
	// Create a temporary test directory.
	testDir := t.TempDir()
	envRecipe, resourceRecipe := getTestInputs()
	secretSuffix, err := generateSecretSuffix(&resourceRecipe)
	require.NoError(t, err)
	expectedTFConfig := TerraformConfig{
		Module: map[string]any{
			testRecipeName: map[string]any{
				moduleSourceKey:       testTemplatePath,
				moduleVersionKey:      testTemplateVersion,
				"resource_group_name": envParams["resource_group_name"],
				"redis_cache_name":    resourceParams["redis_cache_name"],
				"sku":                 resourceParams["sku"],
			},
		},
		Terraform: TerraformDefinition{
			Backend: map[string]interface{}{
				"kubernetes": map[string]interface{}{
					"config_path":   clientcmd.RecommendedHomeFile,
					"secret_suffix": secretSuffix,
					"namespace":     namespace,
				},
			},
		},
	}
	configFilePath, _, err := GenerateTFConfigFile(testcontext.New(t), &envRecipe, &resourceRecipe, testDir, testRecipeName)
	require.NoError(t, err)

	// Assert config file exists and contains data in expected format.
	tfConfig, err := validateConfigIsGenerated(configFilePath)
	require.NoError(t, err)

	// Assert that generated config contains the expected data.
	require.Equal(t, expectedTFConfig, tfConfig)

	// Assert generated config matches expected in JSON format.
	expectedJSON, err := os.ReadFile("testdata/tfconfig.json")
	require.NoError(t, err)
	expectedConfigDatamodel := &TerraformConfig{}
	err = json.Unmarshal(expectedJSON, expectedConfigDatamodel)
	require.NoError(t, err)
	if val, ok := expectedConfigDatamodel.Terraform.Backend["kubernetes"]; ok {
		backend := val.(map[string]interface{})
		backend["config_path"] = clientcmd.RecommendedHomeFile
	}
	expectedJSON, err = json.Marshal(expectedConfigDatamodel)
	require.NoError(t, err)
	generatedJSON, err := os.ReadFile(configFilePath)
	require.NoError(t, err)
	require.JSONEq(t, string(expectedJSON), string(generatedJSON))
}

func TestGenerateTFConfig_EmptyParameters(t *testing.T) {
	// Create a temporary test directory.
	testDir := t.TempDir()

	envRecipe, resourceRecipe := getTestInputs()
	envRecipe.Parameters = nil
	resourceRecipe.Parameters = nil
	secretSuffix, err := generateSecretSuffix(&resourceRecipe)
	require.NoError(t, err)
	expectedTFConfig := TerraformConfig{
		Module: map[string]any{
			testRecipeName: map[string]any{
				moduleSourceKey:  testTemplatePath,
				moduleVersionKey: testTemplateVersion,
			},
		},
		Terraform: TerraformDefinition{
			Backend: map[string]interface{}{
				"kubernetes": map[string]interface{}{
					"config_path":   clientcmd.RecommendedHomeFile,
					"secret_suffix": secretSuffix,
					"namespace":     namespace,
				},
			},
		},
	}

	configFilePath, _, err := GenerateTFConfigFile(testcontext.New(t), &envRecipe, &resourceRecipe, testDir, testRecipeName)
	require.NoError(t, err)

	// Assert config file exists and contains data in expected format.
	tfConfig, err := validateConfigIsGenerated(configFilePath)
	require.NoError(t, err)

	// Assert that generated config contains the expected data.
	require.Equal(t, expectedTFConfig, tfConfig)
}

func TestGenerateTFConfig_InvalidWorkingDir_Error(t *testing.T) {
	envRecipe, resourceRecipe := getTestInputs()

	// Call GenerateMainConfig with a working directory that doesn't exist.
	invalidPath := filepath.Join("invalid", uuid.New().String())
	_, _, err := GenerateTFConfigFile(testcontext.New(t), &envRecipe, &resourceRecipe, invalidPath, testRecipeName)
	require.Error(t, err)
	require.Contains(t, err.Error(), "error creating file")
}

func TestGenerateModuleData(t *testing.T) {
	t.Run("With templateVersion", func(t *testing.T) {
		expectedModuleData := map[string]any{
			moduleSourceKey:       testTemplatePath,
			moduleVersionKey:      testTemplateVersion,
			"resource_group_name": envParams["resource_group_name"],
			"redis_cache_name":    resourceParams["redis_cache_name"],
			"sku":                 resourceParams["sku"],
		}

		moduleData := generateModuleData(testcontext.New(t), testTemplatePath, testTemplateVersion, envParams, resourceParams)

		// Assert that the module data contains the expected data.
		require.Equal(t, expectedModuleData, moduleData)
	})
	t.Run("Without templateVersion", func(t *testing.T) {
		expectedModuleData := map[string]any{
			moduleSourceKey:       testTemplatePath,
			"resource_group_name": envParams["resource_group_name"],
			"redis_cache_name":    resourceParams["redis_cache_name"],
			"sku":                 resourceParams["sku"],
		}

		moduleData := generateModuleData(testcontext.New(t), testTemplatePath, "", envParams, resourceParams)

		// Assert that the module data contains the expected data.
		require.Equal(t, expectedModuleData, moduleData)
	})

}

func TestAddProviders_Success(t *testing.T) {
	ctx := testcontext.New(t)
	// Create a temporary test directory.
	testDir := t.TempDir()
	mProvider, supportedProviders := setup(t)
	envRecipe, resourceRecipe := getTestInputs()
	envConfig := recipes.Configuration{
		Providers: datamodel.Providers{
			AWS: datamodel.ProvidersAWS{
				Scope: "invalid",
			},
		},
	}
	secretSuffix, err := generateSecretSuffix(&resourceRecipe)
	require.NoError(t, err)
	awsProviderConfig := map[string]any{
		"region": "test-region",
	}
	azureProviderConfig := map[string]any{
		"subscription_id": "test-sub",
		"features":        map[string]any{},
	}
	kubernetesProviderConfig := map[string]any{
		"config_path": clientcmd.RecommendedHomeFile,
	}
	expectedTFConfig := TerraformConfig{
		Provider: map[string]any{
			providers.AWSProviderName:        awsProviderConfig,
			providers.AzureProviderName:      azureProviderConfig,
			providers.KubernetesProviderName: kubernetesProviderConfig,
		},
		Module: map[string]any{
			testRecipeName: map[string]any{
				moduleSourceKey:       testTemplatePath,
				moduleVersionKey:      testTemplateVersion,
				"resource_group_name": envParams["resource_group_name"],
				"redis_cache_name":    resourceParams["redis_cache_name"],
				"sku":                 resourceParams["sku"],
			},
		},
		Terraform: TerraformDefinition{
			Backend: map[string]interface{}{
				"kubernetes": map[string]interface{}{
					"config_path":   clientcmd.RecommendedHomeFile,
					"secret_suffix": secretSuffix,
					"namespace":     namespace,
				},
			},
		},
	}

	configFilePath, _, err := GenerateTFConfigFile(ctx, &envRecipe, &resourceRecipe, testDir, testRecipeName)
	require.NoError(t, err)

	mProvider.EXPECT().BuildConfig(ctx, &envConfig).Times(1).Return(awsProviderConfig, nil)
	mProvider.EXPECT().BuildConfig(ctx, &envConfig).Times(1).Return(azureProviderConfig, nil)
	mProvider.EXPECT().BuildConfig(ctx, &envConfig).Times(1).Return(kubernetesProviderConfig, nil)

	err = AddProviders(ctx, configFilePath, []string{providers.AWSProviderName, providers.AzureProviderName, providers.KubernetesProviderName, "sql"}, supportedProviders, &envConfig)
	require.NoError(t, err)

	// Validate that the config file exists and read the updated data.
	tfConfig, err := validateConfigIsGenerated(configFilePath)
	require.NoError(t, err)
	// Assert that the TerraformConfig contains the expected data.
	require.Equal(t, expectedTFConfig, tfConfig)
}

func TestAddProviders_InvalidScope_Error(t *testing.T) {
	ctx := testcontext.New(t)
	// Create a temporary test directory.
	testDir := t.TempDir()
	mProvider, supportedProviders := setup(t)
	envRecipe, resourceRecipe := getTestInputs()
	secretSuffix, err := generateSecretSuffix(&resourceRecipe)
	require.NoError(t, err)
	envConfig := recipes.Configuration{
		Providers: datamodel.Providers{
			AWS: datamodel.ProvidersAWS{
				Scope: "invalid",
			},
		},
	}
	expectedTFConfig := TerraformConfig{
		Module: map[string]any{
			testRecipeName: map[string]any{
				moduleSourceKey:       testTemplatePath,
				moduleVersionKey:      testTemplateVersion,
				"resource_group_name": envParams["resource_group_name"],
				"redis_cache_name":    resourceParams["redis_cache_name"],
				"sku":                 resourceParams["sku"],
			},
		},
		Terraform: TerraformDefinition{
			Backend: map[string]interface{}{
				"kubernetes": map[string]interface{}{
					"config_path":   clientcmd.RecommendedHomeFile,
					"secret_suffix": secretSuffix,
					"namespace":     namespace,
				},
			},
		},
	}

	configFilePath, _, err := GenerateTFConfigFile(ctx, &envRecipe, &resourceRecipe, testDir, testRecipeName)
	require.NoError(t, err)

	mProvider.EXPECT().BuildConfig(ctx, &envConfig).Times(1).Return(nil, errors.New("Invalid AWS provider scope"))

	err = AddProviders(ctx, configFilePath, []string{providers.AWSProviderName}, supportedProviders, &envConfig)
	require.Error(t, err)

	// Validate that the config file still exists and was not updated.
	tfConfig, err := validateConfigIsGenerated(configFilePath)
	require.NoError(t, err)
	require.Equal(t, expectedTFConfig, tfConfig)
}

func TestAddProviders_EmptyProviderConfigurations_Success(t *testing.T) {
	ctx := testcontext.New(t)
	// Create a temporary test directory.
	testDir := t.TempDir()

	mProvider, supportedProviders := setup(t)
	envRecipe, resourceRecipe := getTestInputs()
	envConfig := recipes.Configuration{}
	secretSuffix, err := generateSecretSuffix(&resourceRecipe)
	require.NoError(t, err)
	// Expected config shouldn't contain any provider config
	expectedTFConfig := TerraformConfig{
		Module: map[string]any{
			testRecipeName: map[string]any{
				moduleSourceKey:       testTemplatePath,
				moduleVersionKey:      testTemplateVersion,
				"resource_group_name": envParams["resource_group_name"],
				"redis_cache_name":    resourceParams["redis_cache_name"],
				"sku":                 resourceParams["sku"],
			},
		},
		Terraform: TerraformDefinition{
			Backend: map[string]interface{}{
				"kubernetes": map[string]interface{}{
					"config_path":   clientcmd.RecommendedHomeFile,
					"secret_suffix": secretSuffix,
					"namespace":     namespace,
				},
			},
		},
	}

	configFilePath, _, err := GenerateTFConfigFile(ctx, &envRecipe, &resourceRecipe, testDir, testRecipeName)
	require.NoError(t, err)

	// Expect build config function call for AWS provider with empty output since envConfig has empty AWS scope
	mProvider.EXPECT().BuildConfig(ctx, &envConfig).Times(1).Return(map[string]any{}, nil)

	err = AddProviders(ctx, configFilePath, []string{providers.AWSProviderName}, supportedProviders, &envConfig)
	require.NoError(t, err)

	// Validate that the config file exists and read the updated data.
	tfConfig, err := validateConfigIsGenerated(configFilePath)
	require.NoError(t, err)
	// Assert that the TerraformConfig contains the expected data.
	require.Equal(t, expectedTFConfig, tfConfig)
}

// Empty AWS scope should return empty AWS provider config
func TestAddProviders_EmptyAWSScope(t *testing.T) {
	ctx := testcontext.New(t)
	// Create a temporary test directory.
	testDir := t.TempDir()
	mProvider, supportedProviders := setup(t)
	envRecipe, resourceRecipe := getTestInputs()
	secretSuffix, err := generateSecretSuffix(&resourceRecipe)
	require.NoError(t, err)
	envConfig := recipes.Configuration{
		Providers: datamodel.Providers{
			AWS: datamodel.ProvidersAWS{
				Scope: "",
			},
			Azure: datamodel.ProvidersAzure{
				Scope: "/subscriptions/test-sub/resourceGroups/test-rg",
			},
		},
	}

	// Expected config shouldn't contain any provider config
	expectedTFConfig := TerraformConfig{
		Module: map[string]any{
			testRecipeName: map[string]any{
				moduleSourceKey:       testTemplatePath,
				moduleVersionKey:      testTemplateVersion,
				"resource_group_name": envParams["resource_group_name"],
				"redis_cache_name":    resourceParams["redis_cache_name"],
				"sku":                 resourceParams["sku"],
			},
		},
		Terraform: TerraformDefinition{
			Backend: map[string]interface{}{
				"kubernetes": map[string]interface{}{
					"config_path":   clientcmd.RecommendedHomeFile,
					"secret_suffix": secretSuffix,
					"namespace":     namespace,
				},
			},
		},
	}

	configFilePath, _, err := GenerateTFConfigFile(ctx, &envRecipe, &resourceRecipe, testDir, testRecipeName)
	require.NoError(t, err)

	mProvider.EXPECT().BuildConfig(ctx, &envConfig).Times(1).Return(nil, nil)

	err = AddProviders(ctx, configFilePath, []string{providers.AWSProviderName}, supportedProviders, &envConfig)
	require.NoError(t, err)

	// Validate that the config file exists and read the updated data.
	tfConfig, err := validateConfigIsGenerated(configFilePath)
	require.NoError(t, err)

	// Assert that the TerraformConfig contains the expected data.
	require.Equal(t, expectedTFConfig, tfConfig)
}

func TestAddProviders_MissingAzureProvider(t *testing.T) {
	ctx := testcontext.New(t)
	// Create a temporary test directory.
	testDir := t.TempDir()
	mProvider, supportedProviders := setup(t)
	envRecipe, resourceRecipe := getTestInputs()
	envConfig := recipes.Configuration{}
	secretSuffix, err := generateSecretSuffix(&resourceRecipe)
	require.NoError(t, err)
	azureProviderConfig := map[string]any{
		"features": map[string]any{},
	}
	// Expected config shouldn't contain Azure subscription id in the provider config
	expectedTFConfig := TerraformConfig{
		Provider: map[string]any{
			providers.AzureProviderName: azureProviderConfig,
		},
		Module: map[string]any{
			testRecipeName: map[string]any{
				moduleSourceKey:       testTemplatePath,
				moduleVersionKey:      testTemplateVersion,
				"resource_group_name": envParams["resource_group_name"],
				"redis_cache_name":    resourceParams["redis_cache_name"],
				"sku":                 resourceParams["sku"],
			},
		},
		Terraform: TerraformDefinition{
			Backend: map[string]interface{}{
				"kubernetes": map[string]interface{}{
					"config_path":   clientcmd.RecommendedHomeFile,
					"secret_suffix": secretSuffix,
					"namespace":     namespace,
				},
			},
		},
	}

	configFilePath, _, err := GenerateTFConfigFile(ctx, &envRecipe, &resourceRecipe, testDir, testRecipeName)
	require.NoError(t, err)

	mProvider.EXPECT().BuildConfig(ctx, &envConfig).Times(1).Return(azureProviderConfig, nil)

	err = AddProviders(ctx, configFilePath, []string{providers.AzureProviderName}, supportedProviders, &envConfig)
	require.NoError(t, err)

	// Validate that the config file exists and read the updated data.
	tfConfig, err := validateConfigIsGenerated(configFilePath)
	require.NoError(t, err)

	// Assert that the TerraformConfig contains the expected data.
	require.Equal(t, expectedTFConfig, tfConfig)
}

func TestAddProviders_OpenConfigFileError(t *testing.T) {
	ctx := testcontext.New(t)
	mProvider, supportedProviders := setup(t)
	kubernetesProviderConfig := map[string]any{
		"config_path": clientcmd.RecommendedHomeFile,
	}

	mProvider.EXPECT().BuildConfig(ctx, nil).Times(1).Return(kubernetesProviderConfig, nil)

	// Call AddProviders with a non-existent file path.
	err := AddProviders(ctx, "/path/to/non-existent/file.json", []string{providers.KubernetesProviderName}, supportedProviders, nil)

	// Assert that AddProviders returns an error.
	require.Error(t, err)
	require.Contains(t, err.Error(), "no such file or directory")
}

func TestAddProviders_DecodeError(t *testing.T) {
	ctx := testcontext.New(t)
	mProvider, supportedProviders := setup(t)
	// Create a temporary test directory.
	testDir := t.TempDir()
	// Create a test configuration file with invalid JSON data.
	configFile := filepath.Join(testDir, "test.json")
	err := os.WriteFile(configFile, []byte(`invalid json data`), 0644)
	require.NoError(t, err)

	kubernetesProviderConfig := map[string]any{
		"config_path": clientcmd.RecommendedHomeFile,
	}
	mProvider.EXPECT().BuildConfig(ctx, nil).Times(1).Return(kubernetesProviderConfig, nil)

	// Call AddProviders with the test configuration file and required providers.
	err = AddProviders(ctx, configFile, []string{providers.KubernetesProviderName}, supportedProviders, nil)

	// Assert that AddProviders returns an error.
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid character")
}

func TestAddProviders_WriteConfigFileError(t *testing.T) {
	ctx := testcontext.New(t)
	mProvider, supportedProviders := setup(t)
	// Create a temporary test directory.
	testDir := t.TempDir()

	// Create a test configuration file.
	configFile := filepath.Join(testDir, "test.json")
	err := os.WriteFile(configFile, []byte(`{"module":{}}`), 0644)
	require.NoError(t, err)
	// Mock a write file error by setting the file permissions to read-only.
	err = os.Chmod(configFile, 0400)
	require.NoError(t, err)

	kubernetesProviderConfig := map[string]any{
		"config_path": clientcmd.RecommendedHomeFile,
	}
	mProvider.EXPECT().BuildConfig(ctx, nil).Times(1).Return(kubernetesProviderConfig, nil)

	// Call AddProviders with the test configuration file and required providers.
	err = AddProviders(ctx, configFile, []string{providers.KubernetesProviderName}, supportedProviders, nil)

	// Assert that AddProviders returns an error.
	require.Error(t, err)
	require.Contains(t, err.Error(), "permission denied")
}

func TestGenerateSecretSuffix_invalid_resourceid(t *testing.T) {
	_, resourceRecipe := getTestInputs()
	resourceRecipe.ResourceID = "invalid"
	_, err := generateSecretSuffix(&resourceRecipe)
	require.Equal(t, err.Error(), "'invalid' is not a valid resource id")
}

func TestGenerateSecretSuffix_with_lengthy_resource_name(t *testing.T) {
	_, resourceRecipe := getTestInputs()
	act, err := generateSecretSuffix(&resourceRecipe)
	require.NoError(t, err)
	hasher := sha1.New()
	_, _ = hasher.Write([]byte(strings.ToLower("env-app-" + resourceRecipe.ResourceID)))
	hash := hasher.Sum(nil)
	require.Equal(t, act, "env-app-redis."+fmt.Sprintf("%x", hash))
}
