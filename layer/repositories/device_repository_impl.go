package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"smart-home/layer/models"
)

type DeviceRepositoryImpl struct {
	db *dynamodb.Client
}

func NewDeviceRepository(db *dynamodb.Client) DeviceRepository {
	return &DeviceRepositoryImpl{db: db}
}

func (repo *DeviceRepositoryImpl) CreateDevice(ctx context.Context, device *models.Device) error {
	_, err := repo.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("devices"),
		Item:      device.ToItemWithId(),
	})
	return err
}

func (repo *DeviceRepositoryImpl) GetDevice(ctx context.Context, id string) (*models.Device, error) {
	result, err := repo.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("devices"),
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

	var device *models.Device
	resDevice, err := (*models.Device).FromItem(device, result.Item)
	if err != nil {
		return nil, err
	}

	return resDevice, nil
}

func (repo *DeviceRepositoryImpl) UpdateDevice(ctx context.Context, device *models.Device) error {
	_, err := repo.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String("devices"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: device.ID},
		},
		UpdateExpression:          aws.String("SET mac = :mac, name = :name, type = :type, homeId = :homeId, modifiedAt = :modifiedAt"),
		ExpressionAttributeValues: device.ToItem(),
	})
	return err
}

func (repo *DeviceRepositoryImpl) DeleteDevice(ctx context.Context, id string) error {
	_, err := repo.db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String("devices"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}
