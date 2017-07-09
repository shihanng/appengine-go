package guestbook

import (
	"appengine"
	"appengine/datastore"
)

// Key returns the key used for all guestbook entries.
func Key(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}
