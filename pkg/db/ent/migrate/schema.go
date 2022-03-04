// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppLangsColumns holds the columns for the "app_langs" table.
	AppLangsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppLangsTable holds the schema information for the "app_langs" table.
	AppLangsTable = &schema.Table{
		Name:       "app_langs",
		Columns:    AppLangsColumns,
		PrimaryKey: []*schema.Column{AppLangsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "applang_app_id_lang_id",
				Unique:  true,
				Columns: []*schema.Column{AppLangsColumns[1], AppLangsColumns[2]},
			},
		},
	}
	// LangsColumns holds the columns for the "langs" table.
	LangsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "lang", Type: field.TypeString, Unique: true},
		{Name: "logo", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "short", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// LangsTable holds the schema information for the "langs" table.
	LangsTable = &schema.Table{
		Name:       "langs",
		Columns:    LangsColumns,
		PrimaryKey: []*schema.Column{LangsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "message_id", Type: field.TypeString},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString},
		{Name: "batch_get", Type: field.TypeBool},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "message_app_id_lang_id_message_id",
				Unique:  true,
				Columns: []*schema.Column{MessagesColumns[1], MessagesColumns[3], MessagesColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppLangsTable,
		LangsTable,
		MessagesTable,
	}
)

func init() {
}
