// Code generated by assertions generator; DO NOT EDIT.

package resourceshowoutputassert

import (
	"testing"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

// to ensure sdk package is used
var _ = sdk.Object{}

type StreamShowOutputAssert struct {
	*assert.ResourceAssert
}

func StreamShowOutput(t *testing.T, name string) *StreamShowOutputAssert {
	t.Helper()

	s := StreamShowOutputAssert{
		ResourceAssert: assert.NewResourceAssert(name, "show_output"),
	}
	s.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &s
}

func ImportedStreamShowOutput(t *testing.T, id string) *StreamShowOutputAssert {
	t.Helper()

	s := StreamShowOutputAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "show_output"),
	}
	s.AddAssertion(assert.ValueSet("show_output.#", "1"))
	return &s
}

////////////////////////////
// Attribute value checks //
////////////////////////////

func (s *StreamShowOutputAssert) HasCreatedOn(expected time.Time) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("created_on", expected.String()))
	return s
}

func (s *StreamShowOutputAssert) HasName(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("name", expected))
	return s
}

func (s *StreamShowOutputAssert) HasDatabaseName(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("database_name", expected))
	return s
}

func (s *StreamShowOutputAssert) HasSchemaName(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("schema_name", expected))
	return s
}

func (s *StreamShowOutputAssert) HasTableOn(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("table_on", expected))
	return s
}

func (s *StreamShowOutputAssert) HasOwner(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("owner", expected))
	return s
}

func (s *StreamShowOutputAssert) HasComment(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("comment", expected))
	return s
}

func (s *StreamShowOutputAssert) HasTableName(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("table_name", expected))
	return s
}

func (s *StreamShowOutputAssert) HasSourceType(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("source_type", expected))
	return s
}

func (s *StreamShowOutputAssert) HasBaseTables(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("base_tables", expected))
	return s
}

func (s *StreamShowOutputAssert) HasType(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("type", expected))
	return s
}

func (s *StreamShowOutputAssert) HasStale(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("stale", expected))
	return s
}

func (s *StreamShowOutputAssert) HasMode(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("mode", expected))
	return s
}

func (s *StreamShowOutputAssert) HasStaleAfter(expected time.Time) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("stale_after", expected.String()))
	return s
}

func (s *StreamShowOutputAssert) HasInvalidReason(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("invalid_reason", expected))
	return s
}

func (s *StreamShowOutputAssert) HasOwnerRoleType(expected string) *StreamShowOutputAssert {
	s.AddAssertion(assert.ResourceShowOutputValueSet("owner_role_type", expected))
	return s
}
