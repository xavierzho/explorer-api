package dispatcher

import "testing"

func TestDispatcher_Next(t *testing.T) {
	var dispatcher = New([]string{"11", "22", "33"})
	for i := 0; i < 10; i++ {
		t.Logf("%+v\n", dispatcher.Next().APIKey)
	}
}
