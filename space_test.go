package doclient

import (
	"testing"
)

// TestUpload test object upload
func TestUpload(t *testing.T) {
	spaceservice := NewSpaceService("<<api key>>", "<<secret key>>")
	spaceservice.SetRegion("sgp1")
	spaceservice.SetEndPoint("sgp1.digitaloceanspaces.com")
	spaceservice.SetBucket("dongfeng")

	opts := &UploadOptions{
		FileName: "./test.png",
		Public:   true,
	}

	resp := spaceservice.Upload(opts)
	if resp.Error != nil {
		t.Error(resp.Error)
	}
}
