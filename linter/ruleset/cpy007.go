package ruleset

import (
	"regexp"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
)

var _ = NewRule(
	"CPY007", "Beware of copying sensitive information | COPY . ... can end up copying sensitive information to the Docker image",
	``+"`COPY . ...`"+` command can end up copying sensitive information to the Docker image. 
If you have files containing sensitive information in your folder, either remove them or add them to .dockerignore.`,
	ValWarning, ValidateCpy007)

// checks for recursive copies of the form "COPY ."
func ValidateCpy007(copyCommand *instructions.CopyCommand) RuleValidationResult {
	result := RuleValidationResult{
		isViolated:    false,
		LocationRange: LocationRangeFromCommand(copyCommand),
	}

	// can't use lookahead, lookbehind as it is not supported in Go's regexp.
	regexpWrongNumberOfDashViolation := regexp.MustCompile(`[^\.][ ]{1}(\.)[ ]{1}`)
	if regexpWrongNumberOfDashViolation.MatchString(copyCommand.String()) {
		result.SetViolated()
		result.message = "Beware of copying sensitive information."

		wrongFlagStr := regexpWrongNumberOfDashViolation.FindString(copyCommand.SourcesAndDest.SourcePaths[0])
		result.LocationRange = ParseLocationFromRawParser(wrongFlagStr, copyCommand.Location())
	}

	return result
}
