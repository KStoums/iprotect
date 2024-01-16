package api

type DataServiceFactory interface {
	NewDataService(logger LoggerService) (DataService, error)
}

type DataService interface {
	GetAddressState(address string) bool
	AddBlockedAddress(address string)
	RemoveBlockAddress(address string)
}
