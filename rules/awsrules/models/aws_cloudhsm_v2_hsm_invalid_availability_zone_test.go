// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsCloudhsmV2HsmInvalidAvailabilityZoneRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudhsm_v2_hsm" "foo" {
	availability_zone = "us-east-1"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudhsmV2HsmInvalidAvailabilityZoneRule(),
					Message: `availability_zone does not match valid pattern ^[a-z]{2}(-(gov))?-(east|west|north|south|central){1,2}-\d[a-z]$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudhsm_v2_hsm" "foo" {
	availability_zone = "us-east-1a"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudhsmV2HsmInvalidAvailabilityZoneRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
