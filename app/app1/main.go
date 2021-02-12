// main.go
package main

import (
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const Version = "4.2.3"
const slug = "Jyury11/go-update"

func main() {
	selfupdate.EnableLog()

	println("Hello World!")

	v := semver.MustParse(Version)
	latest, err := selfupdate.UpdateSelf(v, slug)
	log.Println(latest)
	log.Println(latest.Version)
	log.Println(latest.Name)
	log.Println(latest.RepoName)
	log.Println(latest.RepoOwner)
	log.Println(latest.Version.Value())
	log.Println(err)
	if latest.Version.Equals(v) {
		log.Println("some version")
	} else {
		log.Println("update version")
	}
}
