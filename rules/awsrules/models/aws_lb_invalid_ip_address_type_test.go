// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsLbInvalidIPAddressTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_lb" "foo" {
	ip_address_type = "ipv6"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsLbInvalidIPAddressTypeRule(),
					Message: `ip_address_type is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_lb" "foo" {
	ip_address_type = "ipv4"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsLbInvalidIPAddressTypeRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
