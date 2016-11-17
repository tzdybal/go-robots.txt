package robotstxt

import "testing"

func TestParseLine(t *testing.T) {
	r := New()
	r.currentUrl = "http://example.org"
	r.currentAgent = "Example Agent"
	r.sites[r.currentUrl] = &robotsData{map[string][]string{}, map[string][]string{}}

	r.parseLine("User-agent: *")
	if r.currentAgent != "*" {
		t.Error("User-agent directive not working");
	}
	r.parseLine("Disallow: /foo")
	rules := r.sites[r.currentUrl].disallowRules
	if len(rules) != 1 {
		t.Error("Disallow directive: Invalid number of agents")
	} else if len(rules["*"]) != 1 {
		t.Error("Disallow directive: Invalid number of rules")
	} else if rules["*"][0] != "/foo" {
		t.Error("Disallow directive: Invalid path")
	}

	r.parseLine("Allow: /foo/bar")
	rules = r.sites[r.currentUrl].allowRules
	if len(rules) != 1 {
		t.Error("Allow directive: Invalid number of agents")
	} else if len(rules["*"]) != 1 {
		t.Error("Allow directive: Invalid number of rules")
	} else if rules["*"][0] != "/foo/bar" {
		t.Error("Allow directive: Invalid path")
	}
}
