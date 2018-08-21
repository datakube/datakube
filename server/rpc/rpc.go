package rpc

//Server struct to hold RPC Server Information
type Server struct {
	services *Services
}

//Services struct to hold RPC Services Information
type Services struct {
	DumperService *DumperService
	FileHandleService *FileHandleService
}