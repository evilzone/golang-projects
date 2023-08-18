package core

import (
	"redis-lite/storage"
	"strconv"
	"time"
)

type RequestProcessor interface {
	Process(request Request) (Response, error)
}

type CommandProcessor struct {
	Cache storage.Cache[string, []byte]
}

func (cp *CommandProcessor) Process(request Request) (Response, error) {
	switch request.Command {
	case CMDGet:
		return cp.ProcessGet(request)
	case CMDSet:
		return cp.ProcessSet(request)
	case CMDDel:
		return cp.ProcessDel(request)
	case CMDPing:
		return cp.ProcessPing(request)
	default:
		return Response{}, ErrorInvalidCommand
	}
}

func (cp *CommandProcessor) ProcessGet(request Request) (Response, error) {
	res, err := cp.Cache.Get(request.Params[0])
	if err != nil {
		return Response{}, err
	}
	return Response{Success: true, Value: res}, nil
}

func (cp *CommandProcessor) ProcessSet(request Request) (Response, error) {

	key := request.Params[0]
	value := request.Params[1]

	cp.Cache.Set(key, []byte(value))

	if len(request.Params) > 2 {
		ttl, err := strconv.Atoi(request.Params[2])
		if err != nil {
			cp.Cache.Delete(key)
			return Response{}, err
		}
		go cp.processExpiry(key, time.Duration(ttl)*time.Second)
	}
	return Response{Success: true}, nil
}

func (cp *CommandProcessor) ProcessDel(request Request) (Response, error) {
	cp.Cache.Delete(request.Params[0])
	return Response{Success: true}, nil
}

func (cp *CommandProcessor) ProcessPing(request Request) (Response, error) {
	return Response{Success: true, Value: []byte("PONG")}, nil
}

func (cp *CommandProcessor) processExpiry(key string, ttl time.Duration) {
	<-time.After(ttl)
	cp.Cache.Delete(key)
}
