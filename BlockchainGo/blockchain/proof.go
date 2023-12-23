package blockchain

import "math/big"

//implementing proof of work algorithm

//force network to do work to add block to blockchain

//take data from block

//create counter (nonce) which starts at 0

//create a hash fo the data plus the counter

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
	target.Lsh(target, uint(256-Difficulty)) //256 bytes in the hash, Lsh left shifts, so you can get the first 12 bits

	pow := &ProofOfWork{b, target}

	return pow

}
