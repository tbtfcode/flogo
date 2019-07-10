package tcpserver

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/project-flogo/core/trigger"
)

func getJSONMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

const testConfig string = `{
  "id": "mytrigger",
  "ref": "github.com/project-flogo/contrib/trigger/tcp",
  "settings": {
	"network": "tcp",
	"host": "127.0.0.1",
	"port": "8982"
  },
"handlers": [
    {
      "settings": {
      },
      "action" {
	     "id": "dummy"
      }
    }
  ]
}`

func TestCreate(t *testing.T) {

	// New factory
	md := trigger.NewMetadata(getJSONMetadata())
	f := NewFactory(md)

	if f == nil {
		t.Fail()
	}

	// New Trigger
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), config)
	trg := f.New(&config)

	if trg == nil {
		t.Fail()
	}
}
