package pool

import (
	"time"

	"github.com/go-goim/goim/pkg/log"
)

type Conn interface {
	Key() string           // connection unique key in pool
	Close() error          // close the connection
	Done() <-chan struct{} // check if ctx is canceled
	Err() error            // return error if conn has any internal error
}

type idleConn struct {
	p        *namedPool
	c        Conn
	stopChan chan struct{}
}

// close is different form stop
// close is closes the connection and delete it from pool
// stop is a trigger to stop the daemon then call the close
func (i *idleConn) close() {
	_ = i.c.Close()
	i.p.delete(i.c.Key())
}

func (i *idleConn) stop() {
	i.stopChan <- struct{}{}
}

func (i *idleConn) daemon() {
	var (
		timer = time.NewTimer(time.Second * 5)
	)
loop:
	for {

		select {
		case <-timer.C:
			timer.Reset(time.Second * 5)
			if i.c.Err() != nil {
				break loop
			}
		case <-i.stopChan:
			break loop
		case <-i.c.Done():
			log.Error("conn done", "key", i.c.Key())
			break loop
		}
	}

	log.Info("conn daemon exit", "key", i.c.Key())
	i.close()
}
