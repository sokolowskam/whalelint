package ruleset

import (
	"github.com/moby/buildkit/frontend/dockerfile/instructions"

	"github.com/cremindes/whalelint/parser"
)

// TODO: revisit

var _ = NewRule("RUN013", "Update the package manager.", "", ValWarning,
	ValidateRun013)

func ValidateRun013(runCommand *instructions.RunCommand) RuleValidationResult {
	result := RuleValidationResult{
		isViolated:    false,
		LocationRange: LocationRangeFromCommand(runCommand),
	}

	if len(runCommand.CmdLine) == 0 {
		return result
	} else if len(runCommand.CmdLine[0]) == 0 {
		return result
	}

	bashCommandList := parser.ParseBashCommandChain(runCommand).BashCommandList

	//TODO: Add checking if package manager is even used

	for _, bashCommand := range bashCommandList {
		packageManager := bashCommand.Bin()
		if parser.HasPackageUpdateCommand(packageManager, bashCommand) {
			return RuleValidationResult{
				isViolated:    false,
				LocationRange: LocationRangeFromCommand(runCommand),
			}
		}
	}
	return RuleValidationResult{
		isViolated:    true,
		LocationRange: ParseLocationFromRawParser(bashCommandList[0].Bin(), runCommand.Location()),
	}
}
