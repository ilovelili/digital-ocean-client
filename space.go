package doclient

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/minio/minio-go"
	"golang.org/x/net/context"
)

// UploadResponse upload response
type UploadResponse struct {
	// Location location of uploaded file
	Location string
	Error    error
}

// SpaceContext context includes EndPoint, Region and FileName info
type SpaceContext struct {
	EndPoint string
	Region   string
	Bucket   string
}

// SpaceService service defines context
type SpaceService struct {
	Context      *SpaceContext
	AccessKey    string
	AccessSecret string
}

// NewSpaceService service initializer
func NewSpaceService(key, secret string) *SpaceService {
	return &SpaceService{
		Context:      new(SpaceContext),
		AccessKey:    key,
		AccessSecret: secret,
	}
}

// SetEndPoint set endpoint
func (ss *SpaceService) SetEndPoint(endpoint string) {
	ss.Context.check()
	ss.Context.EndPoint = endpoint
}

// GetEndPoint get endpoint
func (ss *SpaceService) GetEndPoint() string {
	return ss.Context.EndPoint
}

// SetBucket set bucket
func (ss *SpaceService) SetBucket(bucket string) {
	ss.Context.check()
	ss.Context.Bucket = bucket
}

// GetBucket get bucket
func (ss *SpaceService) GetBucket() string {
	return ss.Context.Bucket
}

// SetRegion set region
func (ss *SpaceService) SetRegion(region string) {
	ss.Context.check()
	ss.Context.Region = region
}

// GetRegion get region
func (ss *SpaceService) GetRegion() string {
	return ss.Context.Region
}

func (ss *SpaceService) client() (*minio.Client, error) {
	ss.Context.check()
	if ss.Context.EndPoint == "" {
		return nil, fmt.Errorf("invalid endpoint")
	}

	// Initiate a client using DigitalOcean Spaces.
	return minio.New(ss.Context.EndPoint, ss.AccessKey, ss.AccessSecret, true)
}

// Upload upload file
func (ss *SpaceService) Upload(filepath string, timeout ...int) (resp *UploadResponse) {
	resp = new(UploadResponse)

	client, err := ss.client()
	if err != nil {
		resp.Error = err
		return
	}

	if exsit, err := client.BucketExists(ss.GetBucket()); !exsit || err != nil {
		resp.Error = fmt.Errorf("bucket does not exist")
		return
	}

	t := 180 * time.Second
	if len(timeout) > 0 {
		t = time.Duration(timeout[0]) * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	objname := resolveObjName(filepath)
	contenttype, err := resolveContentType(filepath)
	if err != nil {
		resp.Error = err
		return
	}

	_, err = client.FPutObjectWithContext(ctx, ss.GetBucket(), objname, filepath, minio.PutObjectOptions{
		ContentType: contenttype,
	})

	if err != nil {
		resp.Error = err
	} else {
		resp.Location = fmt.Sprintf("https://%s/%s/%s", ss.Context.EndPoint, ss.Context.Bucket, objname)
	}

	return
}

// AsyncUpload async upload
func (ss *SpaceService) AsyncUpload(filepath string) (errchan chan<- *UploadResponse) {
	errchan = make(chan *UploadResponse)
	go func() {
		errchan <- ss.Upload(filepath)
	}()
	return errchan
}

func resolveObjName(fullfilename string) string {
	return filepath.Base(fullfilename)
}

func resolveContentType(fullfilename string) (string, error) {
	f, err := os.Open(fullfilename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func (sc *SpaceContext) check() {
	if sc == nil {
		log.Fatal("invalid context")
	}
}
