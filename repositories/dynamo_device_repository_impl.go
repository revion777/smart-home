package repositories

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"smart-home/models"
	"strconv"
)

type DynamoDeviceRepositoryImpl struct {
	db *dynamodb.Client
}

func NewDynamoDeviceRepository(db *dynamodb.Client) DeviceRepository {
	return &DynamoDeviceRepositoryImpl{db: db}
}

func (repo *DynamoDeviceRepositoryImpl) CreateDevice(ctx context.Context, device *models.Device) error {
	_, err := repo.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("DevicesTable"),
		Item: map[string]types.AttributeValue{
			"id":         &types.AttributeValueMemberS{Value: device.ID},
			"mac":        &types.AttributeValueMemberS{Value: device.MAC},
			"name":       &types.AttributeValueMemberS{Value: device.Name},
			"type":       &types.AttributeValueMemberS{Value: device.Type},
			"homeId":     &types.AttributeValueMemberS{Value: device.HomeID},
			"createdAt":  &types.AttributeValueMemberN{Value: strconv.FormatInt(device.CreatedAt, 10)},
			"modifiedAt": &types.AttributeValueMemberN{Value: strconv.FormatInt(device.ModifiedAt, 10)},
		},
	})
	return err
}

func (repo *DynamoDeviceRepositoryImpl) GetDevice(ctx context.Context, id string) (*models.Device, error) {
	result, err := repo.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("DevicesTable"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	createdAt, err := parseInt64(result.Item["createdAt"])
	if err != nil {
		return nil, err
	}

	modifiedAt, err := parseInt64(result.Item["modifiedAt"])
	if err != nil {
		return nil, err
	}

	device := &models.Device{
		ID:         getString(result.Item, "id"),
		MAC:        getString(result.Item, "mac"),
		Name:       getString(result.Item, "name"),
		Type:       getString(result.Item, "type"),
		HomeID:     getString(result.Item, "homeId"),
		CreatedAt:  createdAt,
		ModifiedAt: modifiedAt,
	}
	return device, nil
}

func getString(item map[string]types.AttributeValue, key string) string {
	if value, ok := item[key].(*types.AttributeValueMemberS); ok {
		return value.Value
	}
	return ""
}

func parseInt64(attr types.AttributeValue) (int64, error) {
	if value, ok := attr.(*types.AttributeValueMemberN); ok {
		return strconv.ParseInt(value.Value, 10, 64)
	}
	return 0, fmt.Errorf("attribute is not a number")
}

func (repo *DynamoDeviceRepositoryImpl) UpdateDevice(ctx context.Context, device *models.Device) error {
	_, err := repo.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String("DevicesTable"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: device.ID},
		},
		UpdateExpression: aws.String("SET mac = :mac, name = :name, type = :type, homeId = :homeId, modifiedAt = :modifiedAt"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":mac":        &types.AttributeValueMemberS{Value: device.MAC},
			":name":       &types.AttributeValueMemberS{Value: device.Name},
			":type":       &types.AttributeValueMemberS{Value: device.Type},
			":homeId":     &types.AttributeValueMemberS{Value: device.HomeID},
			":modifiedAt": &types.AttributeValueMemberN{Value: strconv.FormatInt(device.ModifiedAt, 10)},
		},
	})
	return err
}

func (repo *DynamoDeviceRepositoryImpl) DeleteDevice(ctx context.Context, id string) error {
	_, err := repo.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String("DevicesTable"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}
