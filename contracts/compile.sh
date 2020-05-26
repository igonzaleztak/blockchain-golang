# Compile dataContract
abigen -sol dataContract/data_ledger.sol -pkg dataContract --out=dataContract/dataContract.go

# Compile accessControlContract
abigen -sol accessContract/accessControl.sol -pkg accessControlContract -out accessContract/accessControlContract.go

# Compile balanceContract
abigen --sol balanceContract/balance.sol --pkg balanceContract --out balanceContract/balanceContract.go