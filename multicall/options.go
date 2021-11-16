package multicall

import "fmt"

type Option func(*Config)

type Config struct {
	MulticallAddress string
	Gas              string
}

const (
	MainnetAddress = "0x5eb3fa2dfecdde21c950813c665e9364fa609bd2"
	RopstenAddress = "0xf3ad7e31b052ff96566eedd218a823430e74b406"
	KovanAddress   = "0x2cc8688c5f75e365aaeeb4ea8d6a480405a48d2a"
	RinkebyAddress = "0x42ad527de7d4e9d9d011ac45b31d8551f8fe9821"
	GorliAddress   = "0x77dca2c955b15e9de4dbbcf1246b4b85b651e50e"
	xDaiAddress    = "0xb5b692a88bdfc81ca69dcb1d924f59f0413a602a"
	PolygonAddress = "0x11ce4B23bD875D7F5C6a31084f55fDe1e9A87507"
	MumbaiAddress  = "0x08411ADd0b5AA8ee47563b146743C13b3556c9Cc"
	RSKMainnet     = "0x6c62bf5440de2cb157205b15c424bceb5c3368f5"
	RSKTestnet     = "0x9e469e1fc7fb4c5d17897b68eaf1afc9df39f103"
	BSCMainnet     = "0x41263cba59eb80dc200f3e2544eda4ed6a90e76c"
	BSCTestnet     = "0xae11C5B5f29A6a25e955F0CB8ddCc416f522AF5C"
)

func ContractAddress(address string) Option {
	return func(c *Config) {
		c.MulticallAddress = address
	}
}

func SetGas(gas uint64) Option {
	return func(c *Config) {
		c.Gas = fmt.Sprintf("0x%x", gas)
	}
}

func SetGasHex(gas string) Option {
	return func(c *Config) {
		c.Gas = gas
	}
}
