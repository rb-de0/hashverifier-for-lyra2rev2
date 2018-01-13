package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/rb-de0/lyra2rev2"
	"github.com/rb-de0/lyra2rev2/sha3"

	"github.com/aead/skein"
	"github.com/dchest/blake256"
)

type HashCalculator = func([]byte) []byte

func main() {

	defer func() {
        if r := recover(); r != nil {
			fmt.Println("Recover Panic Error:", r)
        }
    }()

	hashCalculators := map[string]HashCalculator{}

	hashCalculators["lyra2rev2"] = func(input []byte) []byte {
		result, err := lyra2rev2.Sum(input)
		if err != nil {
			panic(err)
		}
		return result
	}

	hashCalculators["blake"] = func(input []byte) []byte {
		blake := blake256.New()
		if _, err := blake.Write(input); err != nil {
			panic(err)
		}
		result := blake.Sum(nil)
		return result
	}

	hashCalculators["keccak"] = func(input []byte) []byte {
		keccak := sha3.NewKeccak256()
		if _, err := keccak.Write(input); err != nil {
			panic(err)
		}
		result := keccak.Sum(nil)
		return result
	}

	hashCalculators["cubehash"] = func(input []byte) []byte {
		return lyra2rev2.Cubehash256(input)
	}

	hashCalculators["lyra2"] = func(input []byte) []byte {
		result := make([]byte, 32)
		lyra2rev2.Lyra2(result, input, input, 1, 4, 4)
		return result
	}

	hashCalculators["skein"] = func(input []byte) []byte {
		var result [32]byte
		skein.Sum256(&result, input, nil)
		return result[:]
	}

	hashCalculators["bmw"] = func(input []byte) []byte {
		return lyra2rev2.Bmw256(input)
	}

	flag.Parse()
	hash := flag.Arg(0)
	calculator, exist := hashCalculators[hash]

	if exist {
		input := flag.Arg(1)
		inputBytes, err := hex.DecodeString(input)
		if err != nil {
			panic(err)
		}

		result := calculator(inputBytes)
		fmt.Println(hash + " result: " + hex.EncodeToString(result))
	} else {
		input := flag.Arg(0)
		inputBytes, err := hex.DecodeString(input)
		if err != nil {
			panic(err)
		}

		for key := range hashCalculators {
			result := hashCalculators[key](inputBytes)
			fmt.Println(key + " : " + hex.EncodeToString(result))
		}
	}
}
