// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsEfsFileSystemInvalidPerformanceModeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_efs_file_system" "foo" {
	performance_mode = "minIO"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsEfsFileSystemInvalidPerformanceModeRule(),
					Message: `performance_mode is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_efs_file_system" "foo" {
	performance_mode = "generalPurpose"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsEfsFileSystemInvalidPerformanceModeRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
