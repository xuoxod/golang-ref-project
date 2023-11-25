package collections

import (
	"testing"

	"github.com/xuoxod/lab/internal/utils"
)

func TestUser(t *testing.T) {
	var uid string = utils.GenerateID()
	var fname string = utils.GenerateName(7)
	var lname string = utils.GenerateName(13)

	var uidPointer *string = &uid
	var fnamePointer *string = &fname
	var lnamePointer *string = &lname

	var user Details = User{
		UID:       *uidPointer,
		FirstName: *fnamePointer,
		LastName:  *lnamePointer,
	}

	got := user.UserInfo()["uid"]
	want := uid

	if got != want {
		t.Fatalf("%T: %s is not equal to %T: %s\n", got, got, want, want)
	}

	got = user.UserInfo()["fname"]
	want = fname

	if got != want {
		t.Fatalf("%T: %s is not equal to %T: %s\n", got, got, want, want)
	}

	got = user.UserInfo()["lname"]
	want = lname

	if got != want {
		t.Fatalf("%T: %s is not equal to %T: %s\n", got, got, want, want)
	}

	user.Info()

}
