package main

type UWSGIConfig struct {
	// URL to uWSGI stats server.
	// Defaults to "tcp://127.0.0.1:1717"
	URL string

	// Period defines how often to read stats in seconds.
	// Defaults to 1 second.
	Period *int64
}

type ConfigSettings struct {
	Input UWSGIConfig
}
