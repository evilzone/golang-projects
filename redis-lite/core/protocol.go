package core

import (
	"errors"
	"strings"
)

var ErrorInvalidCommand = errors.New("invalid command")

type Request struct {
	Command Command
	Params  []string
}

type Response struct {
	Success bool
	Value   []byte
}

type Command struct {
	Cmd               string
	MinRequiredParams int
}

var (
	CMDGet  Command = Command{Cmd: "GET", MinRequiredParams: 1}
	CMDSet  Command = Command{Cmd: "SET", MinRequiredParams: 2}
	CMDDel  Command = Command{Cmd: "DEL", MinRequiredParams: 1}
	CMDPing Command = Command{Cmd: "PING", MinRequiredParams: 0}
)

func parseCommand(cmd string) (Command, error) {
	switch cmd {
	case CMDGet.Cmd:
		return CMDGet, nil
	case CMDSet.Cmd:
		return CMDSet, nil
	case CMDDel.Cmd:
		return CMDDel, nil
	case CMDPing.Cmd:
		return CMDPing, nil
	default:
		return Command{}, ErrorInvalidCommand
	}
}

func (c Command) isParamValid(params []string) bool {
	return len(params) >= c.MinRequiredParams
}

func parseProtocol(input string) (Request, error) {
	tokens := strings.Split(input, " ")

	if len(tokens) == 0 {
		return Request{}, errors.New("empty request")
	}

	command, err := parseCommand(tokens[0])

	if err != nil {
		return Request{}, err
	}

	params := tokens[1:]

	if !command.isParamValid(params) {
		return Request{}, errors.New("invalid params")
	}

	return Request{Command: command, Params: params}, nil
}
