// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
)

// 权限
type Permission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID *int `json:"parent_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Type holds the value of the "type" field.
	Type permission.Type `json:"type,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Sort holds the value of the "sort" field.
	Sort int `json:"sort,omitempty"`
	// Attrs holds the value of the "attrs" field.
	Attrs map[string]interface{} `json:"attrs,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionQuery when eager-loading is set.
	Edges        PermissionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PermissionEdges holds the relations/edges for other nodes in the graph.
type PermissionEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Permission `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Permission `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedRoles    map[string][]*Role
	namedChildren map[string][]*Permission
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionEdges) ParentOrErr() (*Permission, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: permission.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) ChildrenOrErr() ([]*Permission, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldAttrs:
			values[i] = new([]byte)
		case permission.FieldID, permission.FieldParentID, permission.FieldSort:
			values[i] = new(sql.NullInt64)
		case permission.FieldName, permission.FieldKey, permission.FieldType, permission.FieldPath, permission.FieldDesc:
			values[i] = new(sql.NullString)
		case permission.FieldCreatedAt, permission.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case permission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pe.CreatedAt = value.Time
			}
		case permission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pe.UpdatedAt = value.Time
			}
		case permission.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				pe.ParentID = new(int)
				*pe.ParentID = int(value.Int64)
			}
		case permission.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case permission.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				pe.Key = value.String
			}
		case permission.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pe.Type = permission.Type(value.String)
			}
		case permission.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				pe.Path = value.String
			}
		case permission.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				pe.Desc = value.String
			}
		case permission.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				pe.Sort = int(value.Int64)
			}
		case permission.FieldAttrs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field attrs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.Attrs); err != nil {
					return fmt.Errorf("unmarshal field attrs: %w", err)
				}
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Permission.
// This includes values selected through modifiers, order, etc.
func (pe *Permission) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the Permission entity.
func (pe *Permission) QueryRoles() *RoleQuery {
	return NewPermissionClient(pe.config).QueryRoles(pe)
}

// QueryParent queries the "parent" edge of the Permission entity.
func (pe *Permission) QueryParent() *PermissionQuery {
	return NewPermissionClient(pe.config).QueryParent(pe)
}

// QueryChildren queries the "children" edge of the Permission entity.
func (pe *Permission) QueryChildren() *PermissionQuery {
	return NewPermissionClient(pe.config).QueryChildren(pe)
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return NewPermissionClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Permission is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pe.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := pe.ParentID; v != nil {
		builder.WriteString("parent_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(pe.Key)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", pe.Type))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(pe.Path)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(pe.Desc)
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", pe.Sort))
	builder.WriteString(", ")
	builder.WriteString("attrs=")
	builder.WriteString(fmt.Sprintf("%v", pe.Attrs))
	builder.WriteByte(')')
	return builder.String()
}

// NamedRoles returns the Roles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Permission) NamedRoles(name string) ([]*Role, error) {
	if pe.Edges.namedRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Permission) appendNamedRoles(name string, edges ...*Role) {
	if pe.Edges.namedRoles == nil {
		pe.Edges.namedRoles = make(map[string][]*Role)
	}
	if len(edges) == 0 {
		pe.Edges.namedRoles[name] = []*Role{}
	} else {
		pe.Edges.namedRoles[name] = append(pe.Edges.namedRoles[name], edges...)
	}
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pe *Permission) NamedChildren(name string) ([]*Permission, error) {
	if pe.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pe.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pe *Permission) appendNamedChildren(name string, edges ...*Permission) {
	if pe.Edges.namedChildren == nil {
		pe.Edges.namedChildren = make(map[string][]*Permission)
	}
	if len(edges) == 0 {
		pe.Edges.namedChildren[name] = []*Permission{}
	} else {
		pe.Edges.namedChildren[name] = append(pe.Edges.namedChildren[name], edges...)
	}
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission
