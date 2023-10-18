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
	"context"
	"fmt"
	"strings"

	corerp "github.com/radius-project/radius/pkg/corerp/api/v20231001preview"
	ext_ctrl "github.com/radius-project/radius/pkg/corerp/frontend/controller/extenders"
	dapr_ctrl "github.com/radius-project/radius/pkg/daprrp/frontend/controller"
	ds_ctrl "github.com/radius-project/radius/pkg/datastoresrp/frontend/controller"
	msg_ctrl "github.com/radius-project/radius/pkg/messagingrp/frontend/controller"
	recipe_types "github.com/radius-project/radius/pkg/recipes"
	"github.com/radius-project/radius/pkg/to"
	"github.com/radius-project/radius/pkg/ucp/ucplog"
	"github.com/radius-project/radius/pkg/version"
	"oras.land/oras-go/v2/registry/remote"
)

const (
	// RecipeRepositoryPrefix is the prefix for the repository path.
	RecipeRepositoryPrefix = "ghcr.io/radius-project/recipes/local-dev/"
)

// availableDevRecipes returns the list of available dev recipes.
//
// If we want to add a new recipe, we need to add it here.
func availableDevRecipes() []string {
	return []string{
		"mongodatabases",
		"rediscaches",
		"sqldatabases",
		"rabbitmqqueues",
		"pubsubbrokers",
		"secretstores",
		"statestores",
		"extenders",
	}
}

//go:generate mockgen -destination=./mock_devrecipeclient.go -package=radinit -self_package github.com/radius-project/radius/pkg/cli/cmd/radinit github.com/radius-project/radius/pkg/cli/cmd/radinit DevRecipeClient
type DevRecipeClient interface {
	GetDevRecipes(ctx context.Context) (map[string]map[string]corerp.RecipePropertiesClassification, error)
}

type devRecipeClient struct {
}

// NewDevRecipeClient creates a new DevRecipeClient object and returns it.
func NewDevRecipeClient() DevRecipeClient {
	return &devRecipeClient{}
}

// GetDevRecipes is a function that queries a registry for recipes with a specific tag and returns a map of recipes.
// If an error occurs, an error is returned.
func (drc *devRecipeClient) GetDevRecipes(ctx context.Context) (map[string]map[string]corerp.RecipePropertiesClassification, error) {
	logger := ucplog.FromContextOrDiscard(ctx)

	// The tag will be the major.minor version of the release.
	tag := version.Channel()
	if version.IsEdgeChannel() {
		tag = "latest"
	}

	validRepos := []string{}
	for _, recipe := range availableDevRecipes() {
		repoPath := fmt.Sprintf("%s%s", RecipeRepositoryPrefix, recipe)
		repo, err := remote.NewRepository(repoPath)
		if err != nil {
			// This shouldn't cancel the `rad init` flow.
			logger.Error(err, fmt.Sprintf("failed to create client to repository %s", repoPath))
			continue
		}

		// Setting this is not the best way to go.
		repo.TagListPageSize = 1000
		tagExists := false
		err = repo.Tags(ctx, "", func(tags []string) error {
			for _, t := range tags {
				if t == tag {
					tagExists = true
					break
				}
			}
			return nil
		})
		if err != nil {
			continue
		}

		if tagExists {
			validRepos = append(validRepos, repoPath)
		}
	}

	return processRepositories(validRepos, tag), nil
}

// processRepositories processes the repositories and returns the recipes.
func processRepositories(repos []string, tag string) map[string]map[string]corerp.RecipePropertiesClassification {
	recipes := map[string]map[string]corerp.RecipePropertiesClassification{}

	// We are using the default recipe.
	recipeName := "default"

	for _, repo := range repos {
		// An example to a normalized resource type is "mongodatabases".
		// The actual resource type is "Applications.Datastores/mongoDatabases".
		normalizedResourceType := getNormalizedResourceTypeFromPath(repo)
		// If the normalized resource type is empty, it means we don't support the repository.
		if normalizedResourceType == "" {
			continue
		}

		resourceType := getActualResourceType(normalizedResourceType)
		// If the actual resource type is empty, it means we don't support the resource type.
		if resourceType == "" {
			continue
		}

		recipes[resourceType] = map[string]corerp.RecipePropertiesClassification{
			recipeName: &corerp.BicepRecipeProperties{
				TemplateKind: to.Ptr(recipe_types.TemplateKindBicep),
				TemplatePath: to.Ptr(repo + ":" + tag),
			},
		}
	}

	return recipes
}

// getNormalizedResourceTypeFromPath parses the repository path to extract the resource type.
//
// Should be of the form: ghcr.io/radius-project/recipes/local-dev/<resourceType>.
//
// An example to a normalized resource type is "mongodatabases".
func getNormalizedResourceTypeFromPath(repo string) (resourceType string) {
	_, after, found := strings.Cut(repo, RecipeRepositoryPrefix)
	if !found || after == "" {
		return ""
	}

	if strings.Count(after, "/") == 0 {
		resourceType = strings.Split(after, "/")[0]
	}

	return resourceType
}

// getActualResourceType returns the resource type for the given resource.
//
// An example to an actual resource type is "Applications.Datastores/mongoDatabases".
func getActualResourceType(resourceType string) string {
	switch resourceType {
	case "mongodatabases":
		return ds_ctrl.MongoDatabasesResourceType
	case "rediscaches":
		return ds_ctrl.RedisCachesResourceType
	case "sqldatabases":
		return ds_ctrl.SqlDatabasesResourceType
	case "rabbitmqqueues":
		return msg_ctrl.RabbitMQQueuesResourceType
	case "pubsubbrokers":
		return dapr_ctrl.DaprPubSubBrokersResourceType
	case "secretstores":
		return dapr_ctrl.DaprSecretStoresResourceType
	case "statestores":
		return dapr_ctrl.DaprStateStoresResourceType
	case "extenders":
		return ext_ctrl.ResourceTypeName
	default:
		return ""
	}
}
