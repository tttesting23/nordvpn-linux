package service

import (
	"fmt"
	"log"
	"sync"
)

type processType int

const (
	systemd processType = iota
	child
)

type Combined struct {
	mu               sync.Mutex
	systemd          SystemdNorduser
	childProcess     ChildProcessNorduser
	uidToProcessType map[uint32]processType
}

func NewNorduserService() Combined {
	return Combined{
		uidToProcessType: make(map[uint32]processType),
	}
}

func (c *Combined) Enable(uid uint32, gid uint32) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.systemd.Enable(uid)
	if err == nil {
		c.uidToProcessType[uid] = systemd
		return nil
	}

	log.Printf("failed to enable norduserd via systemd: %s, will fallback to fork implementation", err)
	if err := c.childProcess.Enable(uid, gid); err != nil {
		return fmt.Errorf("enabling norduserd via fork: %w", err)
	}

	c.uidToProcessType[uid] = child
	return nil
}

func (c *Combined) Disable(uid uint32) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	switch c.uidToProcessType[uid] {
	case systemd:
		if err := c.systemd.Disable(uid); err != nil {
			return fmt.Errorf("disabling systemd norduserd: %w", err)
		}
	case child:
		if err := c.childProcess.Stop(uid); err != nil {
			return fmt.Errorf("stopping fork norduserd: %w", err)
		}
	}

	delete(c.uidToProcessType, uid)

	return nil
}

func (c *Combined) stop(uid uint32, process processType) error {
	switch process {
	case systemd:
		if err := c.systemd.Stop(uid); err != nil {
			return fmt.Errorf("stopping systemd norduserd: %w", err)
		}
	case child:
		if err := c.childProcess.Stop(uid); err != nil {
			return fmt.Errorf("stopping fork norduserd: %w", err)
		}
	}

	delete(c.uidToProcessType, uid)

	return nil
}

func (c *Combined) Stop(uid uint32) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if processType, ok := c.uidToProcessType[uid]; ok {
		return c.stop(uid, processType)
	}

	return fmt.Errorf("uid not found")
}

func (c *Combined) StopAll() {
	var err error
	for uid, processType := range c.uidToProcessType {
		if err = c.stop(uid, processType); err != nil {
			log.Println("failed to stop norduser for user: ", err.Error())
		}
	}
}
