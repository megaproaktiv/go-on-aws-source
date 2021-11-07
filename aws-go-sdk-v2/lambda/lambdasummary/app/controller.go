package lambdasummary

import (
	"context"
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

func Collect(accounts []string, regions []string) ([]*LambdaSummary, error) {

	var summaries []*LambdaSummary

	for _, account := range accounts {
		cfgSub, err := GetCfgSub(stsclient, account)
		if err != nil {
			log.Warn("Not able to assume account role: ", account)
		} else {
			for _, region := range regions {
			
				
				client := lambda.NewFromConfig(cfgSub)
				collection := ListFunctions(client, &account, &region)
				
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
					summaries = append(summaries, &counter)
			}
		}
	}

	return summaries, nil

}
