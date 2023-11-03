package model

import (
	"database/sql/driver"
	"fmt"

	"strings"

	"github.com/google/uuid"
)

type UUIDEx uuid.UUID

// MarshalJSON marshals the UUIDEx type to a JSON UUID string.
func (my UUIDEx) MarshalJSON() ([]byte, error) {
	u := uuid.UUID(my)
	return []byte(fmt.Sprintf(`"%s"`, strings.ReplaceAll(u.String(), "-", ""))), nil
}

// GormDataType -> sets type to binary(16)
func (my UUIDEx) GormDataType() string {
	return "binary(16)"
}

// Scan --> From DB
func (my *UUIDEx) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = UUIDEx(parseByte)
	return err
}

// Value -> TO DB
func (my UUIDEx) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}

func NewUUIDEx() UUIDEx {
	return UUIDEx(uuid.New())
}

func ValidUUIDExFromIDString(id string) (UUIDEx, error) {
	uuidParsed, err := uuid.Parse(id)
	if err != nil {
		return NewUUIDEx(), err
	}
	return UUIDEx(uuidParsed), nil
}
