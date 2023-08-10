package ruleset

import (
	"github.com/moby/buildkit/frontend/dockerfile/instructions"

	"github.com/cremindes/whalelint/parser"
)

var _ = NewRule("RUN013", "Update the package manager before installing packages.", "Package manager update ensures that the packages are up to date, regardless of when the image was built.", ValWarning,
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

	updateCommandLocation := -1
	installCommandLocation := -1

	for i, bashCommand := range bashCommandList {
		packageManager := bashCommand.Bin()
		if parser.HasPackageUpdateCommand(packageManager, bashCommand) {
			updateCommandLocation = i
		}
		if parser.IsPackageInstall(bashCommand) {
			installCommandLocation = i
		}
	}

	if (updateCommandLocation > installCommandLocation) || (updateCommandLocation == -1 && installCommandLocation != -1) {
		return RuleValidationResult{
			isViolated:    true,
			LocationRange: ParseLocationFromRawParser(bashCommandList[0].Bin(), runCommand.Location()),
		}
	}
	return RuleValidationResult{
		isViolated:    false,
		LocationRange: LocationRangeFromCommand(runCommand),
	}
}
