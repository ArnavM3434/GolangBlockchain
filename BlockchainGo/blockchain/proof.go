package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

//implementing proof of work algorithm

//force network to do work to add block to blockchain

//take data from block

//create counter (nonce) which starts at 0

//create a hash for the data plus the counter

//if hash meets set of requirements, we use that hash and sign the block

//Requirements:
//first few bytes must contain 0s, this is the difficulty

const Difficulty = 12 //in real blockchains this would actually go up, since computational power and number of miners goes up

type ProofOfWork struct {
	Block  *Block
	Target *big.Int //represents requirement above
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) //256 bits in the hash, Lsh left shifts the target

	pow := &ProofOfWork{b, target}

	return pow

}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data

}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte //8 * 32 is 256 bits

	nonce := 0

	for nonce < math.MaxInt64 { //almost an infinite loop
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:]) //turn hash into big integer

		if intHash.Cmp(pow.Target) == -1 { //hash is less than target, so we've signed the block
			break

		} else {
			nonce++

		}

	}
	fmt.Println()
	return nonce, hash[:]

}

func (pow *ProofOfWork) Validate() bool { //after running Run function, you'll have nonce that'll allow you to derive hash that met the target - why can't you just directly use the block's hash instead of running InitData again?
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1

}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num) //writing to the bytes buffer
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()

}
