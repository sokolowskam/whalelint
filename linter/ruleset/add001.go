package ruleset

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"net/url"
	"path"
	"regexp"
)

var _ = NewRule("ADD001", "Use COPY for copying local files into the an image.",
	"Use 'COPY' instead of 'ADD' if the ADD functionalities like - adding files from an URL or from a compressed file is needed.",
	ValWarning, ValidateAdd001)

func ValidateAdd001(addCommand *instructions.AddCommand) RuleValidationResult {
	archiveExtensionList := []string{
		".7z", ".gz", ".lz", "lzo", "lzma", ".tar", ".tb2", ".tbz", ".tbz2", ".tgz",
		".tlz", ".tpz", ".txz", ".tZ", "zx", ".Z", ".zip",
	}

	result := RuleValidationResult{
		isViolated:    false,
		LocationRange: LocationRangeFromCommand(addCommand),
	}

	// ADD command to add remote resources will be evaluated as valid.
	_, err := url.ParseRequestURI(addCommand.SourcesAndDest.SourcePaths[0])

	if err == nil {
		u, err2 := url.Parse(addCommand.SourcesAndDest.SourcePaths[0])
		if err2 != nil || u.Scheme == "" || u.Host == "" {
			fmt.Println("not an url")
		} else {
			return result
		}
	}

	// command to add a private git repository will be evaluated as valid.
	// https://docs.docker.com/engine/reference/builder/#adding-a-private-git-repository
	// TODO: make the regex more precise
	regexpWrongNumberOfDashViolation := regexp.MustCompile(`git@.{1,}`)
	if regexpWrongNumberOfDashViolation.MatchString(addCommand.String()) {
		return result
	}

	// ADD command to extract files will be evaluated as valid.
	fileExt := path.Ext(addCommand.SourcesAndDest.SourcePaths[0])
	for _, archiveExt := range archiveExtensionList {
		if fileExt == archiveExt {
			return result

		}
	}

	// else set violated
	result.SetViolated()
	result.LocationRange.start.charNumber = 0
	result.LocationRange.end.charNumber = len("ADD")
	return result
}
