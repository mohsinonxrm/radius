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

package radinit

import (
	reflect "reflect"
	"testing"

	corerp "github.com/radius-project/radius/pkg/corerp/api/v20231001preview"
	"github.com/radius-project/radius/pkg/recipes"
	"github.com/radius-project/radius/pkg/to"
	"github.com/stretchr/testify/require"
)

func Test_getNormalizedResourceTypeFromPath(t *testing.T) {
	t.Run("Successfully returns metadata", func(t *testing.T) {
		resourceType := getNormalizedResourceTypeFromPath("ghcr.io/radius-project/recipes/local-dev/rediscaches")
		require.Equal(t, "rediscaches", resourceType)
	})

	tests := []struct {
		name     string
		repo     string
		expected string
	}{
		{
			"Path With No Resource Type",
			"randomRepo",
			"",
		},
		{
			"Valid Path",
			"ghcr.io/radius-project/recipes/local-dev/rediscaches",
			"rediscaches",
		},
		{
			"Invalid Path #1",
			"ghcr.io/radius-project/recipes////local-dev/rediscaches",
			"",
		},
		{
			"Invalid Path #2",
			"ghcr.io/radius-project/recipes/local-dev////rediscaches",
			"",
		},
		{
			"Path With Extra Path Argument",
			"ghcr.io/radius-project/recipes/local-dev/rediscaches/testing",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resourceType := getNormalizedResourceTypeFromPath(tt.repo)
			require.Equal(t, tt.expected, resourceType)
		})
	}
}

func Test_getActualResourceType(t *testing.T) {
	tests := []struct {
		name         string
		resourceType string
		want         string
	}{
		{
			"Dapr PubSub Portable Resource",
			"pubsubbrokers",
			"Applications.Dapr/pubSubBrokers",
		},
		{
			"Dapr Secret Store Portable Resource",
			"secretstores",
			"Applications.Dapr/secretStores",
		},
		{
			"Dapr State Store Portable Resource",
			"statestores",
			"Applications.Dapr/stateStores",
		},
		{
			"Rabbit MQ Portable Resource",
			"rabbitmqqueues",
			"Applications.Messaging/rabbitMQQueues",
		},
		{
			"Redis Cache Portable Resource",
			"rediscaches",
			"Applications.Datastores/redisCaches",
		},
		{
			"Mongo Database Portable Resource",
			"mongodatabases",
			"Applications.Datastores/mongoDatabases",
		},
		{
			"SQL Database Portable Resource",
			"sqldatabases",
			"Applications.Datastores/sqlDatabases",
		},
		{
			"Extenders",
			"extenders",
			"Applications.Core/extenders",
		},
		{
			"Invalid Portable Resource",
			"unsupported",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getActualResourceType(tt.resourceType); got != tt.want {
				t.Errorf("getActualResourceType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processRepositories(t *testing.T) {
	tests := []struct {
		name  string
		repos []string
		tag   string
		want  map[string]map[string]corerp.RecipePropertiesClassification
	}{
		{
			"Valid Repository with Redis Cache",
			[]string{
				"ghcr.io/radius-project/recipes/local-dev/rediscaches",
			},
			"0.20",
			map[string]map[string]corerp.RecipePropertiesClassification{
				"Applications.Datastores/redisCaches": {
					"default": &corerp.BicepRecipeProperties{
						TemplateKind: to.Ptr(recipes.TemplateKindBicep),
						TemplatePath: to.Ptr("ghcr.io/radius-project/recipes/local-dev/rediscaches:0.20"),
					},
				},
			},
		},
		{
			"Valid Repository with Redis Cache and Mongo Database",
			[]string{
				"ghcr.io/radius-project/recipes/local-dev/rediscaches",
				"ghcr.io/radius-project/recipes/local-dev/mongodatabases",
			},
			"0.20",
			map[string]map[string]corerp.RecipePropertiesClassification{
				"Applications.Datastores/redisCaches": {
					"default": &corerp.BicepRecipeProperties{
						TemplateKind: to.Ptr(recipes.TemplateKindBicep),
						TemplatePath: to.Ptr("ghcr.io/radius-project/recipes/local-dev/rediscaches:0.20"),
					},
				},
				"Applications.Datastores/mongoDatabases": {
					"default": &corerp.BicepRecipeProperties{
						TemplateKind: to.Ptr(recipes.TemplateKindBicep),
						TemplatePath: to.Ptr("ghcr.io/radius-project/recipes/local-dev/mongodatabases:0.20"),
					},
				},
			},
		},
		{
			"Valid Repository with Redis Cache, Mongo Database, and an unsupported type",
			[]string{
				"ghcr.io/radius-project/recipes/local-dev/rediscaches",
				"ghcr.io/radius-project/recipes/local-dev/mongodatabases",
				"ghcr.io/radius-project/recipes/local-dev/unsupported",
				"ghcr.io/radius-project/recipes/unsupported/rediscaches",
				"ghcr.io/radius-project/recipes/unsupported/unsupported",
			},
			"latest",
			map[string]map[string]corerp.RecipePropertiesClassification{
				"Applications.Datastores/redisCaches": {
					"default": &corerp.BicepRecipeProperties{
						TemplateKind: to.Ptr(recipes.TemplateKindBicep),
						TemplatePath: to.Ptr("ghcr.io/radius-project/recipes/local-dev/rediscaches:latest"),
					},
				},
				"Applications.Datastores/mongoDatabases": {
					"default": &corerp.BicepRecipeProperties{
						TemplateKind: to.Ptr(recipes.TemplateKindBicep),
						TemplatePath: to.Ptr("ghcr.io/radius-project/recipes/local-dev/mongodatabases:latest"),
					},
				},
			},
		},
		{
			"Valid Prod and Dev Repositories with Redis Cache, Mongo Database",
			[]string{
				"ghcr.io/radius-project/recipes/local-dev/rediscaches",
				"ghcr.io/radius-project/recipes/local-dev/mongodatabases",
				"ghcr.io/radius-project/dev/recipes/local-dev/rediscaches",
				"ghcr.io/radius-project/dev/recipes/local-dev/mongodatabases",
			},
			"latest",
			map[string]map[string]corerp.RecipePropertiesClassification{
				"Applications.Datastores/redisCaches": {
					"default": &corerp.BicepRecipeProperties{
						TemplateKind: to.Ptr(recipes.TemplateKindBicep),
						TemplatePath: to.Ptr("ghcr.io/radius-project/recipes/local-dev/rediscaches:latest"),
					},
				},
				"Applications.Datastores/mongoDatabases": {
					"default": &corerp.BicepRecipeProperties{
						TemplateKind: to.Ptr(recipes.TemplateKindBicep),
						TemplatePath: to.Ptr("ghcr.io/radius-project/recipes/local-dev/mongodatabases:latest"),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processRepositories(tt.repos, tt.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processRepositories() = %v, want %v", got, tt.want)
			}
		})
	}
}
