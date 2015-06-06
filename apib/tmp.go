package apib

import (
	"fmt"
	"net/url"
	"strings"
)

var (
	tmpResourceCount = 0
	tmpResourceNames = []string{"Thing", "Foozle", "Bardy", "Wallet"}
	tmpDoc           = &Doc{
		Title:       "JSON Placeholder API",
		Description: "Fake Online REST API for Testing and Prototyping",
		Metadata: &Metadata{
			Format: FORMAT,
			Host:   "http://jsonplaceholder.typicode.com",
		},
		ResourceGroups: []*ResourceGroup{
			&ResourceGroup{
				Title:       "A Lovely Resource Group",
				Description: "All CRUD endpoints for A Lovely Resource Group",
				Resources:   newTmpResources(1),
			},
			&ResourceGroup{
				Title:       "An Awesome Resource Group",
				Description: "All non-CRUD endpoints",
				Resources:   newTmpResources(2),
			},
		},
	}
)

func newTmpResources(n int) []*Resource {
	resources := make([]*Resource, n)

	for i := 0; i < n; i++ {
		resources[i] = newTmpResource()
	}
	return resources
}

func newTmpResource() *Resource {
	tmpResourceCount += 1
	index := tmpResourceCount % len(tmpResourceNames)
	name := tmpResourceNames[index]

	var query string
	if index == 0 {
		query = fmt.Sprintf("?a=0")
	} else if index == 1 {
		query = fmt.Sprintf("?b=hello+world")
	}

	u, err := url.Parse(fmt.Sprintf("/%s/id/%d"+query, strings.ToLower(name), tmpResourceCount))
	if err != nil {
		panic(err.Error())
	}

	return &Resource{
		Title:       fmt.Sprintf("Resource #%d", tmpResourceCount),
		Description: "This resource is just like all the others.",
		URL:         &URL{u, nil},
	}
}
