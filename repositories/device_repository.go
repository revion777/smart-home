package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"smart-home/models"
	"strconv"
)

var db *dynamodb.Client

func InitDynamoDB(client *dynamodb.Client) {
	db = client
}

func CreateDevice(device models.Device) (models.Device, error) {
	_, err := db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("Devices"),
		Item: map[string]types.AttributeValue{
			"ID":         &types.AttributeValueMemberS{Value: device.ID},
			"MAC":        &types.AttributeValueMemberS{Value: device.MAC},
			"Name":       &types.AttributeValueMemberS{Value: device.Name},
			"Type":       &types.AttributeValueMemberS{Value: device.Type},
			"HomeID":     &types.AttributeValueMemberS{Value: device.HomeID},
			"CreatedAt":  &types.AttributeValueMemberN{Value: strconv.FormatInt(device.CreatedAt, 10)},
			"ModifiedAt": &types.AttributeValueMemberN{Value: strconv.FormatInt(device.ModifiedAt, 10)},
		},
	})
	if err != nil {
		return models.Device{}, err
	}
	return device, nil
}

func GetDevice(id string) (models.Device, error) {
	result, err := db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("Devices"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return models.Device{}, err
	}
	if result.Item == nil {
		return models.Device{}, nil // No item found
	}

	createdAt, err := strconv.ParseInt(result.Item["CreatedAt"].(*types.AttributeValueMemberN).Value, 10, 64)
	if err != nil {
		return models.Device{}, err
	}

	modifiedAt, err := strconv.ParseInt(result.Item["ModifiedAt"].(*types.AttributeValueMemberN).Value, 10, 64)
	if err != nil {
		return models.Device{}, err
	}

	return models.Device{
		ID:         result.Item["ID"].(*types.AttributeValueMemberS).Value,
		MAC:        result.Item["MAC"].(*types.AttributeValueMemberS).Value,
		Name:       result.Item["Name"].(*types.AttributeValueMemberS).Value,
		Type:       result.Item["Type"].(*types.AttributeValueMemberS).Value,
		HomeID:     result.Item["HomeID"].(*types.AttributeValueMemberS).Value,
		CreatedAt:  createdAt,
		ModifiedAt: modifiedAt,
	}, nil
}

func UpdateDevice(id string, device models.Device) (models.Device, error) {
	_, err := db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String("Devices"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("set MAC = :mac, Name = :name, Type = :type, HomeID = :homeID, ModifiedAt = :modifiedAt"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":mac":        &types.AttributeValueMemberS{Value: device.MAC},
			":name":       &types.AttributeValueMemberS{Value: device.Name},
			":type":       &types.AttributeValueMemberS{Value: device.Type},
			":homeID":     &types.AttributeValueMemberS{Value: device.HomeID},
			":modifiedAt": &types.AttributeValueMemberN{Value: strconv.FormatInt(device.ModifiedAt, 10)},
		},
	})
	if err != nil {
		return models.Device{}, err
	}
	return device, nil
}

func DeleteDevice(id string) error {
	_, err := db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Devices"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}
