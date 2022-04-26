// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/k3forx/coffee_memo/pkg/ent/driprecipe"
)

// DripRecipe is the model entity for the DripRecipe schema.
type DripRecipe struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int32 `json:"user_id,omitempty"`
	// CoffeeBeanID holds the value of the "coffee_bean_id" field.
	CoffeeBeanID int32 `json:"coffee_bean_id,omitempty"`
	// CoffeeBeanWeight holds the value of the "coffee_bean_weight" field.
	CoffeeBeanWeight float64 `json:"coffee_bean_weight,omitempty"`
	// LiquidWeight holds the value of the "liquid_weight" field.
	LiquidWeight float64 `json:"liquid_weight,omitempty"`
	// Temperature holds the value of the "temperature" field.
	Temperature float64 `json:"temperature,omitempty"`
	// SteamTime holds the value of the "steam_time" field.
	SteamTime int32 `json:"steam_time,omitempty"`
	// DripTime holds the value of the "drip_time" field.
	DripTime int32 `json:"drip_time,omitempty"`
	// Memo holds the value of the "memo" field.
	Memo string `json:"memo,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DripRecipe) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case driprecipe.FieldCoffeeBeanWeight, driprecipe.FieldLiquidWeight, driprecipe.FieldTemperature:
			values[i] = new(sql.NullFloat64)
		case driprecipe.FieldID, driprecipe.FieldUserID, driprecipe.FieldCoffeeBeanID, driprecipe.FieldSteamTime, driprecipe.FieldDripTime:
			values[i] = new(sql.NullInt64)
		case driprecipe.FieldMemo:
			values[i] = new(sql.NullString)
		case driprecipe.FieldCreatedAt, driprecipe.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DripRecipe", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DripRecipe fields.
func (dr *DripRecipe) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case driprecipe.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dr.ID = int32(value.Int64)
		case driprecipe.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				dr.UserID = int32(value.Int64)
			}
		case driprecipe.FieldCoffeeBeanID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coffee_bean_id", values[i])
			} else if value.Valid {
				dr.CoffeeBeanID = int32(value.Int64)
			}
		case driprecipe.FieldCoffeeBeanWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field coffee_bean_weight", values[i])
			} else if value.Valid {
				dr.CoffeeBeanWeight = value.Float64
			}
		case driprecipe.FieldLiquidWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field liquid_weight", values[i])
			} else if value.Valid {
				dr.LiquidWeight = value.Float64
			}
		case driprecipe.FieldTemperature:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field temperature", values[i])
			} else if value.Valid {
				dr.Temperature = value.Float64
			}
		case driprecipe.FieldSteamTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field steam_time", values[i])
			} else if value.Valid {
				dr.SteamTime = int32(value.Int64)
			}
		case driprecipe.FieldDripTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field drip_time", values[i])
			} else if value.Valid {
				dr.DripTime = int32(value.Int64)
			}
		case driprecipe.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				dr.Memo = value.String
			}
		case driprecipe.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dr.CreatedAt = value.Time
			}
		case driprecipe.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dr.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DripRecipe.
// Note that you need to call DripRecipe.Unwrap() before calling this method if this DripRecipe
// was returned from a transaction, and the transaction was committed or rolled back.
func (dr *DripRecipe) Update() *DripRecipeUpdateOne {
	return (&DripRecipeClient{config: dr.config}).UpdateOne(dr)
}

// Unwrap unwraps the DripRecipe entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dr *DripRecipe) Unwrap() *DripRecipe {
	tx, ok := dr.config.driver.(*txDriver)
	if !ok {
		panic("ent: DripRecipe is not a transactional entity")
	}
	dr.config.driver = tx.drv
	return dr
}

// String implements the fmt.Stringer.
func (dr *DripRecipe) String() string {
	var builder strings.Builder
	builder.WriteString("DripRecipe(")
	builder.WriteString(fmt.Sprintf("id=%v", dr.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", dr.UserID))
	builder.WriteString(", coffee_bean_id=")
	builder.WriteString(fmt.Sprintf("%v", dr.CoffeeBeanID))
	builder.WriteString(", coffee_bean_weight=")
	builder.WriteString(fmt.Sprintf("%v", dr.CoffeeBeanWeight))
	builder.WriteString(", liquid_weight=")
	builder.WriteString(fmt.Sprintf("%v", dr.LiquidWeight))
	builder.WriteString(", temperature=")
	builder.WriteString(fmt.Sprintf("%v", dr.Temperature))
	builder.WriteString(", steam_time=")
	builder.WriteString(fmt.Sprintf("%v", dr.SteamTime))
	builder.WriteString(", drip_time=")
	builder.WriteString(fmt.Sprintf("%v", dr.DripTime))
	builder.WriteString(", memo=")
	builder.WriteString(dr.Memo)
	builder.WriteString(", created_at=")
	builder.WriteString(dr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(dr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DripRecipes is a parsable slice of DripRecipe.
type DripRecipes []*DripRecipe

func (dr DripRecipes) config(cfg config) {
	for _i := range dr {
		dr[_i].config = cfg
	}
}
