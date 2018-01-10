package services

//Services Type to expose Services to RPC and HTTP
type Services struct {
	DumperService DumperService
	TargetService TargetService
}
