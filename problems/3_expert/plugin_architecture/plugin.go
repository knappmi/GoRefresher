package pluginarch

import "plugin"

// Load attempts to open a Go plugin (.so) file.
func Load(name string) error { _, err := plugin.Open(name); return err }
