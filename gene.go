package main

type gene []byte

func NewGenes() (genes gene) {
	genes = make([]byte, 18)
	return genes
}
