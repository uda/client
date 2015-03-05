package engine

import (
	"testing"

	"github.com/keybase/client/go/libkb"
)

// TestTrackerList creates a new fake user and has that user track
// t_alice.  It then uses the TrackerList engine to get the
// t_alice's trackers and makes sure that the new fake user is in
// the list.
func TestTrackerList(t *testing.T) {
	tc := SetupEngineTest(t, "trackerlist")
	defer tc.Cleanup()

	fu := CreateAndSignupFakeUser(t, "login")
	trackAlice(t, fu)

	uid := libkb.UsernameToUID("t_alice")
	e := NewTrackerList(&uid)
	ctx := &Context{LogUI: tc.G.UI.GetLogUI()}
	if err := RunEngine(e, ctx, nil, nil); err != nil {
		t.Fatal(err)
	}
	buid := libkb.UsernameToUID(fu.Username)
	trackers := e.List()
	if len(trackers) == 0 {
		t.Errorf("t_alice tracker count: 0.  expected > 0.")
	}

	found := false
	for _, x := range trackers {
		if x.Tracker == buid {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("fake user %q (%s) not included in list of t_alice trackers.", fu.Username, buid)
		t.Logf("tracker list:")
		for i, x := range trackers {
			t.Logf("%d: %s, %d, %d", i, x.Tracker, x.Status, x.Mtime)
		}
	}
}
