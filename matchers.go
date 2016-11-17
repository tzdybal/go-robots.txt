package robotstxt

import (
	"strings"
	"regexp"
)

type matcherFunction func (string, string) bool

func getMatcher(standard ExclusionStandardType) matcherFunction {
	switch standard {
	case RobotsExclusionStandard:
		return standardMatcher
	case GoogleExclusionStandard:
		return googleMatcher
	}
	return standardMatcher
}

func standardMatcher(path, uri string) bool {
	return strings.HasPrefix(uri, path)
}

func googleMatcher(path, uri string) bool {
	pattern := "^" + strings.Replace(path, "*", ".*", -1)
	match, err := regexp.MatchString(pattern, uri)
	if err == nil {
		return match
	} else {
		return false
	}
}
