# Digital Ocean Client

Digital Ocean client designed for easy use. It's used by my own [DongFeng](https://github.com/ilovelili/dongfeng-core) project.

## Dependencies

- [Minio-Go](https://github.com/minio/minio-go/) Client SDK for Amazon S3 Compatible Cloud Storage

## Object Upload

example:

```Go
import doclient

func main() {
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
        log.Fatal(err)
    }
}
```

## Contact

<route666@live.cn>