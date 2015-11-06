# test2doc

Automatically generate API documentation from your Go unit tests: a simple addition to Go's testing pkg

### Diving right in..

Given a handler func:

```go
// GetWidget retrieves a single Widget
func GetWidget(w http.ResponseWriter, req *http.Request) {
	// get widget
	// respond with widget JSON
	// ...
}
```

And a test for this handler func:

```go
func TestGetWidget(t *testing.T) {
	urlPath := fmt.Sprintf("/widgets/%d", 2)

	resp, err := http.Get(server.URL + urlPath)
	// assert all the things...
}
```

Test2doc will generate markdown documentation for this endpoint in the [API Blueprint](https://github.com/apiaryio/api-blueprint/blob/master/API%20Blueprint%20Specification.md) format, like so:

```
# Group widgets

## /widgets/{id}

+ Parameters
    + id: `2` (number)

### Get Widget [GET]
retrieves a single Widget

+ Response 200 

    + Body

            {
                "Id": 2,
                "Name": "Pencil",
                "Role": "Utensil"
            }        
```

Which you can then parse and host w/ *Apiary.io*, eg [here](http://docs.testingit.apiary.io/#). 
Or use a custom parser and host yourself.

![screenshot](https://www.dropbox.com/s/iyzr3lds7sguv2h/Screen%20Shot%202015-11-06%20at%209.38.46%20AM.png?dl=0)


### How can I get this working?

Very few additions, and only to your testing code.

```go

import (
	"github.com/adams-sarah/test2doc/test"
)

func TestMain(m *testing.M) {
	// 1. Tell test2doc how to get URL vars out of your HTTP requests
	//
	//    The 'URLVarExtractor' function must have the following signature:
	//      func(req *http.Request) map[string]string
	//      where the returned map is of the form map[key]value
	test.RegisterURLVarExtractor(myURLVarExtractorFn)

	// 2. You must use test2doc/test's wrapped httptest.Server instead of
	//    the raw httptest.Server, so that test2doc can listen to and
	//    record requests & responses.
	//
	//    NewServer takes 2 arguments:
	//    - your HTTP handler
	//    - the path to where you would like the apib doc file
	server, err := test.NewServer(router, ".")
	if err != nil {
		panic(err.Error())
	}


	// .. then run your tests as usual
	// (remember that os.Exit does not respect defers,
	//	 so you'll )
	exitCode := m.Run()


	// 3. Finally, you must tell the wrapped server when you are done testing
	//    so that the buffer can be flushed to an apib doc file
	server.Finish()

	// note that os.Exit does not respect defers.
	os.Exit(exitCode)
}

```


`gorilla/mux` configurations

```go
	// NOTE: if you are using gorilla/mux, you must set the router's 
	//  'KeepContext' to true, so that url parameters can be accessed
	//  after the request has been handled.
	router := MyNewRouterFn()
	router.KeepContext = true
```

<br>

### Things to note:
1. Go pkg name **=>** `Group` name
2. Go handler name **=>** endpoint title
3. Go handler `godoc` **=>** endpoint description
4. **Everything else is recorded & interpreted directly from the requests and responses**

Eg.
```go
package widget

// GetWidget retrieves a single Widget
func GetWidget(w http.ResponseWriter, req *http.Request)
```

becomes

```
# Group widget

### Get Widget [*]
retrieves a single Widget
```


<br>

### Installation

`go get github.com/adams-sarah/test2doc/...`

<br>


## OUTSTANDING TODOS:
#### 1. Concat package `apib` docs
The biggest outstanding problem with `test2doc` is that after each package generates its own `apib` file, there is no tooling to concatenate the files into your *one* API Blueprint file.

Eg. A package tree like:

```
.
├── foos
│   ├── foos.go
│   └── foos_test.go
└── widgets
    ├── widgets.go
    └── widgets_test.go
```

Will produce separate apib files, eg:

```
.
├── foos
│   ├── ...
│   └── foos.apib
└── widgets
    ├── ...
    └── widgets.apib
```

But we want something like:

```
.
├── apidoc.apib
├── foos
│   └── ...
└── widgets
    └── ...
```

Where `apidoc.apib` includes the contents of both `foos.apib` and `widgets.apib`.

There are, of course, many ways to do this.

But I'm hoping for **your** suggestions on the best way to integrate this into a testing workflow.


#### 2. Query string params
This one will be a simple addition. Just a matter of time.



