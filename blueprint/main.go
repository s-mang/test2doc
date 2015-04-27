// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"net/url"
	"os"
)

func main() {
	reqURI, err := url.Parse("/users/{id}")
	if err != nil {
		panic("url.Parse: " + err.Error())
	}

	jsonBody := "application/json"
	body := bytes.NewBufferString(`"lastName": "Gopher"`)

	apib := &APIBlueprint{
		Metadata:    NewMetadata("http://www.weather-foo.com"),
		Name:        "Weather Foo API",
		Description: "An API for all your weather foo.",
		ResourceGroups: []*ResourceGroup{
			&ResourceGroup{
				Name:        "Users",
				Description: "Account-holders on weather-foo.com",
				Resources: []*Resource{
					&Resource{
						Name:        "User",
						Description: "A specific weather-foo.com account-holder",
						Parameters: []*Parameter{
							&Parameter{
								Name:         "id",
								ExampleValue: "1",
								Type:         Number,
								IsRequired:   true,
								Description:  "user account identifier",
							},
						},
						Actions: []*Action{
							&Action{
								HTTPMethod: "PUT",
								URI:        reqURI,
								Requests: []*Request{
									&Request{
										httpIO{
											Header: map[string][]string{
												"Content-Type": []string{jsonBody},
											},
											ContentType: jsonBody,
											Body:        body,
										},
										[]*Response{
											&Response{
												200,
												httpIO{
													Header: map[string][]string{
														"Content-Type": []string{jsonBody},
														"X-Foo":        []string{"bar"},
													},
													ContentType: jsonBody,
													Body:        body,
												},
											},
											&Response{
												404,
												httpIO{
													Name:        "User not found",
													Description: "No user was found with the given id.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	b, err := json.MarshalIndent(apib, "", "  ")
	if err != nil {
		panic("json.MarshalIndent: " + err.Error())
	}

	os.Stdout.Write(b)

}
