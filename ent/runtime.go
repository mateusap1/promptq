// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/mateusap1/promptq/ent/promptrequest"
	"github.com/mateusap1/promptq/ent/promptresponse"
	"github.com/mateusap1/promptq/ent/schema"
	"github.com/mateusap1/promptq/ent/user"
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
	promptrequest.DefaultIdentifier = promptrequestDescIdentifier.Default.(func() uuid.UUID)
	// promptrequestDescIsQueued is the schema descriptor for is_queued field.
	promptrequestDescIsQueued := promptrequestFields[2].Descriptor()
	// promptrequest.DefaultIsQueued holds the default value on creation for the is_queued field.
	promptrequest.DefaultIsQueued = promptrequestDescIsQueued.Default.(bool)
	// promptrequestDescIsAnswered is the schema descriptor for is_answered field.
	promptrequestDescIsAnswered := promptrequestFields[3].Descriptor()
	// promptrequest.DefaultIsAnswered holds the default value on creation for the is_answered field.
	promptrequest.DefaultIsAnswered = promptrequestDescIsAnswered.Default.(bool)
	// promptrequestDescCreateDate is the schema descriptor for create_date field.
	promptrequestDescCreateDate := promptrequestFields[4].Descriptor()
	// promptrequest.DefaultCreateDate holds the default value on creation for the create_date field.
	promptrequest.DefaultCreateDate = promptrequestDescCreateDate.Default.(func() time.Time)
	promptresponseFields := schema.PromptResponse{}.Fields()
	_ = promptresponseFields
	// promptresponseDescCreateDate is the schema descriptor for create_date field.
	promptresponseDescCreateDate := promptresponseFields[1].Descriptor()
	// promptresponse.DefaultCreateDate holds the default value on creation for the create_date field.
	promptresponse.DefaultCreateDate = promptresponseDescCreateDate.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateDate is the schema descriptor for create_date field.
	userDescCreateDate := userFields[3].Descriptor()
	// user.DefaultCreateDate holds the default value on creation for the create_date field.
	user.DefaultCreateDate = userDescCreateDate.Default.(func() time.Time)
}
