package dns

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/miekg/dns"
)

// domains for the acme challenge are like this:
// _acme-challenge.$callbackdomain.
// the first label may not be _acme-challenge though
// on disk, they are stored in $acmedir/$record/,
// which is a directory of files to return.
// example: /acme/response/_acme-challenge.hotlinecallback.net/
// this function checks if such a dir exists, and if it has files
func doesAcmeChalRespExist(domain string) bool {
	// check if an acme response directory has been configured
	if acmeChallengePath == "" {
		// no acme response, user doens't care about SSL
		// bail
		return false
	}

	// get the path for the acme challenge response, if it exists
	chalPath := getPathForAcmeChallenge(domain)

	// check if the directory exists
	_, err := os.Stat(chalPath)
	if errors.Is(err, os.ErrNotExist) {
		// dir doesn't exist, so no responses available
		return false
	}

	// the directory exists, check if there are any TXT responses available
	files, err := ioutil.ReadDir(chalPath)
	if err != nil {
		log.Println("error while looking for ACME chal responses")
		log.Println(err)
		return false
	}

	// if there are any files in the dir, there are responses available
	return len(files) > 0
}

// this function builds the directory path for looking for ACME
// challenge response files. each file is a TXT record for the domain
func getPathForAcmeChallenge(domain string) string {
	// strip off the trailing period
	domainWithoutTrailingPeriod := strings.TrimSuffix(domain, ".")

	// build the path
	p := path.Join(acmeChallengePath, domainWithoutTrailingPeriod)

	return p
}

// add the TXT record responses to the DNS message for an acme challenge
// this assumes the directory exists and has files
func setAcmeChalRRs(domain string, msg *dns.Msg) error {
	chalPath := getPathForAcmeChallenge(domain)

	files, err := ioutil.ReadDir(chalPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		fp := path.Join(chalPath, f.Name())

		file, err := os.Open(fp)
		if err != nil {
			return err
		}
		defer file.Close()

		responseBytes, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		// set the TXT response
		msg.Answer = append(msg.Answer, &dns.TXT{
			Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 60},
			Txt: []string{string(responseBytes)},
		})
	}

	return nil
}
