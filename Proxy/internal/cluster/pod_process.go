package cluster

import (
	database "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/persistent"
	"github.com/uptrace/bun"
)

type Pod struct {
	bun.BaseModel `bun:"table:Cluster"`
}

func ProcessPod(event Event, db *database.Queries) {
}
