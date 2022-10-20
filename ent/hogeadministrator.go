// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/bug/ent/hoge"
	"entgo.io/bug/ent/hogeadministrator"
	"entgo.io/bug/ulid"
	"entgo.io/ent/dialect/sql"
)

// HogeAdministrator is the model entity for the HogeAdministrator schema.
type HogeAdministrator struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HogeAdministratorQuery when eager-loading is set.
	Edges   HogeAdministratorEdges `json:"edges"`
	hoge_id *ulid.ID
}

// HogeAdministratorEdges holds the relations/edges for other nodes in the graph.
type HogeAdministratorEdges struct {
	// Hoge holds the value of the hoge edge.
	Hoge *Hoge `json:"hoge,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HogeOrErr returns the Hoge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HogeAdministratorEdges) HogeOrErr() (*Hoge, error) {
	if e.loadedTypes[0] {
		if e.Hoge == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: hoge.Label}
		}
		return e.Hoge, nil
	}
	return nil, &NotLoadedError{edge: "hoge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HogeAdministrator) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hogeadministrator.FieldIsActive:
			values[i] = new(sql.NullBool)
		case hogeadministrator.FieldID, hogeadministrator.FieldFirstName, hogeadministrator.FieldLastName, hogeadministrator.FieldEmail:
			values[i] = new(sql.NullString)
		case hogeadministrator.FieldCreatedAt, hogeadministrator.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case hogeadministrator.ForeignKeys[0]: // hoge_id
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type HogeAdministrator", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HogeAdministrator fields.
func (ha *HogeAdministrator) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hogeadministrator.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ha.ID = ulid.ID(value.String)
			}
		case hogeadministrator.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ha.CreatedAt = value.Time
			}
		case hogeadministrator.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ha.UpdatedAt = value.Time
			}
		case hogeadministrator.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				ha.FirstName = value.String
			}
		case hogeadministrator.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				ha.LastName = value.String
			}
		case hogeadministrator.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				ha.Email = value.String
			}
		case hogeadministrator.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				ha.IsActive = value.Bool
			}
		case hogeadministrator.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hoge_id", values[i])
			} else if value.Valid {
				ha.hoge_id = new(ulid.ID)
				*ha.hoge_id = ulid.ID(value.String)
			}
		}
	}
	return nil
}

// QueryHoge queries the "hoge" edge of the HogeAdministrator entity.
func (ha *HogeAdministrator) QueryHoge() *HogeQuery {
	return (&HogeAdministratorClient{config: ha.config}).QueryHoge(ha)
}

// Update returns a builder for updating this HogeAdministrator.
// Note that you need to call HogeAdministrator.Unwrap() before calling this method if this HogeAdministrator
// was returned from a transaction, and the transaction was committed or rolled back.
func (ha *HogeAdministrator) Update() *HogeAdministratorUpdateOne {
	return (&HogeAdministratorClient{config: ha.config}).UpdateOne(ha)
}

// Unwrap unwraps the HogeAdministrator entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ha *HogeAdministrator) Unwrap() *HogeAdministrator {
	_tx, ok := ha.config.driver.(*txDriver)
	if !ok {
		panic("ent: HogeAdministrator is not a transactional entity")
	}
	ha.config.driver = _tx.drv
	return ha
}

// String implements the fmt.Stringer.
func (ha *HogeAdministrator) String() string {
	var builder strings.Builder
	builder.WriteString("HogeAdministrator(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ha.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ha.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ha.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(ha.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(ha.LastName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(ha.Email)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", ha.IsActive))
	builder.WriteByte(')')
	return builder.String()
}

// HogeAdministrators is a parsable slice of HogeAdministrator.
type HogeAdministrators []*HogeAdministrator

func (ha HogeAdministrators) config(cfg config) {
	for _i := range ha {
		ha[_i].config = cfg
	}
}
