package configuration

// ServerConfiguration struct to hold Server config
type ServerConfiguration struct {
	Address           string                  `mapstructure:"address"`
	LogLevel          string                  `mapstructure:"logLevel"`
	Datastore         DatastoreConfiguration  `mapstructure:"datastore"`
	Storage           StorageConfiguration    `mapstructure:"storage"`
	FileTargets       FileTargetsConfguration `mapstructure:"file"`
	KubernetesTargets KubernetesConfiguration `mapstructure:"kubernetes"`
}

type DatastoreConfiguration struct {
	Path string `mapstructure:"path"`
}

type StorageConfiguration struct {
	File FileStorageConfiguration `mapstructure:"file"`
}

type FileStorageConfiguration struct {
	Path string `mapstructure:"path"`
}

// DumperConfiguration struct to hold Dumper config
type DumperConfiguration struct {
	LogLevel string `mapstructure:"logLevel"`
	Token    string `mapstructure:"token"`
	Server   string `mapstructure:"server"`
	Interval int    `mapstructure:"interval"`
}

type FileTargetsConfguration struct {
	Dir  string `mapstructure:"dir"`
	File string `mapstructure:"file"`
}

type KubernetesConfiguration struct {
	Addr      string `mapstructure:"addr"`
	Token     string `mapstructure:"token"`
	CaFile    string `mapstructure:"cafile"`
	Advertise string `mapstructure:"advertise"`
}
