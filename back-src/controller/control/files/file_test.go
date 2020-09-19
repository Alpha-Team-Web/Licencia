package files

import (
	"back-src/model/sql"
	"testing"
)

//execute this after uploading
func TestUpload(t *testing.T) {
	db := sql.NewDb()
	if err := db.Initialize(); err != nil {
		t.Error(err)
	}
	profile, err := db.ProfileTable.GetProfileImage("freelancer", "ashkan")
	if err != nil {
		t.Error(err)
	}
	if string(profile.Data) != "aaa" {
		t.Errorf("Error : %s", string(profile.Data))
	}
}
