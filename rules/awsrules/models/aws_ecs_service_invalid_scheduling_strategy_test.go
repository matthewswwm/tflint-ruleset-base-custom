// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsEcsServiceInvalidSchedulingStrategyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ecs_service" "foo" {
	scheduling_strategy = "SERVER"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsEcsServiceInvalidSchedulingStrategyRule(),
					Message: `scheduling_strategy is not a valid value`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ecs_service" "foo" {
	scheduling_strategy = "REPLICA"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsEcsServiceInvalidSchedulingStrategyRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
