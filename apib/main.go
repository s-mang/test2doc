// +build ignore

package main

import (
	"bytes"
	"fmt"
	"net/url"
)

func main() {
	reqURI, err := url.Parse("/users/{id}")
	if err != nil {
		panic("url.Parse: " + err.Error())
	}

	jsonBody := "application/json"
	body := bytes.NewBufferString(`"lastName": "Gopher"`)

	doc := &APIBlueprint{
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
								Name:        "id",
								Value:       "1",
								Type:        Number,
								IsRequired:  true,
								Description: "user account identifier",
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

	var buf bytes.Buffer
	err = docTemplate.Execute(&buf, doc)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(buf.String())
}
