package carDirectory

import (
	"os"
	"testing"
)

const failTpl = "Unexpected directory format.\nEXPECTED: %s\nGOT:      %s"

func setup() {
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_PATH_SEPARATOR", "")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "3")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_DEPTH_INDICATOR", "...")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "true")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "true")
}

func TestFullPath(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "0")
	directory := "/usr/share/doc/vpn/easy"
	expected := "usrsharedocvpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestWithoutRoot(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "0")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	directory := "/usr/share/doc/vpn/easy"
	expected := "usrsharedocvpneasy"

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
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "5")
	directory := "/usr/share/doc/vpn/easy"
	expected := "|usr|share|doc|vpn|easy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestFirstDirShowingLength2(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "2")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "true")
	directory := "/usr/share/doc/vpn/easy"
	expected := "usr...easy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestFirstDirNotShowingLength2(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "2")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "false")
	directory := "/usr/share/doc/vpn/easy"
	expected := "...vpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestFirstDirNotShowingFullLength(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "0")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "false")
	directory := "/usr/share/doc/vpn/easy"
	expected := "usrsharedocvpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestFirstDirNotShowingLength1(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "1")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "false")
	directory := "/usr/share/doc/vpn/easy"
	expected := "...easy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestFirstDirShowingFullLength(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "0")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "true")
	directory := "/usr/share/doc/vpn/easy"
	expected := "usrsharedocvpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestFirstDirShowingLength1(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "1")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "true")
	directory := "/usr/share/doc/vpn/easy"
	expected := "usr...easy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestFirstDirShowingLength3(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "3")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "true")
	directory := "/usr/share/doc/vpn/easy"
	expected := "usr...vpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestFirstDirNotShowingLength3(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "3")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "false")
	directory := "/usr/share/doc/vpn/easy"
	expected := "...docvpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestCustomDepthIndicator(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "3")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_DEPTH_INDICATOR", "+++")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_ROOT_SHOW", "false")
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_FIRST_DIR_SHOW", "true")
	directory := "/usr/share/doc/vpn/easy"
	expected := "usr+++vpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}

func TestHomeDirectoryShorthand(t *testing.T) {
	setup()
	os.Setenv("BULLETTRAIN_CAR_DIRECTORY_MAX_LENGHT", "5")
	os.Setenv("HOME", "/home/ikon")
	directory := "/home/ikon/doc/vpn/easy"
	expected := "~docvpneasy"

	actual := rebuildDirForRender(directory)

	if actual != expected {
		t.Logf(failTpl, expected, actual)
		t.Fail()
	}
}
