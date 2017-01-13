

# gmodem
`import "github.com/sitec-systems/gmodem"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package gmodem provides simple primitives to open a connection to a serial
modem device and communicate with it by AT-Commands




## <a name="pkg-index">Index</a>
* [type Modem](#Modem)
  * [func (m *Modem) Close()](#Modem.Close)
  * [func (m *Modem) Open() error](#Modem.Open)
  * [func (m *Modem) SendAt(cmd string) (string, error)](#Modem.SendAt)
  * [func (m *Modem) SetReadTimeout(d time.Duration) error](#Modem.SetReadTimeout)


#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/sitec-systems/gmodem/doc.go) [modem.go](/src/github.com/sitec-systems/gmodem/modem.go) 






## <a name="Modem">type</a> [Modem](/src/target/modem.go?s=372:784#L10)
``` go
type Modem struct {
    // DevFile is the absolute path to the device file (e. g. /dev/ttyUSB1)
    DevFile string
    // Baudrate is the speed for data transfer between the you and the modem device (e. g. 115200)
    Baudrate int64
    // ReadTimeout is the timeout which is used for reading operations on the serial device. Some AT-Commands
    // needs longer to execute than others.
    ReadTimeout time.Duration
    // contains filtered or unexported fields
}
```
Modem holds the parameters which are needed for the serial communication
with the modem device










### <a name="Modem.Close">func</a> (\*Modem) [Close](/src/target/modem.go?s=1550:1573#L55)
``` go
func (m *Modem) Close()
```
Close closes the connection on the serial interface




### <a name="Modem.Open">func</a> (\*Modem) [Open](/src/target/modem.go?s=949:977#L25)
``` go
func (m *Modem) Open() error
```
Open opens the connection to ther serial interface with the defined
parameters. The function configures the default read timeout for the serial
interface




### <a name="Modem.SendAt">func</a> (\*Modem) [SendAt](/src/target/modem.go?s=1912:1962#L65)
``` go
func (m *Modem) SendAt(cmd string) (string, error)
```
SendAt sends a AT command over the serial line to the and returns the answer to the caller.
The command will return an error if the Write or Read operation will fail. It doesn't return
an error if the AT command fails.
The function appends the closing <CR> automatically to the command.




### <a name="Modem.SetReadTimeout">func</a> (\*Modem) [SetReadTimeout](/src/target/modem.go?s=1340:1393#L43)
``` go
func (m *Modem) SetReadTimeout(d time.Duration) error
```
SetReadTimeout sets a new value for the read timeout on the serial interface. It altough updates
the ReadTimeout attribute in the Modem structure.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)