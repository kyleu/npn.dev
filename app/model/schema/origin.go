package schema

import "encoding/json"

type Origin struct {
	Key         string `json:"key"`
	T       string `json:"t"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var OriginGraphQL = Origin{Key: "graphql", T: "graphql", Title: "GraphQL", Description: "GraphQL schema and queries"}
var OriginProtobuf = Origin{Key: "protobuf", T: "protobuf", Title: "Protobuf", Description: "File describing proto3 definitions"}
var OriginIntelliJ = Origin{Key: "intellij", T: "database", Title: "Database", Description: "Database system and supporting queries"}
var OriginLiquibase = Origin{Key: "liquibase", T: "database", Title: "Database", Description: "Database system and supporting queries"}
var OriginUnknown = Origin{Key: "unknown", T: "unknown", Title: "Unknown", Description: "Not quite sure what this is"}

var AllOrigins = []Origin{OriginGraphQL, OriginProtobuf, OriginIntelliJ, OriginLiquibase}

func OriginFromString(s string) Origin {
	for _, t := range AllOrigins {
		if t.Key == s {
			return t
		}
	}
	return OriginUnknown
}

func (t *Origin) String() string {
	return t.Key
}

func (t *Origin) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Origin) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*t = OriginFromString(s)
	return nil
}
