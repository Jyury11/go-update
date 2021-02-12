// main.go
package main

const Version = "1.2.3"
const slug = "Jyury11/go-update"

func main() {
	println("Ba dum, tss!")

	v := semver.MustParse(Version)
	latest, err := selfupdate.UpdateSelf(v, slug)
	if !latest.Version.Equals(v) {
		errors.Wrap(err, "Binary update failed")
	}
}
