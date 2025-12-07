package repositories

import (
	"log"

	"github.com/CaioDS/fli/internal/domain/models"
	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type VersionsRepository struct {
	dbContext *context.DbContext
}

func NewVersionsRepository(dbContext *context.DbContext) *VersionsRepository {
	return &VersionsRepository{
		dbContext: dbContext,
	}
}

func (v *VersionsRepository) GetVersion(versionNumber string) (*models.Version, error) {
	attributes := map[string]types.AttributeValue{
			":version": &types.AttributeValueMemberS{Value: versionNumber},
	}
	result, err := v.dbContext.Query(
		"sdk_versions",
		*aws.String("version = :version"), 
		attributes,
		nil,
	)

	log.Println("AOOO", result, len(result), err)

	if err != nil || len(result) == 0 {
		log.Fatalln("Failed to retrieve available versions:", err)
		return nil, err
	}

	var parsedResults []models.Version
	attributevalue.UnmarshalListOfMaps(result, &parsedResults)

	return &parsedResults[0], nil
}