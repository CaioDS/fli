package context

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DbContext struct {
	client *dynamodb.Client
}

func NewDbContext(endpoint string, region string) (*DbContext, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		func(o *config.LoadOptions) error {
			o.Region = region
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	return &DbContext{
		client: dynamodb.NewFromConfig(
			cfg, 
			func(o *dynamodb.Options) {
				o.BaseEndpoint = &endpoint
			},
		),
	}, nil
}

func (db *DbContext) Get(table string) ([]map[string]types.AttributeValue, error) {
	out, err := db.client.Scan(
		context.TODO(),
		&dynamodb.ScanInput{
			TableName: aws.String(table),
		},
	)
	if err != nil {
		log.Fatalln("Failed to retrieve data from table", err)
		return nil, errors.New("failed to retrieve data from table")
	}

	return out.Items, nil
}

func (db *DbContext) Query(
	table string, 
	keyConditionExpression string,
	attributes map[string]types.AttributeValue, 
	filterExpression *string,
) ([]map[string]types.AttributeValue, error) {
	result, err := db.client.Query(
		context.TODO(),
		&dynamodb.QueryInput{
			TableName: aws.String(table),
			KeyConditionExpression: &keyConditionExpression,
			ExpressionAttributeValues: attributes,
			FilterExpression: filterExpression,
		},
	)
	if err != nil {
		return nil, errors.New("Failed to execute query")
	}

	return result.Items, nil
}