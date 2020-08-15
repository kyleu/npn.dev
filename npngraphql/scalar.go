package npngraphql

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func serializeUUID(value interface{}) interface{} {
	switch value := value.(type) {
	case uuid.UUID:
		buff, err := value.MarshalText()
		if err != nil {
			return nil
		}
		return string(buff)
	case *uuid.UUID:
		if value == nil {
			return nil
		}
		return serializeUUID(*value)
	default:
		return nil
	}
}

func unserializeUUID(value interface{}) interface{} {
	switch value := value.(type) {
	case []byte:
		u := uuid.UUID{}
		err := u.UnmarshalText(value)
		if err != nil {
			return nil
		}

		return u
	case string:
		return unserializeUUID([]byte(value))
	case *string:
		if value == nil {
			return nil
		}
		return unserializeUUID([]byte(*value))
	case time.Time:
		return value
	default:
		return nil
	}
}

var ScalarUUID = graphql.NewScalar(graphql.ScalarConfig{
	Name: "UUID",
	Description: "The `UUID` scalar type represents a 128-bit value, represented as hex-encoded " +
		"character sequences, such as [00000000-0000-0000-0123-456789ABCDEF].",
	Serialize:  serializeUUID,
	ParseValue: unserializeUUID,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			ret, err := uuid.FromString(valueAST.Value)
			if err != nil {
				return nil
			}
			return ret
		}
		return nil
	},
})
