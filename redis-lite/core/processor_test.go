package core

import (
	"testing"
	"time"
)

type MockCache struct {
	GetCalled bool
	SetCalled bool
	DelCalled bool
}

func (mc *MockCache) Get(key string) ([]byte, error) {
	mc.GetCalled = true
	return []byte("value"), nil
}

func (mc *MockCache) Set(K string, value []byte) {
	mc.SetCalled = true
}

func (mc *MockCache) Delete(K string) {
	mc.DelCalled = true
}

func TestProcessor(t *testing.T) {

	mockCache := &MockCache{}
	cmd_processor := CommandProcessor{Cache: mockCache}

	t.Run("Test Get", func(t *testing.T) {
		request := Request{Command: CMDGet, Params: []string{"key"}}
		cmd_processor.Process(request)

		if !mockCache.GetCalled {
			t.Errorf("Get not called")
		}
	})

	t.Run("Test Set", func(t *testing.T) {
		request := Request{Command: CMDSet, Params: []string{"key", "value"}}
		cmd_processor.Process(request)

		if !mockCache.SetCalled {
			t.Errorf("Set not called")
		}
	})

	t.Run("Test Del", func(t *testing.T) {
		request := Request{Command: CMDDel, Params: []string{"key"}}
		cmd_processor.Process(request)

		if !mockCache.SetCalled {
			t.Errorf("Delete not called")
		}
	})

	t.Run("Test Set with TTL", func(t *testing.T) {
		request := Request{Command: CMDSet, Params: []string{"key", "value", "2"}}
		cmd_processor.Process(request)

		if !mockCache.SetCalled {
			t.Errorf("Set not called")
		}

		time.Sleep(2 * time.Second)

		if !mockCache.DelCalled {
			t.Errorf("Delete not called for TTL")
		}
	})
}
