// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/message"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreateAt is the schema descriptor for create_at field.
	messageDescCreateAt := messageFields[5].Descriptor()
	// message.DefaultCreateAt holds the default value on creation for the create_at field.
	message.DefaultCreateAt = messageDescCreateAt.Default.(func() uint32)
	// messageDescUpdateAt is the schema descriptor for update_at field.
	messageDescUpdateAt := messageFields[6].Descriptor()
	// message.DefaultUpdateAt holds the default value on creation for the update_at field.
	message.DefaultUpdateAt = messageDescUpdateAt.Default.(func() uint32)
	// message.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	message.UpdateDefaultUpdateAt = messageDescUpdateAt.UpdateDefault.(func() uint32)
	// messageDescDeleteAt is the schema descriptor for delete_at field.
	messageDescDeleteAt := messageFields[7].Descriptor()
	// message.DefaultDeleteAt holds the default value on creation for the delete_at field.
	message.DefaultDeleteAt = messageDescDeleteAt.Default.(func() uint32)
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageFields[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() uuid.UUID)
}
