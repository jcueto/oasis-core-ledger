module github.com/oasisprotocol/oasis-core-ledger

go 1.14

// Updates the version used in spf13/cobra (dependency via tendermint) as
// there is no release yet with the fix. Remove once an updated release of
// spf13/cobra exists and tendermint is updated to include it.
// https://github.com/spf13/cobra/issues/1091
replace github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2

require (
	github.com/oasisprotocol/oasis-core/go v0.20.9
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/zondax/ledger-go v0.12.1
)
