package delayed

import (
	"testing"
	"time"

	delay "gx/ipfs/QmRJVNatYJwTAHgdSM1Xef9QVQ1Ch3XHdmcrykjP5Y4soL/go-ipfs-delay"
	datastore "gx/ipfs/QmVG5gxteQNEMhrS8prJSmU2C9rebtFuTd3SYZ5kE3YZ5k/go-datastore"
)

func TestDelayed(t *testing.T) {
	d := New(datastore.NewMapDatastore(), delay.Fixed(time.Second))
	now := time.Now()
	k := datastore.NewKey("test")
	err := d.Put(k, []byte("value"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = d.Get(k)
	if err != nil {
		t.Fatal(err)
	}
	if time.Since(now) < 2*time.Second {
		t.Fatal("There should have been a delay of 1 second in put and in get")
	}
}