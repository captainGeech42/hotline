package dns

import (
	"errors"
	"os"
	"path"
	"strings"
)

// domains for the acme challenge are like this:
// _acme-challenge.$callbackdomain.
// the first label may not be _acme-challenge though
// on disk, they are stored in $acmedir/$callbackdomain/$record,
// which is a file with the contents to return.
// this function checks if such a response exists
func doesAcmeChalRespExist(domain string) bool {
	// check if an acme response directory has been configured
	if acmeChallengePath == "" {
		// no acme response, user doens't care about SSL
		// bail
		return false
	}

	// get the path for the acme challenge response, if it exists
	chalPath := getPathForAcmeChallenge(domain)

	// check if the file exists
	_, err := os.Stat(chalPath)
	return !errors.Is(err, os.ErrNotExist)
}

func getPathForAcmeChallenge(domain string) string {
	// trim off the callback domain
	// this could leave multiple labels still in the prefix,
	// but that doesn't matter for the usage of this
	prefix := strings.TrimSuffix(domain, callbackDomain+".")

	return path.Join(acmeChallengePath, callbackDomain, prefix)
}
