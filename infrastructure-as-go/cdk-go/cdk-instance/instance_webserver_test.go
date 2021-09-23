package instance_test

import (
	// Go basic
	"strings"
	"testing"
	"time"
	// aws sdk
	aws "github.com/aws/aws-sdk-go-v2/aws"
	
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	
	// cit
	ec2 "github.com/megaproaktiv/cit/citec2"
	
	// Testing
	"github.com/stretchr/testify/assert"
)


func TestPortalDNS(t *testing.T) {
	
	validate := func(i int, s string) bool{
		ok := true;
		if i != 200 {
			ok = false
		}
		if ! strings.Contains(s, "Hello World"){
			ok = false
		}

		return ok
	}


	instance,err := ec2.GetInstance(aws.String("instance"), aws.String("monolith"))

	publicIp := instance.PublicIpAddress

	url :=  "http://" + *publicIp
		
	sleepBetweenRetries, err := time.ParseDuration("10s")
	if err != nil {
		panic("Can't parse duration")
	}

	http_helper.HttpGetWithRetryWithCustomValidation(t, url, nil, 10,  sleepBetweenRetries, validate)
	

}

