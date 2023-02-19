package main

import (
	"fmt"
	"net/url"
	"strings"
)

/**
* This function counts how many unique normalized valid URLs were passed to the function
*
* Accepts a list of URLs
*
* Example:
*
* input: ['https://example.com']
* output: 1
*
* Notes:
*  - assume none of the URLs have authentication information (username, password).
*
* Normalized URL:
*  - process in which a URL is modified and standardized: https://en.wikipedia.org/wiki/URL_normalization
*
#    For example.
#    These 2 urls are the same:
#    input: ["https://example.com", "https://example.com/"]
#    output: 1
#
#    These 2 are not the same:
#    input: ["https://example.com", "http://example.com"]
#    output 2
#
#    These 2 are the same:
#    input: ["https://example.com?", "https://example.com"]
#    output: 1
#
#    These 2 are the same:
#    input: ["https://example.com?a=1&b=2", "https://example.com?b=2&a=1"]
#    output: 1
*/

func CountUniqueUrls(urls []string) int {
	if len(urls) == 1 {
		return 1
	}
	parsedUrls := make(map[string]int)
	for _, u := range urls {
		parsed, err := url.Parse(u)
		if err != nil {
			fmt.Println(err)
			continue
		}
		key := fmt.Sprintf("%s://%s/%s?%s", parsed.Scheme, parsed.Host, parsed.Path, parsed.Query().Encode())
		if _, ok := parsedUrls[key]; !ok {
			parsedUrls[key] += 1
		}
	}
	return len(parsedUrls)
}

/**
 * This function counts how many unique normalized valid URLs were passed to the function per top level domain
 *
 * A top level domain is a domain in the form of example.com. Assume all top level domains end in .com
 * subdomain.example.com is not a top level domain.
 *
 * Accepts a list of URLs
 *
 * Example:
 *
 * input: ["https://example.com"]
 * output: Hash["example.com" => 1]
 *
 * input: ["https://example.com", "https://subdomain.example.com"]
 * output: Hash["example.com" => 2]
 *
 */

func CountUniqueUrlsPerTopLevelDomain(urls []string) map[string]int {
	if len(urls) == 1 {
		return map[string]int{urls[0]: 1}
	}
	parsedUrls := make(map[string]int)
	for _, u := range urls {
		parsed, err := url.Parse(u)
		if err != nil {
			fmt.Println(err)
			continue
		}
		keys := strings.Split(parsed.Hostname(), ".")
		key := fmt.Sprintf("%s.%s", keys[len(keys)-2], keys[len(keys)-1])

		if _, ok := parsedUrls[key]; !ok {
			parsedUrls[key] += 1
		}
	}
	return parsedUrls
}

func main() {
	urls := []string{
		"https://example.com/a?a=b&b=c",
		"https://example.com/a?b=c&a=b",
		"https://example.com/b",
		"https://example.com",
		"https://subdomain.example.com",
		"https://subdomain1.example.com",
	}
	fmt.Println(CountUniqueUrls(urls))
	fmt.Println(CountUniqueUrlsPerTopLevelDomain(urls))
}
