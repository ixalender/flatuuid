package flatuuid_test

import (
	"strings"
	"testing"

	"github.com/ixalender/flatuuid"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Compact(t *testing.T) {
	testUUID := uuid.NewString()

	flatUUID, err := flatuuid.Compact(testUUID)

	assert.Nil(t, err)
	expectedUUID := strings.ReplaceAll(testUUID, "-", "")
	assert.Equal(t, expectedUUID, flatUUID)

	_, err = flatuuid.Compact("")
	assert.NotNil(t, err)
	assert.True(t, len(err.Error()) > 0)

	_, err = flatuuid.Compact("12345")
	assert.NotNil(t, err)

	rfc4122UUID := "urn:uuid:f47ac10b-58cc-4372-0567-0e02b2c3d479"
	flatUUID, err = flatuuid.Compact(rfc4122UUID)
	expectedUUID = strings.Replace(strings.ReplaceAll(rfc4122UUID, "-", ""), "urn:uuid:", "", 1)
	assert.Nil(t, err)
	assert.Equal(t, expectedUUID, flatUUID)

	microsoftGUIDUpper := "{78B83322-84F6-46E8-A0BE-238FB5703A09}"
	flatUUID, err = flatuuid.Compact(microsoftGUIDUpper)
	expectedUUIDLower := strings.ToLower(strings.Trim(strings.ReplaceAll(microsoftGUIDUpper, "-", ""), "{}"))
	assert.Nil(t, err)
	assert.Equal(t, expectedUUIDLower, flatUUID)
}
