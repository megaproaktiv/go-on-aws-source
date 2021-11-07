package lambdasummary

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	log "github.com/sirupsen/logrus"
)

var stsclient *sts.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	stsclient = sts.NewFromConfig(cfg)
}

var wg sync.WaitGroup

func Collect(accounts []string, regions []string) ([]*LambdaSummary, error) {

	var summaries []*LambdaSummary

	for _, account := range accounts {
		cfgSub, err := GetCfgSub(stsclient, account)
		if err != nil {
			log.Warn("Not able to assume account role: ", account)
		} else {
			for _, region := range regions {
				log.Info("Account: ",account, " Region: ", region)
				go func(mysummaries *[]*LambdaSummary, account string, region string) {
					defer wg.Done()

					client := lambda.NewFromConfig(cfgSub)
					collection := ListFunctions(client, account, region)

					counter := LambdaSummary{
						Account:        account,
						Region:         region,
						RuntimeCounter: map[string]int{},
					}
					for _, i := range collection {
						_, ok := counter.RuntimeCounter[i.Runtime]
						if ok {
							counter.RuntimeCounter[i.Runtime] = counter.RuntimeCounter[i.Runtime] + 1
						} else {
							counter.RuntimeCounter[i.Runtime] = 1
						}

					}
					*mysummaries = append(*mysummaries, &counter)
				}(&summaries, account, region)
				wg.Add(1)

			}
		}
	}
	wg.Wait()

	return summaries, nil

}
