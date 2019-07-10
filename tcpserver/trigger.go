package tcpserver

import (
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

/*
var triggerMetadata = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}
*/

// Trigger is a stub for your Trigger implementation
type Trigger struct {
	logger   log.Logger
	handlers []trigger.Handler
	//	settings *Settings
	metadata *trigger.Metadata
	config   *trigger.Config
}

// Factory My Trigger factory
type Factory struct {
}

// Metadata implements trigger.Trigger.Metadata
func (t *Trigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// New Creates a new trigger instance for a given id
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	return &Trigger{config: config}, nil
}

// Initialize implements trigger.Init.Initialize
func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()
	t.handlers = ctx.GetHandlers()

	t.logger.Infof("Initialize implements trigger.Init.Initialize : %s", "")

	return nil
}

// Start implements trigger.Trigger.Start
func (t *Trigger) Start() error {
	// start the trigger
	t.logger.Infof("start the trigger : %s", "")
	return nil
}

// Stop implements trigger.Trigger.Start
func (t *Trigger) Stop() error {
	// stop the trigger
	t.logger.Infof("stop the trigger : %s", "")
	return nil
}
