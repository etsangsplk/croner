//************************************************************************//
// API "croner": Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/rightscale/croner/design
// --out=$(GOPATH)/src/github.com/rightscale/croner
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// JobHref returns the resource href.
func JobHref() string {
	return fmt.Sprintf("/job")
}
