package tcpserver

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Network   string `md:"network"`       // The network type
	Host      string `md:"host"`          // The host name or IP for TCP server.
	Port      string `md:"port,required"` // The port to listen on
	Delimiter string `md:"delimiter"`     // Data delimiter for read and write
	TimeOut   int    `md:"timeout"`
}

type HandlerSettings struct {
}

type Output struct {
	//	Content interface{} `md:"content"`     // incomming data
	Content string `md:"content"` // incomming data
}

type Reply struct {
	Code int    `md: "code"`
	Data string `md: "data"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"content": o.Content,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Content, err = coerce.ToString(values["content"])
	if err != nil {
		return err
	}

	return nil
}

func (r *Reply) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code": r.Code,
		"data": r.Data,
	}
}

func (r *Reply) FromMap(values map[string]interface{}) error {

	var err error

	r.Code, err = coerce.ToInt(values["code"])
	if err != nil {
		return err
	}

	r.Data, err = coerce.ToString(values["data"])
	if err != nil {
		return err
	}

	return nil
}
