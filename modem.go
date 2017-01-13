package gmodem

import (
	"time"

	"github.com/pkg/term"
)

type Modem struct {
	DevFile     string
	Baudrate    int64
	ReadTimeout time.Duration
	t           *term.Term
}

// Open opens the connection to ther serial interface with the defined
// parameters
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

func (m *Modem) Close() {
	if m.t != nil {
		m.t.Close()
	}
}

func (m *Modem) At(cmd string) (string, error) {
	_, err := m.t.Write([]byte(cmd))
	if err != nil {
		return "", err
	}

	return "", nil
}
