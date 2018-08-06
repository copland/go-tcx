package tcx

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type TrainingCenterDatabase struct {
	XMLName xml.Name `xml:"TrainingCenterDatabase"`
	Activities Activities `xml:"Activities"`
}

type Activities struct {
   XMLName xml.Name `xml:"Activities"`
   Activities []Activity `xml:"Activity"`
}

type Activity struct {
  XMLName xml.Name `xml:"Activity"`
  Sport   string `xml:"Sport,attr"`
  Id      string `xml:"Id"`
  Laps    []Lap  `xml:"Lap"`
  Creator Creator `xml:"Creator"`
}

type Creator struct {
	XMLName xml.Name `xml:"Creator"`
	Name string `xml:"Name"`
	UnitId int `xml:"UnitId"`
	ProductID int `xml:"ProductID"`
}

type Lap struct {
  XMLName xml.Name `xml:"Lap"`
  StartTime string `xml:"StartTime,attr"`
  TotalTimeSeconds float64 `xml:"TotalTimeSeconds"`
  DistanceMeters float64 `xml:"DistanceMeters"`
  Calories int16 `xml:"Calories"`
  Intensity string `xml:"Intensity"`
  TriggerMethod string `xml:"TriggerMethod"`
  Track Track `xml: "Track"`

}

type Track struct {
	XMLName xml.Name `xml:"Track"`
	Trackpoints []Trackpoint `xml:"Trackpoint"`
}

type Trackpoint struct  {
	XMLName xml.Name `xml:"Trackpoint"`
	Time string `xml:"Time"`
	Position Position `xml:"Position"`
	AltitudeMeters float32 `xml:"AltitudeMeters"`
	DistanceMeters float64 `xml:"DistanceMeters"`
        HeartRateBpm HeartRateBpm `xml:"HeartRateBpm"`
}

type Position struct {
	XMLName xml.Name `xml:"Position"`
	LatitudeDegrees float64 `xml:"LatitudeDegrees"`
	LongitudeDegrees float64 `xml:"LongitudeDegrees"`
}

type HeartRateBpm struct {
	XMLName xml.Name `xml:"HeartRateBpm"`
	Value int `xml:"Value"`
}

func ReadTCX(reader io.Reader) (Activities, error) {
	var database TrainingCenterDatabase
	if err := xml.NewDecoder(reader).Decode(&database); err != nil {
		return Activities{}, err
	}
	return database.Activities, nil
}

func main() {

	tcxPath, err := filepath.Abs("15612758058.tcx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tcxFile, err := os.Open(tcxPath)
        defer tcxFile.Close()
	if err != nil {
	    log.Fatal(err)
	}

	activities, err := ReadTCX(tcxFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
        for _, activity := range activities.Activities {
		fmt.Printf("Sport: %s ID: %s\n", activity.Sport, activity.Id)
		for _, lap := range activity.Laps {
			fmt.Printf("  Total Time: %f Total Distance: %f\n", lap.TotalTimeSeconds, lap.DistanceMeters)
			for _, trackpoint := range lap.Track.Trackpoints {
				fmt.Printf("   Time: %s Position: (%f, %f)\n", trackpoint.Time, trackpoint.Position.LatitudeDegrees, trackpoint.Position.LongitudeDegrees)
			}
		}
	}

}


