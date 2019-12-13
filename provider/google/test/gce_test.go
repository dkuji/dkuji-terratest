package test

import (
	"os"
	"path"
  "testing"
  "fmt"

	"github.com/gruntwork-io/terratest/modules/terraform"
	terraformCore "github.com/hashicorp/terraform/terraform"
)

func TestUT_StorageAccountName(t *testing.T) {
	t.Parallel()

	// Test cases for storage account name conversion logic
	testCases := map[string]string{
    "test": "test1",
    /*
		"TestWebsiteName": "testwebsitenamedata001",
		"ALLCAPS":         "allcapsdata001",
		"S_p-e(c)i.a_l":   "specialdata001",
		"A1phaNum321":     "a1phanum321data001",
    "E5e-y7h_ng":      "e5ey7hngdata001",
    */
	}

	//for input, expected := range testCases {
	for expected := range testCases {
		// Specify the test case folder and "-var" options
		tfOptions := &terraform.Options{
      TerraformDir: "../",
      /*
			Vars: map[string]interface{}{
				"website_name": input,
      },
      */
		}

		// Terraform init and plan only
		tfPlanOutput := "terraform.tfplan"
		terraform.Init(t, tfOptions)
		terraform.RunTerraformCommand(t, tfOptions, terraform.FormatArgs(tfOptions, "plan", "-out="+tfPlanOutput)...)

		// Read and parse the plan output
		f, err := os.Open(path.Join(tfOptions.TerraformDir, tfPlanOutput))
    fmt.Println(f)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
    plan, err := terraformCore.ReadPlan(f)
    fmt.Println("aaaaaaaaaaaaaaaaa")
    fmt.Println(err)
		if err != nil {
			t.Fatal(err)
		}

		// Validate the test result
		for _, mod := range plan.Diff.Modules {
			if len(mod.Path) == 2 && mod.Path[0] == "root" && mod.Path[1] == "google" {
				actual := mod.Resources["google_compute_instance.default"].Attributes["name"].New
				if actual != expected {
					t.Fatalf("Expect %v, but found %v", expected, actual)
				}
			}
		}
	}
}
