package reload

// ReloadFunc are functions that are called when a reload is requested.
type ReloadFunc func(map[string]interface{}) error
