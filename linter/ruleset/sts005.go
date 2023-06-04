package ruleset

import (
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"strings"

	"github.com/cremindes/whalelint/utils"
)

// STS -> Stage Single.
var _ = NewRule("STS005", "Make sure you are using a trusted registry. Avoid using images of unknown origin and from users you do not trust. When in doubt, use docker official images.", "", ValInfo, ValidateSts005)

func ValidateSts005(stage instructions.Stage) RuleValidationResult {
	result := RuleValidationResult{isViolated: false, LocationRange: BKRangeSliceToLocationRange(stage.Location)}

	if strings.Contains(stage.BaseName, "/") {
		registry, _ := utils.SplitKeyValue(stage.BaseName, '/')
		if registry != "docker.io" && registry != "hub.docker.com" {
			result.SetViolated() // as latest is the default
		}
		if result.IsViolated() {
			result.message = "Registry \"" + registry + "\" might not be trusted."
			result.LocationRange = ParseLocationFromRawParser(stage.BaseName, stage.Location)
		}
	}

	return result
}
