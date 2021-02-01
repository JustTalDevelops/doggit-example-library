package doggit_example

import "github.com/JustTalDevelops/doggit"

// BuildLibrary is called from Doggit.
func BuildLibrary(_ *doggit.Service) *doggit.Library {
	library := doggit.NewLibrary("Example Library", "An example of the Doggit API", "JustTal")

	library.AddHandler(DoggitHandler{})

	return library
}
