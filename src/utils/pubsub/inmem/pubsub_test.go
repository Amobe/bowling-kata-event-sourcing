package inmem

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestSubscriber(t *testing.T) {
	ctx := context.Background()
	h := NewPubSub[data]()
	sub01 := NewSubscriber[data]("sub01", fakeNotifier)
	sub02 := NewSubscriber[data]("sub02", fakeNotifier)
	sub03 := NewSubscriber[data]("sub03", fakeNotifier)

	_ = h.Subscribe(ctx, sub01)
	_ = h.Subscribe(ctx, sub02)
	_ = h.Subscribe(ctx, sub03)

	assert.Equal(t, 3, h.Subscribers())

	_ = h.Unsubscribe(ctx, sub01)
	_ = h.Unsubscribe(ctx, sub02)
	_ = h.Unsubscribe(ctx, sub03)

	assert.Equal(t, 0, h.Subscribers())
}

type data struct {
	payload string
}

func (d data) String() string {
	return d.payload
}

func fakeNotifier(d data) {}
