package nhexporter

import (
	"fmt"
	"github.com/bitbandi/go-nicehash-api"
	"strings"
)

func AlgoFromStringToInt(algoName string) ([]nicehash.AlgoType, error) {
	switch strings.ToLower(algoName) {
	case "scrypt":
		return []nicehash.AlgoType{nicehash.AlgoTypeScrypt}, nil
	case "sha256":
		return []nicehash.AlgoType{nicehash.AlgoTypeSHA256}, nil
	case "scryptnf":
		return []nicehash.AlgoType{nicehash.AlgoTypeScryptNf}, nil
	case "x11":
		return []nicehash.AlgoType{nicehash.AlgoTypeX11}, nil
	case "x13":
		return []nicehash.AlgoType{nicehash.AlgoTypeX13}, nil
	case "keccak":
		return []nicehash.AlgoType{nicehash.AlgoTypeKeccak}, nil
	case "x15":
		return []nicehash.AlgoType{nicehash.AlgoTypeX15}, nil
	case "nist5":
		return []nicehash.AlgoType{nicehash.AlgoTypeNist5}, nil
	case "neoscrypt":
		return []nicehash.AlgoType{nicehash.AlgoTypeNeoScrypt}, nil
	case "lyra2re":
		return []nicehash.AlgoType{nicehash.AlgoTypeLyra2RE}, nil
	case "whirpoolx":
		return []nicehash.AlgoType{nicehash.AlgoTypeWhirlpoolX}, nil
	case "qubit":
		return []nicehash.AlgoType{nicehash.AlgoTypeQubit}, nil
	case "quark":
		return []nicehash.AlgoType{nicehash.AlgoTypeQuark}, nil
	case "axiom":
		return []nicehash.AlgoType{nicehash.AlgoTypeAxiom}, nil
	case "lyra2rev2":
		return []nicehash.AlgoType{nicehash.AlgoTypeLyra2REv2}, nil
	case "scryptjanenf16":
		return []nicehash.AlgoType{nicehash.AlgoTypeScryptJaneNf16}, nil
	case "blake256r8":
		return []nicehash.AlgoType{nicehash.AlgoTypeBlake256r8}, nil
	case "blake256r14":
		return []nicehash.AlgoType{nicehash.AlgoTypeBlake256r14}, nil
	case "blake256r8vnl":
		return []nicehash.AlgoType{nicehash.AlgoTypeBlake256r8vnl}, nil
	case "hodl":
		return []nicehash.AlgoType{nicehash.AlgoTypeHodl}, nil
	case "daggerhashimoto":
		return []nicehash.AlgoType{nicehash.AlgoTypeDaggerHashimoto}, nil
	case "decred":
		return []nicehash.AlgoType{nicehash.AlgoTypeDecred}, nil
	case "cryptonight":
		return []nicehash.AlgoType{nicehash.AlgoTypeCryptoNight}, nil
	case "lbry":
		return []nicehash.AlgoType{nicehash.AlgoTypeLbry}, nil
	case "equihash":
		return []nicehash.AlgoType{nicehash.AlgoTypeEquihash}, nil
	case "pascal":
		return []nicehash.AlgoType{nicehash.AlgoTypePascal}, nil
	case "x11gost":
		return []nicehash.AlgoType{nicehash.AlgoTypeX11Gost}, nil
	case "sia":
		return []nicehash.AlgoType{nicehash.AlgoTypeSia}, nil
	case "blake2s":
		return []nicehash.AlgoType{nicehash.AlgoTypeBlake2s}, nil
	case "max":
		return []nicehash.AlgoType{nicehash.AlgoTypeMAX}, nil
	}
	return []nicehash.AlgoType{}, fmt.Errorf("No Algo with name %s was found", algoName)
}
