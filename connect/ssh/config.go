package ssh

type Config struct {
	LocalEndpoint struct {
		Host string
		Port int
	}
	ServerEndpoint struct {
		Host string
		Port int
	}
	RemoteEndpoint struct {
		Host string
		Port int
	}
	RemoteUser    string

	// Full path to id_rsa key
	KeyPath string
}