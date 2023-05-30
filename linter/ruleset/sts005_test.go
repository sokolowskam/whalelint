// nolint:dupl
package ruleset_test

import (
	"testing"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/stretchr/testify/assert"

	RuleSet "github.com/cremindes/whalelint/linter/ruleset"
)

func TestValidateSts005(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		IsViolation   bool
		StageBaseName string
		ExampleName   string
		DocsContext   string
	}{
		{
			IsViolation:   false,
			ExampleName:   "One stage FROM ubuntu:20.04.",
			StageBaseName: "ubuntu:20.04",
			DocsContext:   "`FROM` {{ .StageBaseName }}",
		},
		{
			IsViolation:   true,
			ExampleName:   "One stage FROM ubuntu:latest.",
			StageBaseName: "test/ubuntu:20.04",
			DocsContext:   "`FROM` {{ .StageBaseName }}",
		},
		{
			IsViolation:   false,
			ExampleName:   "One stage FROM ubuntu;20.04.",
			StageBaseName: "docker.io/ubuntu:20.04",
			DocsContext:   "`FROM` {{ .StageBaseName }}",
		},
		{
			IsViolation:   false,
			ExampleName:   "One stage FROM ubuntu",
			StageBaseName: "hub.docker.com/ubuntu:20.04",
			DocsContext:   "`FROM` {{ .StageBaseName }}",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.ExampleName, func(t *testing.T) {
			t.Parallel()

			// nolint:exhaustivestruct
			stage := instructions.Stage{BaseName: testCase.StageBaseName, SourceCode: "FROM " + testCase.StageBaseName}

			assert.Equal(t, testCase.IsViolation, RuleSet.ValidateSts005(stage).IsViolated())
		})
	}
}
