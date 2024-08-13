package models

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"strconv"
)

type Device struct {
	ID         string
	MAC        string
	Name       string
	Type       string
	HomeID     string
	CreatedAt  int64
	ModifiedAt int64
}

func (d *Device) ToItemWithId() map[string]types.AttributeValue {
	resMap := d.ToItem()
	resMap["id"] = &types.AttributeValueMemberS{Value: d.ID}
	return resMap
}

func (d *Device) ToItem() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"mac":        &types.AttributeValueMemberS{Value: d.MAC},
		"name":       &types.AttributeValueMemberS{Value: d.Name},
		"type":       &types.AttributeValueMemberS{Value: d.Type},
		"homeId":     &types.AttributeValueMemberS{Value: d.HomeID},
		"createdAt":  &types.AttributeValueMemberN{Value: strconv.FormatInt(d.CreatedAt, 10)},
		"modifiedAt": &types.AttributeValueMemberN{Value: strconv.FormatInt(d.ModifiedAt, 10)},
	}
}

func (d *Device) FromItem(item map[string]types.AttributeValue) (*Device, error) {
	createdAt, err := parseInt64FromAttrValue(item["createdAt"])
	if err != nil {
		return nil, err
	}
	modifiedAt, err := parseInt64FromAttrValue(item["modifiedAt"])
	if err != nil {
		return nil, err
	}

	return &Device{
		ID:         getStringFromAttrValue(item, "id"),
		MAC:        getStringFromAttrValue(item, "mac"),
		Name:       getStringFromAttrValue(item, "name"),
		Type:       getStringFromAttrValue(item, "type"),
		HomeID:     getStringFromAttrValue(item, "homeId"),
		CreatedAt:  createdAt,
		ModifiedAt: modifiedAt,
	}, nil
}

func getStringFromAttrValue(item map[string]types.AttributeValue, key string) string {
	if value, ok := item[key].(*types.AttributeValueMemberS); ok {
		return value.Value
	}
	return ""
}

func parseInt64FromAttrValue(attr types.AttributeValue) (int64, error) {
	if value, ok := attr.(*types.AttributeValueMemberN); ok {
		return strconv.ParseInt(value.Value, 10, 64)
	}
	return 0, fmt.Errorf("attribute is not a number")
}
