package robotstxt

import (
	"testing"
)

type testCase struct {
	path, uri string
	expected bool
}


func TestStandardMatcher(t *testing.T) {
	for _, c := range getStandardCases() {
		if c.expected != standardMatcher(c.path, c.uri) {
			t.Error("path=", c.path, ", uri=", c.uri, " expected= ", c.expected)
		}
	}
}

func TestGoogleMatcher(t *testing.T) {
	cases := append(getStandardCases(), getGoogleCases()...)
	for _, c := range cases {
		if c.expected != googleMatcher(c.path, c.uri) {
			t.Error("path=", c.path, ", uri=", c.uri, " expected= ", c.expected)
		}
	}
}

func getStandardCases() []testCase {
	return []testCase {
		{ "/", "/", true },
		{ "/", "/foo", true },
		{ "/", "/foo/bar", true },

		{ "/foo", "/foo", true },
		{ "/foo", "/bar", false },
		{ "/foo", "/foo.html", true },
		{ "/foo", "/foo/", true},
		{ "/foo", "/foo/bar", true },
		{ "/foo", "/f", false },
		{ "/foo", "/", false },

		{ "/foo/", "/foo", false },
		{ "/foo/", "/bar", false },
		{ "/foo/", "/foo/", true },
		{ "/foo/", "/foo/bar", true },
		{ "/foo/", "/foo/bar/baz", true },
		{ "/foo/", "/f", false },
		{ "/foo/", "/", false },
	}
}

func getGoogleCases() []testCase {
	return []testCase {
		{ "/fish", "/fish", true },
		{ "/fish", "/fish.html", true },
		{ "/fish", "/fish/salmon.html", true },
		{ "/fish", "/fishheads", true },
		{ "/fish", "/fishheads/yummy.html", true },
		{ "/fish", "/fish.php?id=anything", true },
		{ "/fish", "/Fish.asp", false },
		{ "/fish", "/catfish", false },
		{ "/fish", "/?id=fish", false },

		{ "/fish*", "/fish", true },
		{ "/fish*", "/fish.html", true },
		{ "/fish*", "/fish/salmon.html", true },
		{ "/fish*", "/fishheads", true },
		{ "/fish*", "/fishheads/yummy.html", true },
		{ "/fish*", "/fish.php?id=anything", true },
		{ "/fish*", "/Fish.asp", false },
		{ "/fish*", "/catfish", false },
		{ "/fish*", "/?id=fish", false },

		{ "/fish/", "/fish/", true },
		{ "/fish/", "/fish/?id=anything", true },
		{ "/fish/", "/fish/salmon.htm", true },
		{ "/fish/", "/fish", false },
		{ "/fish/", "/fish.html", false },
		{ "/fish/", "/Fish/Salmon.asp", false },

		{ "/*.php", "/filename.php", true },
		{ "/*.php", "/folder/filename.php", true },
		{ "/*.php", "/folder/filename.php?parameters", true },
		{ "/*.php", "/folder/any.php.file.html", true },
		{ "/*.php", "/filename.php/", true },
		{ "/*.php", "/", false },
		{ "/*.php", "/windows.PHP", false },

		{ "/*.php$", "/filename.php", true },
		{ "/*.php$", "/folder/filename.php", true },
		{ "/*.php$", "/filename.php?parameters", false },
		{ "/*.php$", "/filename.php/", false },
		{ "/*.php$", "/filename.php5", false },
		{ "/*.php$", "/windows.PHP", false },

		{ "/fish*.php", "/fish.php", true },
		{ "/fish*.php", "/fishheads/catfish.php?parameters", true },
		{ "/fish*.php", "Fish.PHP", false },
	}
}
