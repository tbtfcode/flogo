package tcpserver

import (
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

// TCPTrigger is a stub for your Trigger implementation
type TCPTrigger struct {
	logger   log.Logger
	handlers []trigger.Handler
	//	settings *Settings
	metadata *trigger.Metadata
	config   *trigger.Config
}

// TCPTriggerFactory My Trigger factory
type TCPTriggerFactory struct {
	metadata *trigger.Metadata
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &TCPTriggerFactory{metadata: md}
}

// New Creates a new trigger instance for a given id
func (t *TCPTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &TCPTrigger{metadata: t.metadata, config: config}
}

// Metadata implements trigger.Trigger.Metadata
func (t *TCPTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Initialize implements trigger.Init.Initialize
func (t *TCPTrigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()
	t.handlers = ctx.GetHandlers()

	t.logger.Infof("Initialize implements trigger.Init.Initialize : %s", "")

	return nil
}

// Start implements trigger.Trigger.Start
func (t *TCPTrigger) Start() error {
	// start the trigger
	t.logger.Infof("start the trigger : %s", "")
	return nil
}

// Stop implements trigger.Trigger.Start
func (t *TCPTrigger) Stop() error {
	// stop the trigger
	t.logger.Infof("stop the trigger : %s", "")
	return nil
}
