package showtable

import (
	"context"
	"fmt"
	"net/http"

	// "fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/gin-gonic/gin"
)

type QueryInput struct {
	QueryString string `form:"querystring"`
}



func Query(c *gin.Context) {
	var input QueryInput
	c.Bind(&input)

	tableName := GetTableName(ClientSSM)
	results := QueryDDB(ClientDDB, tableName)
	
	c.HTML(http.StatusOK, "result.tmpl", gin.H{
		"Query":   input.QueryString,
		"Results": results,
	})
	
}

func QueryDDB(client *dynamodb.Client, tablename *string) []Result {
	parms := &dynamodb.ScanInput{
		TableName:                 tablename,
	}

	resp, err := client.Scan(context.TODO(), parms)
	if err != nil {
		panic("dynamodb error, " + err.Error())
	}

	items := []Result{}
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &items)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}
	
	return items
}
