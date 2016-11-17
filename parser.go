package robotstxt

import (
	"bufio"
	"net/http"
	"strings"
)

func (r *RobotsTxt) fetchAndParse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return;
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	switch {
	case 400 <= status && status < 500:
		// full allow
		r.sites[url] = &robotsData{map[string][]string{}, map[string][]string{}}
	case 500 <= status && status < 600:
		// full disallow
		r.sites[url] = &robotsData{map[string][]string{"*":{"/"}}, map[string][]string{}}
	case 200 <= status && status < 300:
		// conditional allow
		r.currentUrl = url
		r.sites[url] = &robotsData{map[string][]string{}, map[string][]string{}}
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			r.parseLine(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			// TODO: handle error
		}

		r.currentUrl = ""
	}

}

func (r *RobotsTxt) parseLine(line string) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	scanner.Scan();
	directive := scanner.Text()
	scanner.Scan();

	switch directive {
	case "User-agent:":
		r.currentAgent = scanner.Text()
	case "Disallow:":
		rules := r.sites[r.currentUrl].disallowRules
		rules[r.currentAgent] = append(rules[r.currentAgent], scanner.Text())
	case "Allow:":
		rules := r.sites[r.currentUrl].allowRules
		rules[r.currentAgent] = append(rules[r.currentAgent], scanner.Text())
	}
}
