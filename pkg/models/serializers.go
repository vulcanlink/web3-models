package models

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/schema"
)

type StringerJoinerSerializer struct{}
type ByteArraySerializer struct{}

var StringArrReflectType reflect.Type = reflect.TypeOf("string")
var HashArrReflectType reflect.Type = reflect.TypeOf([]common.Hash{})
var AddressArrReflectType reflect.Type = reflect.TypeOf([]common.Address{})

// Assumes field a slice
func (StringerJoinerSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	fieldArr := reflect.ValueOf(fieldValue)
	fieldArrLen := fieldArr.Len()

	strArr := make([]string, fieldArrLen)
	for i := 0; i < fieldArrLen; i++ {
		elem := fieldArr.Index(i).Interface()
		strArr[i] = fmt.Sprint(elem)
	}

	strJoin := strings.Join(strArr, ",")
	return strJoin, nil
}

// Assumes field is a slice
func (StringerJoinerSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	fieldValue := reflect.New(field.FieldType)

	if nil == dbValue {
		field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
		return nil
	}

	var dbValueStr string
	switch v := dbValue.(type) {
	case string:
		dbValueStr = v
	case []byte:
		dbValueStr = string(v)
	default:
		return fmt.Errorf("failed to unmarshal text value of field %s: %#v", field.Name, dbValue)
	}

	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())

	dbValueStrArr := strings.Split(dbValueStr, ",")

	// fieldValue should be of type *[]interface{}
	fieldElemType := fieldValue.Type().Elem()

	switch fieldElemType {
	case HashArrReflectType:
		hashArr := make([]common.Hash, len(dbValueStrArr))
		for i, hashStr := range dbValueStrArr {
			// TODO: Fix if necessary with clickhouse
			hashArr[i] = common.HexToHash(hashStr)
		}

		fieldValue = reflect.ValueOf(hashArr)
	case AddressArrReflectType:
		addressArr := make([]common.Address, len(dbValueStrArr))
		for i, addressStr := range dbValueStrArr {
			// TODO: Fix if necessary with clickhouse
			addressArr[i] = common.HexToAddress(addressStr)
		}

		fieldValue = reflect.ValueOf(addressArr)
	case StringArrReflectType:
		fieldValue = reflect.ValueOf(dbValueStrArr)
	default:
		return fmt.Errorf("destination type not supported by StringerJoinerSerializer: %T", dst.Type())
	}

	// We know the field is a slice, so we Set it directly to fieldValue, and not a pointer to fieldValue
	field.ReflectValueOf(ctx, dst).Set(fieldValue)
	return
}

func (StringerJoinerSerializer) String() string {
	return "stringer_joiner"
}
