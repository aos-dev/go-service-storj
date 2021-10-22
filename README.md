# go-service-storj

[Storj DCS] (Decentralized Cloud Storage) support for [go-storage].

[Storj DCS]: https://www.storj.io
[go-storage]: https://github.com/beyondstorage/go-storage

## Notes

**This package has been moved to [go-storage](https://github.com/beyondstorage/go-storage/tree/master/services/storj).**

```shell
go get go.beyondstorage.io/services/storj
```

## Install

```go
go get github.com/beyondstorage/go-service-storj
```

## Usage

```go
import (
	"log"

	_ "github.com/beyondstorage/go-service-storj"
	"github.com/beyondstorage/go-storage/v4/services"
)

func main() {
	store, err := services.NewStoragerFromString("storj://bucket_name/path/to/workdir")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/beyondstorage/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/storj) about go-service-storj.
