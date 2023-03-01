#!/bin/bash
libPath="./build/lib"
libs=("ERC721A/contracts" "forge-std/src" "murky/src" "openzeppelin-contracts/contracts" "operator-filter-registry/src" "solmate/src" "utility-contracts/src")

for libElement in ${libs[@]}
do
  if [ ! -d "$libPath/${libElement%/*}" ]; then
     mkdir -p "$libPath/${libElement%/*}"
  fi

  if [ "`ls -A $libPath/${libElement%/*}`"="" ];then
     cp -rf "./lib/$libElement/." "$libPath/${libElement%/*}"
  fi
done

solcjs --optimize --include-path build/lib/ --base-path . --abi --bin src/market/SeaPort.sol -o ./build/SeaPort
solcjs --optimize --include-path build/lib/ --base-path . --abi --bin src/token/SeaDrop.sol -o ./build/SeaDrop
solcjs --optimize --include-path build/lib/ --base-path . --abi --bin src/token/ERC721SeaDrop.sol -o ./build/SeaDrop
solcjs --optimize --include-path build/lib/ --base-path . --abi --bin src/token/ERC721PartnerSeaDrop.sol -o ./build/SeaDrop
solcjs --optimize --include-path build/lib/ --base-path . --abi --bin src-upgradeable/src/footage/ERC721ClubLockUpgradeable.sol -o ./build/footage
solcjs --optimize --include-path build/lib/ --base-path . --abi --bin src-upgradeable/src/footage/ERC721ClubUpgradeable.sol -o ./build/footage
solcjs --optimize --include-path build/lib/ --base-path . --abi --bin src-upgradeable/src/footage/ERC721PartnerTokenUpgradeable.sol -o ./build/footage
solcjs --optimize --include-path build/lib/ --base-path . --abi --bin src-upgradeable/src/footage/ERC721TokenUpgradeable.sol -o ./build/footage


rm -rf "./tools/go-package"

if [ ! -d "./tools/go-package/seaport" ];then
    mkdir -p "./tools/go-package/seaport"
fi 
abigen --bin ./build/SeaPort/src_market_SeaPort_sol_Seaport.bin --abi ./build/SeaPort/src_market_SeaPort_sol_Seaport.abi --type=seaport --pkg=seaport --out=./tools/go-package/seaport/seaport.go


if [ ! -d "./tools/go-package/seadrop" ];then
    mkdir -p "./tools/go-package/seadrop"
fi 
abigen --bin ./build/SeaDrop/src_token_SeaDrop_sol_SeaDrop.bin --abi ./build/SeaDrop/src_token_SeaDrop_sol_SeaDrop.abi --type=seadrop --pkg=seadrop --out=./tools/go-package/seadrop/seadrop.go

if [ ! -d "./tools/go-package/seadrop/token" ];then
    mkdir -p "./tools/go-package/seadrop/token"
fi 
abigen --bin ./build/SeaDrop/src_token_ERC721SeaDrop_sol_ERC721SeaDrop.bin --abi ./build/SeaDrop/src_token_ERC721SeaDrop_sol_ERC721SeaDrop.abi --type=token --pkg=token --out=./tools/go-package/seadrop/token/erc721_seadrop.go

if [ ! -d "/tools/go-package/seadrop/partner" ];then
    mkdir -p "./tools/go-package/seadrop/partner"
fi 
abigen --bin ./build/SeaDrop/src_token_ERC721PartnerSeaDrop_sol_ERC721PartnerSeaDrop.bin --abi ./build/SeaDrop/src_token_ERC721PartnerSeaDrop_sol_ERC721PartnerSeaDrop.abi --type=partner --pkg=partner --out=./tools/go-package/seadrop/partner/erc721_partner_seadrop.go

if [ ! -d "/tools/go-package/footagedrop/partner" ];then
    mkdir -p "./tools/go-package/footagedrop/partner"
fi 
abigen --bin ./build/footage/src-upgradeable_src_footage_ERC721PartnerTokenUpgradeable_sol_ERC721PartnerTokenUpgradeable.bin --abi ./build/footage/src-upgradeable_src_footage_ERC721PartnerTokenUpgradeable_sol_ERC721PartnerTokenUpgradeable.abi --type=partner --pkg=partner --out=./tools/go-package/footagedrop/partner/erc721_partner_token_upgradeable.go


if [ ! -d "/tools/go-package/footagedrop/token" ];then
    mkdir -p "./tools/go-package/footagedrop/token"
fi 
abigen --bin ./build/footage/src-upgradeable_src_footage_ERC721TokenUpgradeable_sol_ERC721TokenUpgradeable.bin --abi ./build/footage/src-upgradeable_src_footage_ERC721TokenUpgradeable_sol_ERC721TokenUpgradeable.abi --type=token --pkg=token --out=./tools/go-package/footagedrop/token/erc721_token_upgradeable.go


if [ ! -d "/tools/go-package/footagedrop/club" ];then
    mkdir -p "./tools/go-package/footagedrop/club"
fi 
abigen --bin ./build/footage/src-upgradeable_src_footage_ERC721ClubUpgradeable_sol_ERC721ClubUpgradeable.bin --abi ./build/footage/src-upgradeable_src_footage_ERC721ClubUpgradeable_sol_ERC721ClubUpgradeable.abi --type=club --pkg=club --out=./tools/go-package/footagedrop/club/erc721_club_upgradeable.go


if [ ! -d "/tools/go-package/footagedrop/clublock" ];then
    mkdir -p "./tools/go-package/footagedrop/clublock"
fi 
abigen --bin ./build/footage/src-upgradeable_src_footage_ERC721ClubLockUpgradeable_sol_ERC721ClubLockUpgradeable.bin --abi ./build/footage/src-upgradeable_src_footage_ERC721ClubLockUpgradeable_sol_ERC721ClubLockUpgradeable.abi --type=clublock --pkg=clublock --out=./tools/go-package/footagedrop/clublock/erc721_clublock_upgradeable.go

rm -rf ./build

echo "Done"