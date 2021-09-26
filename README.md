[![Build Status](https://github.com/beyondstorage/go-service-storj/workflows/Unit%20Test/badge.svg?branch=master)](https://github.com/beyondstorage/go-service-storj/actions?query=workflow%3A%22Unit+Test%22)
[![Integration Tests](https://teamcity.beyondstorage.io/app/rest/builds/buildType:(id:Services_Storj_IntegrationTest)/statusIcon)](https://teamcity.beyondstorage.io/buildConfiguration/Services_Storj_IntegrationTest)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/Xuanwo/storage/blob/master/LICENSE)
[![](https://img.shields.io/matrix/beyondstorage@go-storage:matrix.org.svg?logo=matrix)](https://matrix.to/#/#beyondstorage@go-storage:matrix.org)

# go-service-storj

[Storj DCS] (Decentralized Cloud Storage) support for [go-storage].

[Storj DCS]: https://www.storj.io
[go-storage]: https://github.com/beyondstorage/go-storage

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