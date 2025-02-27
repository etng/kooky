package kooky_test

import (
	"fmt"

	"github.com/etng/kooky"
	_ "github.com/etng/kooky/browser/all" // This registers all cookiestore finders!
	// _ "github.com/etng/kooky/browser/chrome" // load only the chrome cookiestore finder
)

func ExampleReadCookies_all() {
	// try to find cookie stores in default locations and
	// read the cookies from them.
	// decryption is handled automatically.
	cookies := kooky.ReadCookies()

	for _, cookie := range cookies {
		fmt.Println(cookie)
	}
}

var _ struct{} // ignore this - for correct working of the documentation tool
