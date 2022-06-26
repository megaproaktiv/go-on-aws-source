// Example
// See https://docs.aws.amazon.com/code-samples/latest/catalog/javascriptv3-dynamodb-src-ddb_putitem.js.html

import { PutItemCommand } from "@aws-sdk/client-dynamodb";
import { ddbClient } from "./ddbClient.js";

var params = {
    Item: {
     "ID": {
       S: "NODECLIENT"
      }, 
     "Status": {
       S: "OK"
      },
    }, 
    ReturnConsumedCapacity: "TOTAL", 
    TableName: "crud"
};

export const run = async () => {
    try {
      const data = await ddbClient.send(new PutItemCommand(params));
      console.log(data);
      return data;
    } catch (err) {
      console.error(err);
    }
};
run();