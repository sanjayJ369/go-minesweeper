package maelstrom

import "math/rand"

const (
	INIT    = "init"
	INIT_OK = "init_ok"
)

type BaseMessage struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
	Body Body   `json:"body"`
}

type Body struct {
	Type string `json:"type"`

	// optional fields
	MsgId     int `json:"msg_id,omitempty"`
	InReplyTo int `json:"in_reply_to,omitempty"`
	Value     int `json:"value,omitempty"`

	// init fields
	NodeId  int `json: "node_id,omitempty"`
	NodeIds int `json:"node_ids,omitempty"`

	// error fields
	Code int    `json:"code,omitempty"`
	Text string `json:"text,omitempty"`
}

func setClientAddress(req BaseMessage) BaseMessage {
	var resp BaseMessage
	resp.Src = req.Dest // or node id
	resp.Dest = req.Src
	resp.Body = Body{}
	resp.Body.MsgId = rand.Int()
	resp.Body.InReplyTo = req.Body.MsgId
	return resp
}

func HandleInit(req BaseMessage) BaseMessage {
	resp := setClientAddress(req)
	resp.Body.Type = INIT_OK
	return resp
}
