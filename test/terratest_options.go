package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func createTestGKEClusterTerraformOptions(
	t *testing.T,
	uniqueID,
	project string,
	region string,
	templatePath string,
) *terraform.Options {
	repoName := strings.ToLower(fmt.Sprintf("cloud-build-repo-%s", uniqueID))
	gkeClusterName := strings.ToLower(fmt.Sprintf("gke-cluster-%s", uniqueID))
	gkeServiceAccountName := strings.ToLower(fmt.Sprintf("gke-cluster-sa-%s", uniqueID))

	terraformVars := map[string]interface{}{
		"region":                       region,
		"location":                     region,
		"project":                      project,
		"gcr_region":                   lookupMultiRegion(region),
		"repository_name":              repoName,
		"cluster_name":                 gkeClusterName,
		"cluster_service_account_name": gkeServiceAccountName,
	}

	terratestOptions := terraform.Options{
		TerraformDir: templatePath,
		Vars:         terraformVars,
	}

	return &terratestOptions
}

func createTestGitHubTerraformOptions(
	t *testing.T,
	uniqueID,
	project string,
	region string,
	githubOrg string,
	githubRepo string,
	templatePath string,
) *terraform.Options {
	gkeClusterName := strings.ToLower(fmt.Sprintf("gke-cluster-%s", uniqueID))
	gkeServiceAccountName := strings.ToLower(fmt.Sprintf("gke-cluster-sa-%s", uniqueID))

	terraformVars := map[string]interface{}{
		"region":                       region,
		"location":                     region,
		"project":                      project,
		"gcr_region":                   lookupMultiRegion(region),
		"github_owner":                 githubOrg,
		"github_repository":            githubRepo,
		"cluster_name":                 gkeClusterName,
		"cluster_service_account_name": gkeServiceAccountName,
	}

	terratestOptions := terraform.Options{
		TerraformDir: templatePath,
		Vars:         terraformVars,
	}

	return &terratestOptions
}

// lookupMultiRegion returns the appropriate multi-region depending on the GCP region passed in.
// https://cloud.google.com/storage/docs/locations#location-mr
func lookupMultiRegion(region string) string {
	parts := strings.Split(region, "-")

	switch mr := parts[0]; mr {
	case "europe":
		return "eu"
	case "asia":
		return "asia"
	default:
		return "us"
	}
}