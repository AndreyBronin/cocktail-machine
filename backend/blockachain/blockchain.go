// Blockchain
package blockachain

import "github.com/AndreyBronin/cocktail-machine/backend/coctailhal"

// BlockChain interface to public blockchain SDK
type BlockChain interface {
	CheckCoctail(contract string) coctailhal.Contains
}
