package robotstxt

import urlLib "net/url"

type AccessType bool
const (
	Allowed AccessType = true
	Disallowed AccessType = false
)

type ExclusionStandardType int
const (
	RobotsExclusionStandard ExclusionStandardType = iota
	GoogleExclusionStandard ExclusionStandardType = iota
)

func CheckAccess(url, userAgent string) (AccessType, error) {
	return defaultRobotsTxtInstance.CheckAccess(url, userAgent, GoogleExclusionStandard)
}

func (r *RobotsTxt) CheckAccess(url, userAgent string, standard ExclusionStandardType) (AccessType, error) {
	u, err := urlLib.Parse(url)
	if err != nil {
		return Disallowed, err
	}

	robotsUrl := getRobotsUrl(u)
	robots, found := r.sites[robotsUrl]
	if !found {
		r.fetchAndParse(robotsUrl)
		robots, found = r.sites[robotsUrl]
	}

	return robots.checkAccess(u.RequestURI(), userAgent, standard)
}


type RobotsTxt struct {
	sites map [string]*robotsData // robotx.txt URL -> robots data
}

func New() *RobotsTxt {
	return &RobotsTxt{make(map[string]*robotsData)}
}

type robotsData struct {
	disallowRules map [string][]string // User-agent -> list of rules
	allowRules     map [string][]string // User-agent -> list of rules
}

func (r *robotsData) checkAccess(uri, userAgent string, standard ExclusionStandardType) (AccessType, error) {
	matcher := getMatcher(standard)
	access := Allowed

	effectiveAgent := getEffectiveAgent(userAgent, r)
	for _, rule := range r.disallowRules[effectiveAgent] {
		if matcher(rule, uri) {
			access = Disallowed
			break
		}
	}

	if access == Disallowed {
		for _, rule := range r.allowRules[effectiveAgent] {
			if matcher(rule, uri) {
				access = Allowed
				break
			}
		}
	}

	return access, nil
}

var defaultRobotsTxtInstance *RobotsTxt = New()

