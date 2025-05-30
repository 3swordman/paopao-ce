// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// BooleanFalse is a Boolean of type False.
	// false
	BooleanFalse Boolean = iota
	// BooleanTrue is a Boolean of type True.
	// true
	BooleanTrue
)

var ErrInvalidBoolean = fmt.Errorf("not a valid Boolean, try [%s]", strings.Join(_BooleanNames, ", "))

const _BooleanName = "FalseTrue"

var _BooleanNames = []string{
	_BooleanName[0:5],
	_BooleanName[5:9],
}

// BooleanNames returns a list of possible string values of Boolean.
func BooleanNames() []string {
	tmp := make([]string, len(_BooleanNames))
	copy(tmp, _BooleanNames)
	return tmp
}

// BooleanValues returns a list of the values for Boolean
func BooleanValues() []Boolean {
	return []Boolean{
		BooleanFalse,
		BooleanTrue,
	}
}

var _BooleanMap = map[Boolean]string{
	BooleanFalse: _BooleanName[0:5],
	BooleanTrue:  _BooleanName[5:9],
}

// String implements the Stringer interface.
func (x Boolean) String() string {
	if str, ok := _BooleanMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Boolean(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Boolean) IsValid() bool {
	_, ok := _BooleanMap[x]
	return ok
}

var _BooleanValue = map[string]Boolean{
	_BooleanName[0:5]:                  BooleanFalse,
	strings.ToLower(_BooleanName[0:5]): BooleanFalse,
	_BooleanName[5:9]:                  BooleanTrue,
	strings.ToLower(_BooleanName[5:9]): BooleanTrue,
}

// ParseBoolean attempts to convert a string to a Boolean.
func ParseBoolean(name string) (Boolean, error) {
	if x, ok := _BooleanValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _BooleanValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Boolean(0), fmt.Errorf("%s is %w", name, ErrInvalidBoolean)
}

// MustParseBoolean converts a string to a Boolean, and panics if is not valid.
func MustParseBoolean(name string) Boolean {
	val, err := ParseBoolean(name)
	if err != nil {
		panic(err)
	}
	return val
}

// MarshalText implements the text marshaller method.
func (x Boolean) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Boolean) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseBoolean(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errBooleanNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *Boolean) Scan(value interface{}) (err error) {
	if value == nil {
		*x = Boolean(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = Boolean(v)
	case string:
		*x, err = ParseBoolean(v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(v); verr == nil {
				*x, err = Boolean(val), nil
			}
		}
	case []byte:
		*x, err = ParseBoolean(string(v))
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(string(v)); verr == nil {
				*x, err = Boolean(val), nil
			}
		}
	case Boolean:
		*x = v
	case int:
		*x = Boolean(v)
	case *Boolean:
		if v == nil {
			return errBooleanNilPtr
		}
		*x = *v
	case uint:
		*x = Boolean(v)
	case uint64:
		*x = Boolean(v)
	case *int:
		if v == nil {
			return errBooleanNilPtr
		}
		*x = Boolean(*v)
	case *int64:
		if v == nil {
			return errBooleanNilPtr
		}
		*x = Boolean(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = Boolean(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errBooleanNilPtr
		}
		*x = Boolean(*v)
	case *uint:
		if v == nil {
			return errBooleanNilPtr
		}
		*x = Boolean(*v)
	case *uint64:
		if v == nil {
			return errBooleanNilPtr
		}
		*x = Boolean(*v)
	case *string:
		if v == nil {
			return errBooleanNilPtr
		}
		*x, err = ParseBoolean(*v)
		if err != nil {
			// try parsing the integer value as a string
			if val, verr := strconv.Atoi(*v); verr == nil {
				*x, err = Boolean(val), nil
			}
		}
	}

	return
}

// Value implements the driver Valuer interface.
func (x Boolean) Value() (driver.Value, error) {
	return int64(x), nil
}

type NullBoolean struct {
	Boolean Boolean
	Valid   bool
	Set     bool
}

func NewNullBoolean(val interface{}) (x NullBoolean) {
	x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	return
}

// Scan implements the Scanner interface.
func (x *NullBoolean) Scan(value interface{}) (err error) {
	x.Set = true
	if value == nil {
		x.Boolean, x.Valid = Boolean(0), false
		return
	}

	err = x.Boolean.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x NullBoolean) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	// driver.Value accepts int64 for int values.
	return int64(x.Boolean), nil
}

// MarshalJSON correctly serializes a NullBoolean to JSON.
func (n NullBoolean) MarshalJSON() ([]byte, error) {
	const nullStr = "null"
	if n.Valid {
		return json.Marshal(n.Boolean)
	}
	return []byte(nullStr), nil
}

// UnmarshalJSON correctly deserializes a NullBoolean from JSON.
func (n *NullBoolean) UnmarshalJSON(b []byte) error {
	n.Set = true
	var x interface{}
	err := json.Unmarshal(b, &x)
	if err != nil {
		return err
	}
	err = n.Scan(x)
	return err
}
