// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/mateusap1/promptq/ent/promptrequest"
	"github.com/mateusap1/promptq/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	promptrequestFields := schema.PromptRequest{}.Fields()
	_ = promptrequestFields
	// promptrequestDescIdentifier is the schema descriptor for identifier field.
	promptrequestDescIdentifier := promptrequestFields[0].Descriptor()
	// promptrequest.DefaultIdentifier holds the default value on creation for the identifier field.
	promptrequest.DefaultIdentifier = promptrequestDescIdentifier.Default.(func() string)
	// promptrequestDescIsQueued is the schema descriptor for is_queued field.
	promptrequestDescIsQueued := promptrequestFields[2].Descriptor()
	// promptrequest.DefaultIsQueued holds the default value on creation for the is_queued field.
	promptrequest.DefaultIsQueued = promptrequestDescIsQueued.Default.(bool)
}
