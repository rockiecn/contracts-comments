cd ..
abigen --sol ./contracts-sol/contracts/role/RoleFS.sol --out ./contracts/rolefs/RoleFS.go --pkg rolefs --type RoleFS --bin ~/Documents/ContractABI/RoleFS.bin
abigen --sol ./contracts-sol/contracts/role/Role.sol --out ./contracts/role/Role.go --pkg role --type Role --bin ~/Documents/ContractABI/Role.bin
abigen --sol ./contracts-sol/contracts/pledgePool/PledgePool.sol --out ./contracts/pledgepool/PledgePool.go --pkg pledgepool --type PledgePool --bin ~/Documents/ContractABI/PledgePool.bin
abigen --sol ./contracts-sol/contracts/fileSystem/FileSys.sol --out ./contracts/filesystem/FileSys.go --pkg filesys --type FileSys --bin ~/Documents/ContractABI/FileSys.bin
abigen --sol ./contracts-sol/contracts/role/Issuance.sol --out ./contracts/issu/Issuance.go --pkg issu --type Issuance --bin ~/Documents/ContractABI/Issuance.bin
abigen --sol ./contracts-sol/contracts/token/ERC20.sol --out ./contracts/erc20/ERC20.go --bin ~/Documents/ContractABI/ERC20.bin --pkg erc20 --type ERC20
