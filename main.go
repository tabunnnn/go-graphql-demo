package main

import "tabu4n.me/graphql-demo/provider"

func main() {
	if err := provider.Start(); err != nil {
		panic(err)
		return
	}
	select {}
}
