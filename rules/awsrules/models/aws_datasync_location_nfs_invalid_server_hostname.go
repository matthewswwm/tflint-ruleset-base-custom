// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsDatasyncLocationNfsInvalidServerHostnameRule checks the pattern is valid
type AwsDatasyncLocationNfsInvalidServerHostnameRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationNfsInvalidServerHostnameRule returns new rule with default attributes
func NewAwsDatasyncLocationNfsInvalidServerHostnameRule() *AwsDatasyncLocationNfsInvalidServerHostnameRule {
	return &AwsDatasyncLocationNfsInvalidServerHostnameRule{
		resourceType:  "aws_datasync_location_nfs",
		attributeName: "server_hostname",
		max:           255,
		pattern:       regexp.MustCompile(`^(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationNfsInvalidServerHostnameRule) Name() string {
	return "aws_datasync_location_nfs_invalid_server_hostname"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationNfsInvalidServerHostnameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsDatasyncLocationNfsInvalidServerHostnameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationNfsInvalidServerHostnameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationNfsInvalidServerHostnameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"server_hostname must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`server_hostname does not match valid pattern ^(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}