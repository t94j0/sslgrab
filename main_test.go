package sslgrab

import "fmt"

func ExampleGrab() {
	certs, _ := Grab("google.com:443")
	fmt.Println(certs[0].Subject.CommonName)
	// output:
	// *.google.com
}
