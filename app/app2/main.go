// main.go
package main

import (
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const Version = "1.2.3"
const slug = "Jyury11/go-update"

func main() {
	println("Ba dum, tss!")

	v := semver.MustParse(Version)
	latest, _ := selfupdate.UpdateSelf(v, slug)
	if latest.Version.Equals(v) {
		log.Println("some version")
	}
}
