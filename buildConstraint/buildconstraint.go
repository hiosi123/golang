//go:build ignore
// +build ignore

package buildconstraint

// A build constraint or a build tag, is an identifier added to a piece of code that determines
// when the file should be inclueded in a package during the build process.

// main_dev.go
var explain1 = "+build !prod"

// main_prod.go
var explain2 = "+build prod"

// main.go

// go build -tags prod

//Usage
// test set you run, CI/CD pipeline, intergration docker
