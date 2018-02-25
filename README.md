# test2doc

[![wercker status](https://app.wercker.com/status/fd0bda35af1bdd28e08eee90783127c2/s/master "wercker status")](https://app.wercker.com/project/bykey/fd0bda35af1bdd28e08eee90783127c2)

Generate documentation for your REST/HTTP API from your Go unit tests - a simple addition to Go's testing package.

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

![screenshot](http://s17.postimg.org/6mz3ich1b/Screen_Shot_2015_11_06_at_9_38_46_AM.png)


<br>

### Things to note:
1. Go pkg name **becomes** `Group` name
2. Go handler name **becomes** endpoint title
3. Go handler `godoc` string **becomes** endpoint description
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



### Integrating test2doc

Very few additions, and only to your testing code.

#### 1. Add 3 things to your TestMain:

```go

import (
	"github.com/adams-sarah/test2doc/test"
)

var server *test.Server

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
	//    NewServer takes your HTTP handler as an argument
	server, err := test.NewServer(router)
	if err != nil {
		panic(err.Error())
	}

	// .. then run your tests as usual
	// (remember that os.Exit does not respect defers)
	exitCode := m.Run()


	// 3. Finally, you must tell the wrapped server when you are done testing
	//    so that the buffer can be flushed to an apib doc file
	server.Finish()

	// note that os.Exit does not respect defers.
	os.Exit(exitCode)
}

```

#### Router-specific configurations

`gorilla/mux` configurations

```go
	import "github.com/adams-sarah/test2doc/vars"
	// ...
	
	extractor := vars.MakeGorillaMuxExtractor(myGorillaRouter)
	test.RegisterURLVarExtractor(extractor)
```

`julienschmidt/httprouter` configurations

```go
	import "github.com/adams-sarah/test2doc/vars"
	// ...

	extractor := vars.MakeHTTPRouterExtractor(myHTTPRouter)
	test.RegisterURLVarExtractor(extractor)
```

#### 2. Combine the output `.apib` files
`test2doc` will spit out one doc file per package.

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


You will need to add the doc header (below) and combine all of the package doc files after your tests run.

eg.:

```bash
# find all *.apib files (after tests have run, generated files)
files=`find . -type f -name "*.apib"`

# copy template file to new apiary.apib file
cp apib.tmpl apiary.apib

# copy contents of each generated apib file into apiary.apib
# and delete the apib file
for f in ${files[@]}; do
	cat $f >> apiary.apib
	rm $f
done
```

where `apib.tmpl` includes the doc header information. 
Something like:

```
FORMAT: 1A
HOST: https://api.mysite.com

# The API for My Site

My Site is a fancy site. The API is also fancy.

```

