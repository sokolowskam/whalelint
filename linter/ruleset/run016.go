package ruleset

import (
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	log "github.com/sirupsen/logrus"
	"regexp"

	Parser "github.com/cremindes/whalelint/parser"
)

var _ = NewRule("RUN016", "Consider checking the checksum of every file you download.", "", ValWarning,
	ValidateRun016)

func ValidateRun016(runCommand *instructions.RunCommand) RuleValidationResult {
	downloadCmdSet := []string{"curl", "wget", "torrent"}

	result := RuleValidationResult{
		isViolated:    false,
		LocationRange: LocationRangeFromCommand(runCommand),
	}
	result.LocationRange.end.charNumber += len(runCommand.ShellDependantCmdLine.CmdLine)

	regexpChecksumPattern := regexp.MustCompile(`(sha1sum|sha224sum|sha384sum|sha256sum|sha512sum|md5sum|hmac|base64)`)

	bashCommandChain := Parser.ParseBashCommandChain(runCommand.CmdLine)
	bashCommandList := bashCommandChain.BashCommandList

	for i, bashCommand := range bashCommandList {
		for _, downloadCmd := range downloadCmdSet {
			if bashCommand.Bin() == downloadCmd {
				hasChecksum := false
				matched := ""
				for j := i; j < len(bashCommandList); j++ {
					if match := regexpChecksumPattern.FindString(bashCommandList[j].String()); len(match) > 0 {
						hasChecksum = true
						matched = match
					}
				}
				if !hasChecksum {
					result.SetViolated()
					result.LocationRange = ParseLocationFromRawParser(matched, runCommand.Location())
				}
			}
		}
	}

	log.Trace("ValidateRun016 result:", result)

	return result
}
