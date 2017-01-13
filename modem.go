// Copyright (c) 2017 sitec systems GmbH
// All rights reserved. Use of this source code is goverened by a
// BSD-style license that can be found in the license file

// +build !windows

package gmodem

import (
	"bufio"
	"bytes"
	"io"
	"time"

	"github.com/pkg/term"
)

// Modem holds the parameters which are needed for the serial communication
// with the modem device
type Modem struct {
	// DevFile is the absolute path to the device file (e. g. /dev/ttyUSB1)
	DevFile string
	// Baudrate is the speed for data transfer between the you and the modem device (e. g. 115200)
	Baudrate int64
	// ReadTimeout is the timeout which is used for reading operations on the serial device. Some AT-Commands
	// needs longer to execute than others.
	ReadTimeout time.Duration

	t *term.Term
}

// Open opens the connection to ther serial interface with the defined
// parameters. The function configures the default read timeout for the serial
// interface
func (m *Modem) Open() error {
	t, err := term.Open(m.DevFile, term.Speed(int(m.Baudrate)), term.RawMode)
	if err != nil {
		return err
	}

	err = t.SetReadTimeout(m.ReadTimeout)
	if err != nil {
		return err
	}

	m.t = t

	return nil
}

// SetReadTimeout sets a new value for the read timeout on the serial interface. It altough updates
// the ReadTimeout attribute in the Modem structure.
func (m *Modem) SetReadTimeout(d time.Duration) error {
	err := m.t.SetReadTimeout(d)
	if err != nil {
		return err
	}

	m.ReadTimeout = d

	return nil
}

// Close closes the connection on the serial interface
func (m *Modem) Close() {
	if m.t != nil {
		m.t.Close()
	}
}

// SendAt sends a AT command over the serial line to the and returns the answer to the caller.
// The command will return an error if the Write or Read operation will fail. It doesn't return
// an error if the AT command fails.
// The function appends the closing <CR> automatically to the command.
func (m *Modem) SendAt(cmd string) (string, error) {
	var data bytes.Buffer
	buf := bufio.NewWriter(&data)

	_, err := m.t.Write([]byte(cmd + "\r\n"))
	if err != nil {
		return "", err
	}

	_, err = io.Copy(buf, m.t)
	if err != nil {
		return "", err
	}
	buf.Flush()
	m.t.Flush()

	return data.String(), nil
}
