package maelstrom

import (
	"bufio"
	"encoding/json"
	"io"
)

const (
	ECHO    = "echo"
	ECHO_OK = "echo_ok"
)

// EchoBody contain echo meessage
type EchoBody struct {
	Body
	Echo string `json:"echo,omitempty"`
}

type Logger interface {
	Logf(format string, args ...any)
}

type EchoServer struct {
	lgr Logger
	in  io.Reader
	out io.Writer
}

func NewEchoServer(in io.Reader, out io.Writer, lgr Logger) EchoServer {
	return EchoServer{
		lgr: lgr,
		in:  in,
		out: out,
	}
}

func (e *EchoServer) ListenAndServe() error {
	scanner := bufio.NewScanner(e.in)
	for scanner.Scan() {
		var req BaseMessage
		line := scanner.Bytes()
		err := json.Unmarshal(line, &req)
		if err != nil {
			e.lgr.Logf("unable to unmarshall data %s: %w", line, err)
			return err
		}
		var resp BaseMessage
		switch req.Body.Type {
		case INIT:
			resp = HandleInit(req)
		case ECHO:

		}

	}
}
