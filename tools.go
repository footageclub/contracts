package tools

//go:generate solcjs --abi --include-path node_modules/ --base-path . contracts/LandToken.sol contracts/Club.sol contracts/ClubERC20.sol contracts/ClubNFT.sol contracts/ClubNFTPublic.sol
//go:generate solcjs --bin --include-path node_modules/ --base-path . contracts/LandToken.sol contracts/Club.sol contracts/ClubERC20.sol contracts/ClubNFT.sol contracts/ClubNFTPublic.sol
//go:generate abigen --bin contracts_LandToken_sol_LandToken.bin --abi contracts_LandToken_sol_LandToken.abi --type=LandToken --pkg=token --out=nft_go/token/land_token.go
//go:generate abigen --bin contracts_Club_sol_Club.bin --abi contracts_Club_sol_Club.abi --type=Club --pkg=token --out=nft_go/token/club.go
//go:generate abigen --bin contracts_ClubERC20_sol_ClubERC20.bin --abi contracts_ClubERC20_sol_ClubERC20.abi --type=ClubERC20 --pkg=token --out=nft_go/token/club_erc20.go
//go:generate abigen --bin contracts_ClubNFT_sol_ClubNFT.bin --abi contracts_ClubNFT_sol_ClubNFT.abi --type=ClubNFT --pkg=token --out=nft_go/token/club_nft.go
//go:generate abigen --bin contracts_ClubNFTPublic_sol_ClubNFTPublic.bin --abi contracts_ClubNFTPublic_sol_ClubNFTPublic.abi --type=ClubNFTPublic --pkg=token --out=nft_go/token/club_nft_public.go
