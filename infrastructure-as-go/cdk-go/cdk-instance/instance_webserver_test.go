package instance_test

import (
	// Go basic
	"strings"
	"testing"
	"time"

	// aws sdk
	aws "github.com/aws/aws-sdk-go-v2/aws"

	// begin terratest
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	// end terratest

	// begin import
	// cit
	ec2 "github.com/megaproaktiv/cit/citec2"
	// end import
	// Testing
)

func TestPortalDNS(t *testing.T) {
	// begin validate
	validate := func(i int, s string) bool {
		ok := true
		if i != 200 {
			ok = false
		}
		if !strings.Contains(s, "Hello World") {
			ok = false
		}

		return ok
	}
	//end validate
	//begin getip
	instance, err := ec2.GetInstance(aws.String("instance"), aws.String("monolith"))
	publicIp := instance.PublicIpAddress
	//end getip

	// begin itest
	url := "http://" + *publicIp
	sleepBetweenRetries, err := time.ParseDuration("10s")
	if err != nil {
		panic("Can't parse duration")
	}
	http_helper.HttpGetWithRetryWithCustomValidation(t, url, nil, 10, sleepBetweenRetries, validate)
	// end itest
}
