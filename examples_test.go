package ddtags_test

import (
	"fmt"
	"strings"

	"github.com/webdestroya/ddtags"
)

type SimpleTags struct {
	Service   string `ddtag:"service_name"`
	AvailZone string `ddtag:"availability-zone"`
	Tier      string `ddtag:"tier"`
}

func ExampleExtract() {
	tags := &SimpleTags{
		Service: "dummy-service",
		Tier:    "gold",
	}

	fmt.Println(strings.Join(ddtags.Extract(tags), ","))
	// Output:
	// service_name:dummy-service,tier:gold
}
