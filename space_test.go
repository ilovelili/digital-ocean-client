package doclient

import (
	"testing"
)

// TestUpload test object upload
func TestUpload(t *testing.T) {
	spaceservice := NewSpaceService("xxxxxxx", "yyyyyyy")
	spaceservice.SetRegion("sgp1")
	spaceservice.SetEndPoint("sgp1.digitaloceanspaces.com")
	spaceservice.SetBucket("dongfeng")

	err := spaceservice.Upload("./test.png")
	if err != nil {
		t.Error(err)
	}
}
