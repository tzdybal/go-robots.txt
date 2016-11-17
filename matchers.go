package robotstxt

import "strings"

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
	// TODO: fix this
	return standardMatcher(path, uri)
}
