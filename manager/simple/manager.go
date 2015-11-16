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

	"github.com/mediaFORGE/gol"
)

// entry the internal structure that links a logger to a status.
type entry struct {
	logger gol.Logger
	status bool
}

// Manager generic struct for a logger manager.
type Manager struct {
	loggers map[string]entry
}

var _ gol.LoggerManager = (*Manager)(nil)

// New creates a simple implementation of a logger manager.
func New() gol.LoggerManager {
	return &Manager{
		loggers: make(map[string]entry),
	}
}

// Deregister removes the logger with the given name from the manager.
func (m *Manager) Deregister(n string) (err error) {
	if _, ok := m.loggers[n]; ok {
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
	m.loggers[n] = entry{
		logger: l,
		status: true,
	}
	return nil
}
