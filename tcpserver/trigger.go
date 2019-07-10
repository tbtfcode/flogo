package tcpserver

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMetadata = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

// Factory My Trigger factory
type Factory struct {
}

// Trigger is a stub for your Trigger implementation
type Trigger struct {
	logger      log.Logger
	handlers    []trigger.Handler
	settings    *Settings
	delimiter   byte
	listener    net.Listener
	connections []net.Conn
}

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

// Metadata implements trigger.Trigger.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMetadata
}

// New Creates a new trigger instance for a given id
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	settings := &Settings{}
	err := metadata.MapToStruct(config.Settings, settings, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{settings: settings}, nil
}

// Initialize implements trigger.Init.Initialize
func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()
	t.handlers = ctx.GetHandlers()

	host := t.settings.Host
	port := t.settings.Port
	delimiter := t.settings.Delimiter

	if port == "" {
		return errors.New("Valid port must be set")
	}

	if delimiter != "" {
		r, _ := utf8.DecodeRuneInString(delimiter)
		t.delimiter = byte(r)
	}

	listener, err := net.Listen(t.settings.Network, host+":"+port)
	if err != nil {
		return err
	}
	// defer listener.Close()
	t.listener = listener

	return err
}

func (t *Trigger) connectionWaiting() {
	for {
		// Listen for an incoming connection.
		conn, err := t.listener.Accept()
		if err != nil {
			errString := err.Error()
			if !strings.Contains(errString, "use of closed network connection") {
				t.logger.Error("Error accepting connection: ", err.Error())
			}
			return
		} else {
			t.logger.Debugf("Handling new connection from client - %s", conn.RemoteAddr().String())
			// Handle connections in a new goroutine.
			go t.connectionHandler(conn)
		}
	}
}

// Start implements trigger.Trigger.Start
func (t *Trigger) Start() error {
	go t.connectionWaiting()
	t.logger.Infof("Started listener on Port - %s, Network - %s", t.settings.Port, t.settings.Network)
	return nil
}

func (t *Trigger) connectionHandler(conn net.Conn) {

	//Gather connection list for later cleanup
	t.connections = append(t.connections, conn)

	for {

		if t.settings.TimeOut > 0 {
			t.logger.Info("Setting timeout: ", t.settings.TimeOut)
			conn.SetDeadline(time.Now().Add(time.Duration(t.settings.TimeOut) * time.Millisecond))
		}

		output := &Output{}

		if t.delimiter != 0 {
			data, err := bufio.NewReader(conn).ReadBytes(t.delimiter)
			if err != nil {
				errString := err.Error()
				if !strings.Contains(errString, "use of closed network connection") {
					t.logger.Error("Error reading data from connection: ", err.Error())
				} else {
					t.logger.Info("Connection is closed.")
				}
				if nerr, ok := err.(net.Error); !ok || !nerr.Timeout() {
					// Return if not timeout error
					return
				}

			} else {
				output.Content = string(data[:len(data)-1])
			}
		} else {
			var buf bytes.Buffer
			_, err := io.Copy(&buf, conn)
			if err != nil {
				errString := err.Error()
				if !strings.Contains(errString, "use of closed network connection") {
					t.logger.Error("Error reading data from connection: ", err.Error())
				} else {
					t.logger.Info("Connection is closed.")
				}
				if nerr, ok := err.(net.Error); !ok || !nerr.Timeout() {
					// Return if not timeout error
					return
				}
			} else {
				output.Content = string(buf.Bytes())
			}
		}
		t.logger.Info(output.Content)
		/*
			if output.Content != "" {
				var replyData []string
				for i := 0; i < len(t.handlers); i++ {
					results, err := t.handlers[i].Handle(context.Background(), output)
					if err != nil {
						t.logger.Error("Error invoking action : ", err.Error())
						continue
					}

					reply := &Reply{}
					err = reply.FromMap(results)
					if err != nil {
						t.logger.Error("Failed to convert flow output : ", err.Error())
						continue
					}
					if reply.Reply != "" {
						replyData = append(replyData, reply.Reply)
					}
				}

				if len(replyData) > 0 {
					replyToSend := strings.Join(replyData, string(t.delimiter))
					// Send a response back to client contacting us.
					_, err := conn.Write([]byte(replyToSend + "\n"))
					if err != nil {
						t.logger.Error("Failed to write to connection : ", err.Error())
					}
				}
			}
		*/
	}
}

// Stop implements trigger.Trigger.Start
func (t *Trigger) Stop() error {

	for i := 0; i < len(t.connections); i++ {
		t.connections[i].Close()
	}

	t.connections = nil

	if t.listener != nil {
		t.listener.Close()
	}

	t.logger.Info("Stopped listener")

	return nil
}
