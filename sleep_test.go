package sleep

import (
	"context"
	"testing"
	"time"
)

func TestLightly(t *testing.T) {
	ch := make(chan struct{})
	close(ch)
	awoken := Lightly(time.Second, ch)
	if !awoken {
		t.Error("Lightly should have been awoken")
	}

	awoken = Lightly(10*time.Millisecond, nil)
	if awoken {
		t.Error("Lightly should not have been awoken")
	}
}

func TestLightlyCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	awoken := Lightly(time.Second, ctx.Done())
	if !awoken {
		t.Error("Lightly should have been awoken")
	}

	ctx = context.Background()
	awoken = Lightly(10*time.Millisecond, ctx.Done())
	if awoken {
		t.Error("Lightly should not have been awoken")
	}
}
