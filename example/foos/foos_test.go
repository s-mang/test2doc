package foos

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetFoos(t *testing.T) {
	urlPath, err := router.Get("GetFoos").URL()
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	resp, err := http.Get(server.URL + urlPath.String() + "?n=10")
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 'resp.StatusCode' (%v) to equal 'http.StatusOK' (%v)", resp.StatusCode, http.StatusOK)
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	fs := map[string]Foo{}
	err = decoder.Decode(&fs)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if len(fs) != len(AllFoos) {
		t.Fatalf("expected 'len(fs)' (%v) to equal 'len(AllFoos)' (%v)", len(fs), len(AllFoos))
	}

	for k, foo := range AllFoos {
		if fs[k].B != foo.B {
			t.Fatalf("expected 'fs[k].B' (%v) to equal 'foo.B' (%v)", fs[k].B, foo.B)
		}
		if fs[k].A != foo.A {
			t.Fatalf("expected 'fs[k].A' (%v) to equal 'foo.A' (%v)", fs[k].A, foo.A)
		}
		if fs[k].R != foo.R {
			t.Fatalf("expected 'fs[k].R' (%v) to equal 'foo.R' (%v)", fs[k].R, foo.R)
		}
	}
}

func TestGetFoo(t *testing.T) {
	key := "ABeeSee"
	urlPath, err := router.Get("GetFoo").URL("key", key)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	resp, err := http.Get(server.URL + urlPath.String())
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 'resp.StatusCode' (%v) to equal 'http.StatusOK' (%v)", resp.StatusCode, http.StatusOK)
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var foo Foo
	err = decoder.Decode(&foo)
	if err != nil {
		t.Fatalf("expected 'err' (%v) be nil", err)
	}

	if foo.B != AllFoos[key].B {
		t.Fatalf("expected 'foo.B' (%v) to equal 'AllFoos[key].B' (%v)", foo.B, AllFoos[key].B)
	}
	if foo.A != AllFoos[key].A {
		t.Fatalf("expected 'foo.A' (%v) to equal 'AllFoos[key].A' (%v)", foo.A, AllFoos[key].A)
	}
	if foo.R != AllFoos[key].R {
		t.Fatalf("expected 'foo.R' (%v) to equal 'AllFoos[key].R' (%v)", foo.R, AllFoos[key].R)
	}
}
