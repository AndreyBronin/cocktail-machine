// Blockchain
package blockachain

import "github.com/AndreyBronin/cocktail-machine/backend/coctalhal"

// BlockChain interface to public blockchain SDK
type BlockChain interface {
	CheckCoctail(contract string) coctalhal.Contains
}
