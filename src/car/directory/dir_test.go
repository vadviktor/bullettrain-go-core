package carDirectory

import (
	"os"
	"path/filepath"
	"testing"
)

const failTpl = "Unexpected directory format.\nEXPECTED: %s\nGOT:      %s"

func setup() {
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_PATH_SEPARATOR", "")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_DEPTH_INDICATOR", "...")
}

func TestFullPath(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FRONT_MAX_LENGTH", "-1")
	directory := filepath.Clean("/usr/share/doc/vpn/easy")
	expected := "usrsharedocvpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestCustomPathSeparator(t *testing.T) {
	setup()
	sep := "|"
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_PATH_SEPARATOR", sep)
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FRONT_MAX_LENGTH", "-1")
	directory := filepath.Clean("/usr/share/doc/vpn/easy")
	expected := "|usr|share|doc|vpn|easy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestMergeMode(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FRONT_MAX_LENGTH", "2")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_TAIL_MAX_LENGTH", "1")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ACRONYM_MODE", "merge")
	directory := filepath.Clean("/usr/share/doc/vpn/easy")
	expected := "usr...easy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestAcronymMode(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FRONT_MAX_LENGTH", "2")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_TAIL_MAX_LENGTH", "1")
	directory := filepath.Clean("/usr/share/doc/vpn/easy")
	expected := "usrs*d*v*easy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestCustomAcronymSymbol(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ELLIPSIS", "~")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FRONT_MAX_LENGTH", "2")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_TAIL_MAX_LENGTH", "1")
	directory := filepath.Clean(filepath.Clean("/usr/share/doc/vpn/easy"))
	expected := "usrs~d~v~easy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestCustomDepthIndicator(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ACRONYM_MODE", "merge")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_DEPTH_INDICATOR", "+++")
	directory := filepath.Clean(filepath.Clean("/usr/share/doc/vpn/easy"))
	expected := "usr+++vpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestHomeDirectoryShorthand(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FRONT_MAX_LENGTH", "-1")
	os.Setenv("HOME", filepath.Clean("/home/ikon"))
	directory := filepath.Clean("/home/ikon/doc/vpn/easy")
	expected := "~docvpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}
