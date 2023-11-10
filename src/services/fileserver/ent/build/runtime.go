// Code generated by ent, DO NOT EDIT.

package build

import (
	"time"

	"github.com/varsotech/varsoapi/src/services/fileserver/ent/build/checkpoint"
	"github.com/varsotech/varsoapi/src/services/fileserver/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	checkpointFields := schema.Checkpoint{}.Fields()
	_ = checkpointFields
	// checkpointDescCreateTime is the schema descriptor for create_time field.
	checkpointDescCreateTime := checkpointFields[0].Descriptor()
	// checkpoint.DefaultCreateTime holds the default value on creation for the create_time field.
	checkpoint.DefaultCreateTime = checkpointDescCreateTime.Default.(func() time.Time)
}
