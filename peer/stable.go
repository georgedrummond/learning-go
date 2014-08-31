package peer

import (
	"fmt"
)

type Server struct {
	Ip        string
	Signature string
}

func StableList() {
	stable_servers := []*Server{
		&Server{Ip: "124.33.23.554", Signature: "def"},
		&Server{Ip: "124.33.23.553", Signature: "dsdsdef"},
	}

	fmt.Println("Stable servers: ...")
	fmt.Println(stable_servers[0].Ip)
}
