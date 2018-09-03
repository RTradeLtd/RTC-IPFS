package rtfs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/RTradeLtd/Temporal/rtfs"
)

const testPIN = "QmNZiPk974vDsPmQii3YbrMKfi12KTSNM7XMiYyiea4VYZ"

func TestInitialize(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	im, err := rtfs.Initialize("", "")
	if err != nil {
		t.Fatal(err)
	}
	info, err := im.Shell.ID()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(info)
}

func TestDHTFindProvs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	im, err := rtfs.Initialize("", "")
	if err != nil {
		t.Fatal(err)
	}
	err = im.DHTFindProvs("QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv", "10")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildCustomRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	im, err := rtfs.Initialize("", "")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := im.BuildCustomRequest(context.Background(), "127.0.0.1:5001", "dht/findprovs", nil, "QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestPin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	im, err := rtfs.Initialize("", "")
	if err != nil {
		t.Fatal(err)
	}
	err = im.Pin(testPIN)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetObjectFileSizeInBytes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	im, err := rtfs.Initialize("", "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = im.GetObjectFileSizeInBytes(testPIN)
	if err != nil {
		t.Fatal(err)
	}
}

func TestObjectStat(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	im, err := rtfs.Initialize("", "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = im.ObjectStat(testPIN)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseLocalPinsForHash(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	im, err := rtfs.Initialize("", "")
	if err != nil {
		t.Fatal(err)
	}
	exists, err := im.ParseLocalPinsForHash(testPIN)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal(err)
	}
}

func TestPubSub(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	im, err := rtfs.Initialize("", "")
	if err != nil {
		t.Fatal(err)
	}
	err = im.PublishPubSubMessage(im.PubTopic, "data")
	if err != nil {
		t.Fatal(err)
	}
}
