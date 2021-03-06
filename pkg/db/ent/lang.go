// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/lang"
	"github.com/google/uuid"
)

// Lang is the model entity for the Lang schema.
type Lang struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Lang holds the value of the "lang" field.
	Lang string `json:"lang,omitempty"`
	// Short holds the value of the "short" field.
	Short string `json:"short,omitempty"`
	// Logo holds the value of the "logo" field.
	Logo string `json:"logo,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Lang) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case lang.FieldCreateAt, lang.FieldUpdateAt, lang.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case lang.FieldLang, lang.FieldShort, lang.FieldLogo, lang.FieldName:
			values[i] = new(sql.NullString)
		case lang.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Lang", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Lang fields.
func (l *Lang) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lang.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case lang.FieldLang:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lang", values[i])
			} else if value.Valid {
				l.Lang = value.String
			}
		case lang.FieldShort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short", values[i])
			} else if value.Valid {
				l.Short = value.String
			}
		case lang.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				l.Logo = value.String
			}
		case lang.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case lang.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				l.CreateAt = uint32(value.Int64)
			}
		case lang.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				l.UpdateAt = uint32(value.Int64)
			}
		case lang.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				l.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Lang.
// Note that you need to call Lang.Unwrap() before calling this method if this Lang
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Lang) Update() *LangUpdateOne {
	return (&LangClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Lang entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Lang) Unwrap() *Lang {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Lang is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Lang) String() string {
	var builder strings.Builder
	builder.WriteString("Lang(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", lang=")
	builder.WriteString(l.Lang)
	builder.WriteString(", short=")
	builder.WriteString(l.Short)
	builder.WriteString(", logo=")
	builder.WriteString(l.Logo)
	builder.WriteString(", name=")
	builder.WriteString(l.Name)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", l.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", l.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", l.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// Langs is a parsable slice of Lang.
type Langs []*Lang

func (l Langs) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
