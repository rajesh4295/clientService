package common

import (
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertDateToGRPCTimestamp(datetimeStr string) *timestamppb.Timestamp {
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, datetimeStr)
	if err != nil {
		fmt.Println("Error parsing datetime:", err)
		return nil
	}
	return timestamppb.New(parsedTime)
}

func ConvertGRPCTimestampToDate(timestamp *timestamppb.Timestamp) string {
	timeObj := timestamp.AsTime()
	layout := "2006-01-02T15:04:05.999999999Z"
	datetimeStr := timeObj.Format(layout)
	return datetimeStr
}

func ConvertProtoFieldsToBson(fields map[string]*structpb.Value) (bson.M, error) {
	bsonMap := bson.M{}

	for key, value := range fields {
		convertedValue, err := convertProtoValueToGoValue(value)
		if err != nil {
			return nil, err
		}
		bsonMap[key] = convertedValue
	}
	log.Printf("fields %v", fields)
	log.Printf("bsonMap %v", bsonMap)
	return bsonMap, nil
}

// convertProtoValueToGoValue converts *structpb.Value to its corresponding Go type
func convertProtoValueToGoValue(value *structpb.Value) (interface{}, error) {
	switch kind := value.Kind.(type) {
	case *structpb.Value_NullValue:
		return nil, nil
	case *structpb.Value_BoolValue:
		return kind.BoolValue, nil
	case *structpb.Value_NumberValue:
		return kind.NumberValue, nil
	case *structpb.Value_StringValue:
		return kind.StringValue, nil
	case *structpb.Value_ListValue:
		// Convert each item in the list
		list := make([]interface{}, len(kind.ListValue.Values))
		for i, v := range kind.ListValue.Values {
			convertedValue, err := convertProtoValueToGoValue(v)
			if err != nil {
				return nil, err
			}
			list[i] = convertedValue
		}
		return list, nil
	case *structpb.Value_StructValue:
		// Recursively convert the Struct fields
		return ConvertProtoFieldsToBson(kind.StructValue.Fields)
	default:
		return nil, fmt.Errorf("unsupported Protobuf value type: %T", value.Kind)
	}
}
