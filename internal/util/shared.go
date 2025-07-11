package util

import "regexp"

func IsValidGitURL(url string) bool {
	const pattern = `^((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)(/)?$`

	re := regexp.MustCompile(pattern)
	return re.MatchString(url)
}
