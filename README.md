# Go DataDog Tagging Struct

Building up a slice of tags every time you want to do a Datadog metric is tedious. 

This library makes passing tags to datadog much easier

## Installation

```
go get -u github.com/webdestroya/ddtags
```

## Usage

```go
package example

import (
  "github.com/webdestroya/ddtags"
	ddlambda "github.com/DataDog/datadog-lambda-go"
)

type DDTags struct {
  Service          string `ddtag:"service_name"`
  AvailabilityZone string `ddtag:"availability-zone"`
  Tier             string `ddtag:"tier"`
}

func Something() {

  ddlambda.Metric("some.metric.name", 123, ddtags.Extract(&DDTags{
    Service: "some_service",
    Tier:    "gold",
  })...)

}
```

All scalar types are supported (as well as their pointer variants). The zero value is ignored. If you want to explicitly tag the zero value of a type, you'll need to make it a pointer