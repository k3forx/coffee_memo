// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeebean"
	"github.com/k3forx/coffee_memo/pkg/ent/user"
	"github.com/k3forx/coffee_memo/pkg/ent/userscoffeebean"
)

// UsersCoffeeBean is the model entity for the UsersCoffeeBean schema.
type UsersCoffeeBean struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status int32 `json:"status,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int32 `json:"user_id,omitempty"`
	// CoffeeBeanID holds the value of the "coffee_bean_id" field.
	CoffeeBeanID int32 `json:"coffee_bean_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UsersCoffeeBeanQuery when eager-loading is set.
	Edges UsersCoffeeBeanEdges `json:"edges"`
}

// UsersCoffeeBeanEdges holds the relations/edges for other nodes in the graph.
type UsersCoffeeBeanEdges struct {
	// CoffeeBean holds the value of the coffee_bean edge.
	CoffeeBean *CoffeeBean `json:"coffee_bean,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CoffeeBeanOrErr returns the CoffeeBean value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UsersCoffeeBeanEdges) CoffeeBeanOrErr() (*CoffeeBean, error) {
	if e.loadedTypes[0] {
		if e.CoffeeBean == nil {
			// The edge coffee_bean was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: coffeebean.Label}
		}
		return e.CoffeeBean, nil
	}
	return nil, &NotLoadedError{edge: "coffee_bean"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UsersCoffeeBeanEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UsersCoffeeBean) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userscoffeebean.FieldID, userscoffeebean.FieldStatus, userscoffeebean.FieldUserID, userscoffeebean.FieldCoffeeBeanID:
			values[i] = new(sql.NullInt64)
		case userscoffeebean.FieldCreatedAt, userscoffeebean.FieldUpdatedAt, userscoffeebean.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UsersCoffeeBean", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UsersCoffeeBean fields.
func (ucb *UsersCoffeeBean) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userscoffeebean.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ucb.ID = int32(value.Int64)
		case userscoffeebean.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ucb.Status = int32(value.Int64)
			}
		case userscoffeebean.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ucb.UserID = int32(value.Int64)
			}
		case userscoffeebean.FieldCoffeeBeanID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coffee_bean_id", values[i])
			} else if value.Valid {
				ucb.CoffeeBeanID = int32(value.Int64)
			}
		case userscoffeebean.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ucb.CreatedAt = value.Time
			}
		case userscoffeebean.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ucb.UpdatedAt = value.Time
			}
		case userscoffeebean.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ucb.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// QueryCoffeeBean queries the "coffee_bean" edge of the UsersCoffeeBean entity.
func (ucb *UsersCoffeeBean) QueryCoffeeBean() *CoffeeBeanQuery {
	return (&UsersCoffeeBeanClient{config: ucb.config}).QueryCoffeeBean(ucb)
}

// QueryUser queries the "user" edge of the UsersCoffeeBean entity.
func (ucb *UsersCoffeeBean) QueryUser() *UserQuery {
	return (&UsersCoffeeBeanClient{config: ucb.config}).QueryUser(ucb)
}

// Update returns a builder for updating this UsersCoffeeBean.
// Note that you need to call UsersCoffeeBean.Unwrap() before calling this method if this UsersCoffeeBean
// was returned from a transaction, and the transaction was committed or rolled back.
func (ucb *UsersCoffeeBean) Update() *UsersCoffeeBeanUpdateOne {
	return (&UsersCoffeeBeanClient{config: ucb.config}).UpdateOne(ucb)
}

// Unwrap unwraps the UsersCoffeeBean entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ucb *UsersCoffeeBean) Unwrap() *UsersCoffeeBean {
	tx, ok := ucb.config.driver.(*txDriver)
	if !ok {
		panic("ent: UsersCoffeeBean is not a transactional entity")
	}
	ucb.config.driver = tx.drv
	return ucb
}

// String implements the fmt.Stringer.
func (ucb *UsersCoffeeBean) String() string {
	var builder strings.Builder
	builder.WriteString("UsersCoffeeBean(")
	builder.WriteString(fmt.Sprintf("id=%v", ucb.ID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", ucb.Status))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", ucb.UserID))
	builder.WriteString(", coffee_bean_id=")
	builder.WriteString(fmt.Sprintf("%v", ucb.CoffeeBeanID))
	builder.WriteString(", created_at=")
	builder.WriteString(ucb.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ucb.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(ucb.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UsersCoffeeBeans is a parsable slice of UsersCoffeeBean.
type UsersCoffeeBeans []*UsersCoffeeBean

func (ucb UsersCoffeeBeans) config(cfg config) {
	for _i := range ucb {
		ucb[_i].config = cfg
	}
}
