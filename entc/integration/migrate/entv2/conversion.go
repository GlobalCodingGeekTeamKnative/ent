// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/migrate/entv2/conversion"
)

// Conversion is the model entity for the Conversion schema.
type Conversion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Int8ToString holds the value of the "int8_to_string" field.
	Int8ToString string `json:"int8_to_string,omitempty"`
	// Uint8ToString holds the value of the "uint8_to_string" field.
	Uint8ToString string `json:"uint8_to_string,omitempty"`
	// Int16ToString holds the value of the "int16_to_string" field.
	Int16ToString string `json:"int16_to_string,omitempty"`
	// Uint16ToString holds the value of the "uint16_to_string" field.
	Uint16ToString string `json:"uint16_to_string,omitempty"`
	// Int32ToString holds the value of the "int32_to_string" field.
	Int32ToString string `json:"int32_to_string,omitempty"`
	// Uint32ToString holds the value of the "uint32_to_string" field.
	Uint32ToString string `json:"uint32_to_string,omitempty"`
	// Int64ToString holds the value of the "int64_to_string" field.
	Int64ToString string `json:"int64_to_string,omitempty"`
	// Uint64ToString holds the value of the "uint64_to_string" field.
	Uint64ToString string `json:"uint64_to_string,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Conversion) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // int8_to_string
		&sql.NullString{}, // uint8_to_string
		&sql.NullString{}, // int16_to_string
		&sql.NullString{}, // uint16_to_string
		&sql.NullString{}, // int32_to_string
		&sql.NullString{}, // uint32_to_string
		&sql.NullString{}, // int64_to_string
		&sql.NullString{}, // uint64_to_string
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Conversion fields.
func (c *Conversion) assignValues(values ...interface{}) error {
	if m, n := len(values), len(conversion.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		c.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field int8_to_string", values[1])
	} else if value.Valid {
		c.Int8ToString = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field uint8_to_string", values[2])
	} else if value.Valid {
		c.Uint8ToString = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field int16_to_string", values[3])
	} else if value.Valid {
		c.Int16ToString = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field uint16_to_string", values[4])
	} else if value.Valid {
		c.Uint16ToString = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field int32_to_string", values[5])
	} else if value.Valid {
		c.Int32ToString = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field uint32_to_string", values[6])
	} else if value.Valid {
		c.Uint32ToString = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field int64_to_string", values[7])
	} else if value.Valid {
		c.Int64ToString = value.String
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field uint64_to_string", values[8])
	} else if value.Valid {
		c.Uint64ToString = value.String
	}
	return nil
}

// Update returns a builder for updating this Conversion.
// Note that, you need to call Conversion.Unwrap() before calling this method, if this Conversion
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Conversion) Update() *ConversionUpdateOne {
	return (&ConversionClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Conversion) Unwrap() *Conversion {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("entv2: Conversion is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Conversion) String() string {
	var builder strings.Builder
	builder.WriteString("Conversion(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", int8_to_string=")
	builder.WriteString(c.Int8ToString)
	builder.WriteString(", uint8_to_string=")
	builder.WriteString(c.Uint8ToString)
	builder.WriteString(", int16_to_string=")
	builder.WriteString(c.Int16ToString)
	builder.WriteString(", uint16_to_string=")
	builder.WriteString(c.Uint16ToString)
	builder.WriteString(", int32_to_string=")
	builder.WriteString(c.Int32ToString)
	builder.WriteString(", uint32_to_string=")
	builder.WriteString(c.Uint32ToString)
	builder.WriteString(", int64_to_string=")
	builder.WriteString(c.Int64ToString)
	builder.WriteString(", uint64_to_string=")
	builder.WriteString(c.Uint64ToString)
	builder.WriteByte(')')
	return builder.String()
}

// Conversions is a parsable slice of Conversion.
type Conversions []*Conversion

func (c Conversions) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}