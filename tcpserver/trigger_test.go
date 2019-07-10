package tcpserver

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

func getJSONMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

const testConfig string = `{
  "id": "tcpserver",
  "ref": "github.com/tbtfcode/flogo/tcpserver",
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
	// md := trigger.NewMetadata(getJSONMetadata())

	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig), config)
	// assert.Nil(t, err)

	// actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
	// 	//do nothing
	// })}

	// trg, err := test.InitTrigger(f, config, actions)
	trg, err := f.New(config)
	assert.Nil(t, err)
	assert.NotNil(t, trg)

	err = trg.Start()
	assert.Nil(t, err)
	err = trg.Stop()
	assert.Nil(t, err)
}
