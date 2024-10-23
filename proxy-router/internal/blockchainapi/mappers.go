package blockchainapi

import (
	"encoding/hex"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi/structs"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	m "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/marketplace"
	s "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
	"github.com/ethereum/go-ethereum/common"
)

func mapBids(bidIDs [][32]byte, bids []m.IBidStorageBid) []*structs.Bid {
	result := make([]*structs.Bid, len(bidIDs))
	for i, value := range bids {
		result[i] = mapBid(bidIDs[i], value)
	}
	return result
}

func mapBid(bidID common.Hash, bid m.IBidStorageBid) *structs.Bid {
	return &structs.Bid{
		Id:             bidID,
		ModelAgentId:   bid.ModelId,
		Provider:       bid.Provider,
		Nonce:          &lib.BigInt{Int: *bid.Nonce},
		CreatedAt:      &lib.BigInt{Int: *bid.CreatedAt},
		DeletedAt:      &lib.BigInt{Int: *bid.DeletedAt},
		PricePerSecond: &lib.BigInt{Int: *bid.PricePerSecond},
	}
}

func mapSessions(sessionIDs [][32]byte, sessions []s.ISessionStorageSession, bids []m.IBidStorageBid) []*structs.Session {
	result := make([]*structs.Session, len(sessionIDs))
	for i := 0; i < len(sessionIDs); i++ {
		result[i] = mapSession(sessionIDs[i], sessions[i], bids[i])
	}
	return result
}

func mapSession(ID common.Hash, ses s.ISessionStorageSession, bid m.IBidStorageBid) *structs.Session {
	return &structs.Session{
		Id:                      lib.BytesToString(ID[:]),
		Provider:                bid.Provider,
		User:                    ses.User,
		ModelAgentId:            lib.BytesToString(bid.ModelId[:]),
		BidID:                   lib.BytesToString(ses.BidId[:]),
		Stake:                   ses.Stake,
		PricePerSecond:          bid.PricePerSecond,
		CloseoutReceipt:         hex.EncodeToString(ses.CloseoutReceipt),
		CloseoutType:            ses.CloseoutType,
		ProviderWithdrawnAmount: ses.ProviderWithdrawnAmount,
		OpenedAt:                ses.OpenedAt,
		EndsAt:                  ses.EndsAt,
		ClosedAt:                ses.ClosedAt,
	}
}
