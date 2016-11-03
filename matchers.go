package robotstxt

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
	return false;
}

func googleMatcher(path, uri string) bool {
	return false;
}
