package robotstxt

import urlLib "net/url"

func getRobotsUrl(url *urlLib.URL) string {
	return  url.Scheme + "://" + url.Host + "/robots.txt"
}

func getEffectiveAgent(userAgent string, data *robotsData) string {
	if _, found := data.disallowRules[userAgent]; found {
		return userAgent
	} else if _, found := data.allowRules[userAgent]; found {
		return userAgent
	} else {
		return "*"
	}
}
