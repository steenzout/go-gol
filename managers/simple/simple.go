//
// Copyright 2015 Rakuten Marketing LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package simple

import (
	"fmt"
	"sync"

	"github.com/mediaFORGE/gol"
)

// entry the internal structure that links a logger to a status.
type entry struct {
	channel chan *gol.LogMessage
	logger  gol.Logger
	status  bool
}

// Manager generic struct for a logger manager.
type Manager struct {
	capacity  int
	loggers   map[string]entry
	channel   chan *gol.LogMessage
	mutex     sync.Mutex
	waitGroup sync.WaitGroup
	status    bool
}

// New creates a simple implementation of a logger manager.
func New() gol.LoggerManager {
	return &Manager{
		loggers: make(map[string]entry),
		mutex:   sync.Mutex{},
	}
}

// Close closes the log message channel and waits for all processing to complete.
func (m *Manager) Close() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	close(m.channel)
	m.waitGroup.Wait()

	for _, e := range m.loggers {
		e.status = false
		close(e.channel)
		e.logger.Close()
	}
	m.status = false
}

// Deregister removes the logger with the given name from the manager.
func (m *Manager) Deregister(n string) (err error) {
	if _, ok := m.loggers[n]; ok {
		e := m.loggers[n]
		e.status = false
		close(e.channel)

		delete(m.loggers, n)
	} else {
		err = fmt.Errorf("No logger registered as %s", n)
	}
	return
}

// Disable sets the given logger name as disabled.
func (m *Manager) Disable(n string) (err error) {
	if _, ok := m.loggers[n]; ok {
		if m.loggers[n].status {
			m.loggers[n] = entry{
				logger: m.loggers[n].logger,
				status: false,
			}
		}
	} else {
		err = fmt.Errorf("No logger registered as %s", n)
	}
	return
}

// Enable sets the given logger name as enabled.
func (m *Manager) Enable(n string) (err error) {
	if _, ok := m.loggers[n]; ok {
		if !m.loggers[n].status {
			m.loggers[n] = entry{
				logger: m.loggers[n].logger,
				status: true,
			}
		}
	} else {
		err = fmt.Errorf("No logger registered as %s", n)
	}
	return
}

// IsEnabled returns the status of the given logger.
func (m *Manager) IsEnabled(n string) (st bool, err error) {
	if _, ok := m.loggers[n]; ok {
		st = m.loggers[n].status
		err = nil
	} else {
		st = false
		err = fmt.Errorf("No logger registered as %s", n)
	}
	return
}

// List returns the list of logger names.
func (m *Manager) List() (keys []string) {
	keys = make([]string, 0, len(m.loggers))
	for k := range m.loggers {
		keys = append(keys, k)
	}
	return
}

// Register registers the given logger with the given name.
func (m *Manager) Register(n string, l gol.Logger) error {
	if l == nil {
		return fmt.Errorf("Cannot register a nil logger")
	}
	lchan := make(chan *gol.LogMessage, 1)
	l.Run(lchan)
	m.loggers[n] = entry{
		channel: lchan,
		logger:  l,
		status:  true,
	}
	return nil
}

// Run start a goroutine that will distribute all messages in
// the LogManager channel to each registered and enabled logger.
func (m *Manager) Run(c chan *gol.LogMessage) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if !m.status {
		m.status = true
		m.waitGroup.Add(1)
		m.channel = c

		go m.process()
	}
}

// Send process log message.
func (m *Manager) Send(msg *gol.LogMessage) (err error) {
	if !m.status {
		return fmt.Errorf("manager.simple.LogManager is not running")
	}
	m.channel <- msg
	return nil
}

func (m *Manager) process() {
	for msg := range m.channel {
		m.sendMessageToLoggers(msg)
	}
	m.waitGroup.Done()
}

func (m *Manager) sendMessageToLoggers(msg *gol.LogMessage) {
	for _, e := range m.loggers {
		if e.status {
			e.channel <- msg
		}
	}
}

var _ gol.LoggerManager = (*Manager)(nil)
