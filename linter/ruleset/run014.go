package ruleset

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
)

var _ = NewRule("RUN014", "Do not hardcode credentials. Consider passing vulnerable data via environment variables or secret files.", "", ValWarning,
	ValidateRun014)

func ValidateRun014(runCommand *instructions.RunCommand) RuleValidationResult {
	result := RuleValidationResult{
		isViolated:    false,
		LocationRange: LocationRangeFromCommand(runCommand),
	}

	regexpInvalidPattern := regexp.MustCompile(`(?i)(password|token|client_secret)`)
	if match := regexpInvalidPattern.FindString(runCommand.String()); len(match) > 0 {
		result.SetViolated()
		result.LocationRange = ParseLocationFromRawParser(match, runCommand.Location())
	}

	return result
}
