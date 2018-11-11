# Digital Ocean Client

Digital Ocean client designed for easy use. It's used by my own [DongFeng](https://github.com/ilovelili/dongfeng-core) project.

## Dependencies

- [Minio-Go](https://github.com/minio/minio-go/) Client SDK for Amazon S3 Compatible Cloud Storage

## Object Upload

example:

```Go
import doclient

func main() {
    spaceservice := NewSpaceService("{{your digitalocean space API key}}", "{{your digitalocean space API secret}}")
    spaceservice.SetRegion("nyc1")
    spaceservice.SetEndPoint("{{nyc1.digitaloceanspaces.com}}")
    spaceservice.SetBucket("{{your digitalocean bucketname}}")

    err := spaceservice.Upload("./test.png")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Contact

<route666@live.cn>