// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/k3forx/coffee_memo/pkg/ent/goosedbversion"
)

// GooseDbVersion is the model entity for the GooseDbVersion schema.
type GooseDbVersion struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// VersionID holds the value of the "version_id" field.
	VersionID int `json:"version_id,omitempty"`
	// IsApplied holds the value of the "is_applied" field.
	IsApplied bool `json:"is_applied,omitempty"`
	// Tstamp holds the value of the "tstamp" field.
	Tstamp time.Time `json:"tstamp,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GooseDbVersion) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goosedbversion.FieldIsApplied:
			values[i] = new(sql.NullBool)
		case goosedbversion.FieldID, goosedbversion.FieldVersionID:
			values[i] = new(sql.NullInt64)
		case goosedbversion.FieldTstamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GooseDbVersion", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GooseDbVersion fields.
func (gdv *GooseDbVersion) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goosedbversion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gdv.ID = uint64(value.Int64)
		case goosedbversion.FieldVersionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version_id", values[i])
			} else if value.Valid {
				gdv.VersionID = int(value.Int64)
			}
		case goosedbversion.FieldIsApplied:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_applied", values[i])
			} else if value.Valid {
				gdv.IsApplied = value.Bool
			}
		case goosedbversion.FieldTstamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field tstamp", values[i])
			} else if value.Valid {
				gdv.Tstamp = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GooseDbVersion.
// Note that you need to call GooseDbVersion.Unwrap() before calling this method if this GooseDbVersion
// was returned from a transaction, and the transaction was committed or rolled back.
func (gdv *GooseDbVersion) Update() *GooseDbVersionUpdateOne {
	return (&GooseDbVersionClient{config: gdv.config}).UpdateOne(gdv)
}

// Unwrap unwraps the GooseDbVersion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gdv *GooseDbVersion) Unwrap() *GooseDbVersion {
	tx, ok := gdv.config.driver.(*txDriver)
	if !ok {
		panic("ent: GooseDbVersion is not a transactional entity")
	}
	gdv.config.driver = tx.drv
	return gdv
}

// String implements the fmt.Stringer.
func (gdv *GooseDbVersion) String() string {
	var builder strings.Builder
	builder.WriteString("GooseDbVersion(")
	builder.WriteString(fmt.Sprintf("id=%v", gdv.ID))
	builder.WriteString(", version_id=")
	builder.WriteString(fmt.Sprintf("%v", gdv.VersionID))
	builder.WriteString(", is_applied=")
	builder.WriteString(fmt.Sprintf("%v", gdv.IsApplied))
	builder.WriteString(", tstamp=")
	builder.WriteString(gdv.Tstamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// GooseDbVersions is a parsable slice of GooseDbVersion.
type GooseDbVersions []*GooseDbVersion

func (gdv GooseDbVersions) config(cfg config) {
	for _i := range gdv {
		gdv[_i].config = cfg
	}
}
