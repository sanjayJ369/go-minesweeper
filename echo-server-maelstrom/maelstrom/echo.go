package maelstrom

import (
	"bufio"
	"echo-server/internal/echo"
	"encoding/json"
	"fmt"
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

func HandleEcho(req BaseMessage) BaseMessage {
	resp := setClientAddress(req)
	resp.Body.Type = ECHO_OK

	msg := req.Body.Echo
	res := echo.Echo(msg)

	resp.Body.Echo = res
	return resp
}

func (e *EchoServer) ListenAndServe() error {
	scanner := bufio.NewScanner(e.in)
	writer := bufio.NewWriter(e.out)
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
			e.lgr.Logf("initalizing server...")
			e.lgr.Logf("received: %v", req)
			resp = HandleInit(req)
		case ECHO:
			e.lgr.Logf("echoing: %v", req)
			resp = HandleEcho(req)
		}

		respJSON, err := json.Marshal(resp)
		if err != nil {
			e.lgr.Logf("unable to marshall response %v: %s", resp, err)
			return err
		}

		_, err = writer.WriteString(fmt.Sprintf("%s\n", respJSON))
		if err != nil {
			e.lgr.Logf("unable to write response %v: %s", respJSON, err)
			return err
		}

		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		e.lgr.Logf("unable to scan input buffer: %s", err)
		return err
	}

	return nil
}
