// main.go
package main

import (
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const Version = "2.2.3"
const slug = "Jyury11/go-update"

func main() {
	println("Hello World!")

	v := semver.MustParse(Version)
	latest, err := selfupdate.UpdateSelf(v, slug)
	if latest.Version.Equals(v) {
		log.Println(latest)
		log.Println(err)
		log.Println("some version")
	} else {
		log.Println(latest)
		log.Println(err)
		log.Println("update version")
	}
}
