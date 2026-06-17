package bookingappabfactory

import "fmt"

type Storer interface {
	SaveInStore() 
}

type Cacher interface {
	CacheThis() 
}

type Logger interface {
	Log()
}


type MemoryStore struct{}

func (ms MemoryStore) SaveInStore() {
	fmt.Println("Data successfully saved in MemoryStore.")
}

type PostgresStore struct{}

func (ps PostgresStore) SaveInStore() {
	fmt.Println("Data successfully persisted in PostgresStore database.")
}

type MemoryCache struct{}

func (mc MemoryCache) CacheThis() {
	fmt.Println("Item cached in local MemoryCache.")
}

type RedisCache struct{}

func (rc RedisCache) CacheThis() {
	fmt.Println("Item cached remotely in RedisCache cluster.")
}


type ConsoleLogger struct{}

func (cl ConsoleLogger) Log() {
	fmt.Println("[ConsoleLogger] Standard log message outputted to stdout.")
}

type ZapLogger struct{}

func (zl ZapLogger) Log() {
	fmt.Println("[ZapLogger] Structured JSON log entry produced efficiently.")
}


type InfraFactory interface {
	CreateStore() Storer 
	CreateCache() Cacher
	CreateLogger() Logger
}

type ProdFactory struct {}

type DevFactory struct {}

func (pf ProdFactory) CreateStore() Storer {
	return PostgresStore{}
}

func (pf ProdFactory) CreateCache() Cacher {
	return RedisCache{}
}

func (pf ProdFactory) CreateLogger() Logger {
	return ZapLogger{}
}

func (df DevFactory) CreateStore() Storer {
	return MemoryStore{}
}

func (df DevFactory) CreateCache() Cacher {
	return MemoryCache{}
}

func (df DevFactory) CreateLogger() Logger {
	return ConsoleLogger{}
}


func main(){

	var factory InfraFactory= ProdFactory{}
	logger:=factory.CreateLogger()
	Cacher:=factory.CreateCache()
	Storer:=factory.CreateStore()

	logger.Log()
	Cacher.CacheThis()
	Storer.SaveInStore()

}