package dbutil

import (
	"encoding/json"
	"github.com/jinzhu/gorm/dialects/postgres"
)

func EmptyJsonArray() postgres.Jsonb {
	return postgres.Jsonb{RawMessage: json.RawMessage(`[]`)}
}
