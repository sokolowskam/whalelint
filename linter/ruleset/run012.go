package ruleset

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
)

var _ = NewRule("RUN012", "Avoid interacting with /etc/sudoers file. Consider using gosu to enforce root instead.", "", ValWarning,
	ValidateRun012)

func ValidateRun012(runCommand *instructions.RunCommand) RuleValidationResult {
	result := RuleValidationResult{
		isViolated:    false,
		LocationRange: LocationRangeFromCommand(runCommand),
	}

	regexpInvalidPattern := regexp.MustCompile(`/etc/sudoers`)
	if match := regexpInvalidPattern.FindString(runCommand.String()); len(match) > 0 {
		result.SetViolated()
		result.LocationRange = ParseLocationFromRawParser(match, runCommand.Location())
	}

	return result
}
