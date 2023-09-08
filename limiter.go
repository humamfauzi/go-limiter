package limiter

// Limiter is an interface that allows to limit the number of events that can happen in a given time.
type Limiter interface {
	// AddCount adds one event to the limiter.
	AddCount()

	// GetCount returns the current number of events in the limiter.
	GetCount() int

	// IsAllowed returns true if the event is allowed to happen. Some limiter have different kind of way to limit
	IsAllowed() bool
}
