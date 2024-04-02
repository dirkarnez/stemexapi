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

/*
return either
*/
func ValidUUIDExPointerFromIDString(id string) (*UUIDEx, error) {
	trimmed := strings.TrimSpace(id)
	if len(trimmed) > 1 {
		uuidParsed, err := uuid.Parse(trimmed)
		if err != nil {
			return nil, err
		} else {
			uuidex := UUIDEx(uuidParsed)
			return &uuidex, nil
		}
	} else {
		return nil, nil
	}
}
func ValidUUIDExFromIDString(id string) (UUIDEx, error) {
	trimmed := strings.TrimSpace(id)
	if len(trimmed) > 1 {
		uuidParsed, err := uuid.Parse(trimmed)
		if err != nil {
			return UUIDEx(uuid.Nil), err
		} else {
			return UUIDEx(uuidParsed), nil
		}
	} else {
		return UUIDEx(uuid.Nil), nil
	}
}

func (my UUIDEx) ToString() string {
	u := uuid.UUID(my)
	return strings.ReplaceAll(u.String(), "-", "")
}

func (my UUIDEx) IsEmpty() bool {
	empty := UUIDEx(uuid.Nil).ToString()
	return strings.TrimSpace(my.ToString()) == empty
}
