package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

type Student struct {
	PersonID   uuid.UUID  `gorm:"primaryKey"`
	Person     Person     `json:",omitempty"`
	Index      string     `json:"index"`
	Major      string     `json:"major"`
	RandomData RandomData `json:"randomData" gorm:"type:jsonb;"`
}

type RandomData struct {
	Years int `json:"years"`
}

func (r RandomData) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *RandomData) Scan(value interface{}) error {
	if value == nil {
		*r = RandomData{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, r)
}
