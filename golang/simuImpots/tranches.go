package main

type TrancheBuilder interface {
	Build() []OutBoundTranche
}
