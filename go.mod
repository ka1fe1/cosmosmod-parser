module github.com/kaifei-bianjie/cosmosmod-parser

go 1.16

require (
	github.com/bianjieai/spartan-cosmos v1.0.2
	github.com/cosmos/cosmos-sdk v0.45.5-0.20220523154235-2921a1c3c918
	github.com/gogo/protobuf v1.3.3
	github.com/stretchr/testify v1.7.5
	github.com/tendermint/tendermint v0.35.0
	github.com/tharsis/ethermint v0.10.3
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.45.1-irita-20220816.0.20220816095307-845547d9c19e
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20210908054213-781a5fed16d6
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.10.2-irita-20220826
)