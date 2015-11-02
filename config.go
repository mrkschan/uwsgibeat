package main

type UWSGIConfig struct {
	// URL to uWSGI stats server.
	// Defaults to "127.0.0.1:1717"
	URL string

	// The method used to dial to uWSGI stats server.
	// It can be either "tcp", "unix", or "http".
	// Defaults "tcp".
	Method string

	// Period defines how often to read stats in seconds.
	// Defaults to 1 second.
	Period *int64
}

type ConfigSettings struct {
	Input UWSGIConfig
}
