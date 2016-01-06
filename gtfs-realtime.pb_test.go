package transit_realtime

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"testing"
)

func TestDecode(t *testing.T) {
	bytes, err := ioutil.ReadFile("vehicle_position.pb")
	if err != nil {
		t.Errorf("Unable to read file: %s", err.Error())
	}

	feed := &FeedMessage{}
	err = proto.Unmarshal(bytes, feed)
	if err != nil {
		t.Errorf("Unexpected error unmarshalling protobuf: %s", err.Error())
	}

	if len(feed.Entity) != 1 {
		t.Errorf("Expected number of entities to be 1, got: %v", len(feed.Entity))
	}

	entity := feed.Entity[0]

	if *entity.Id != "1" {
		t.Errorf("Expected entity Id to be '1', got: %s", entity.Id)
	}

	if entity.Vehicle == nil {
		t.Errorf("Expected entity to have vehicle, got nil")
	}
}
