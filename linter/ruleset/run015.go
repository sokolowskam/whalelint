package ruleset

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
)

var _ = NewRule("RUN015", "Do not hardcode the private key in the Dockerfile.", "", ValError,
	ValidateRun015)

func ValidateRun015(runCommand *instructions.RunCommand) RuleValidationResult {
	result := RuleValidationResult{
		isViolated:    false,
		LocationRange: LocationRangeFromCommand(runCommand),
	}

	regexpInvalidPattern := regexp.MustCompile(`PRIVATE KEY-----`)
	if match := regexpInvalidPattern.FindString(runCommand.String()); len(match) > 0 {
		result.SetViolated()
		result.LocationRange = ParseLocationFromRawParser(match, runCommand.Location())
	}

	return result
}
