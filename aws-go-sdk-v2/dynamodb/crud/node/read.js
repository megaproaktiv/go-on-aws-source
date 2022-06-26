// Example
// See https://docs.aws.amazon.com/code-samples/latest/catalog/javascriptv3-dynamodb-src-ddb_putitem.js.html

import { GetItemCommand } from "@aws-sdk/client-dynamodb";
import { ddbClient } from "./ddbClient.js";

var params = {
  TableName: "crud",
  Key: {
      ID: { S: "NODECLIENT" }, 
  }, 
};

export const run = async () => {
    try {
      const data = await ddbClient.send(new GetItemCommand(params));
      // console.log(data);
      // Item: { ID: { S: 'NODECLIENT' }, Status: { S: 'OK' } }
      console.log("Status is: ",data.Item.Status.S);
      return data;
    } catch (err) {
      console.error(err);
    }
  };
  run();