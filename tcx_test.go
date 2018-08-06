package tcx_test

import (
	"os"
	"path/filepath"
	"testing"
	"github.com/copland/go-tcx"
)

func loadTestFile(file_path string) (tcx.Activities, error) {

	tcxPath, err := filepath.Abs(file_path)
	if err != nil {
		return tcx.Activities{}, err
	}
	tcxFile, err := os.Open(tcxPath)
        defer tcxFile.Close()
	if err != nil {
		return tcx.Activities{}, err
	}

	activities, err := tcx.ReadTCX(tcxFile)
	if err != nil {
		return tcx.Activities{}, err
	}
	return activities, nil
}

func TestReadTCXMalformedFile(t *testing.T) {
	_, err := loadTestFile("test_resources/bad_format.tcx")
	if err == nil {
		t.Errorf("should have failed with error on bad format")
	}
}

func TestReadTCXNoActivities(t *testing.T) {
	activities, err := loadTestFile("test_resources/no_activities.tcx")
	if err != nil {
		t.Errorf("error loading file: test_resources/no_activities.tcx")
	}
	if len(activities.Activities) > 0 {
		t.Errorf("should have loaded 0 activities")
	}
}

func TestReadTCXCorrectActivityCount(t *testing.T) {
	activities, err := loadTestFile("test_resources/valid_activities.tcx")
	if err != nil {
		t.Errorf("error loading file: test_resources/valid_activities.tcx")
	}
	if len(activities.Activities) != 2 {
		t.Errorf("should have loaded exactly 2 activities")
	}
}

func TestReadTCXCorrectSport(t *testing.T) {
	activities, err := loadTestFile("test_resources/valid_activities.tcx")
	if err != nil {
		t.Errorf("error loading file: test_resources/valid_activities.tcx")
	}
	expected_sports := []string{"Running", "Other"}
	for index, activity := range activities.Activities {
		if expected_sports[index] != activity.Sport {
		    t.Errorf("did not load correct sport for Activity[%d]", index)
		}
	}
}

func TestReadTCXCorrectTrackpointCount(t *testing.T) {
	activities, err := loadTestFile("test_resources/valid_activities.tcx")
	if err != nil {
		t.Errorf("error loading file: test_resources/valid_activities.tcx")
	}

	expected := 4
	actual := 0
	for _, lap := range activities.Activities[0].Laps {
             actual += len(lap.Track.Trackpoints) 
	}
	if actual != expected {
		t.Errorf("expected %d trackpoints, got %d", expected, actual)
	}
}
