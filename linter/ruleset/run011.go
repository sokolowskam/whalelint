package ruleset

import (
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	log "github.com/sirupsen/logrus"

	Parser "github.com/cremindes/whalelint/parser"
)

var _ = NewRule("RUN011", "Generating keys in the build phase renders the keys static.", "Consider mounting the keys at runtime instead.", ValWarning,
	ValidateRun011)

func ValidateRun011(runCommand *instructions.RunCommand) RuleValidationResult {

	result := RuleValidationResult{
		isViolated:    false,
		LocationRange: LocationRangeFromCommand(runCommand),
	}
	result.LocationRange.end.charNumber += len(runCommand.ShellDependantCmdLine.CmdLine)

	bashCommandChain := Parser.ParseBashCommandChain(runCommand.CmdLine)

	for _, bashCommand := range bashCommandChain.BashCommandList {
		if bashCommand.Bin() == "ssh-keygen" {
			result.SetViolated()
		}
	}

	log.Trace("ValidateRun011 result:", result)

	return result
}
