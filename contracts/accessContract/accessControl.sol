pragma solidity >=0.4.0 <0.6.2;

contract accessControlContract
{
    
    address admin = 0x21A018606490C031A8c02Bb6b992D8AE44ADD37f;
    mapping(bytes32 => address) whitePages;
    
    event newAddrRegistered(bytes32 indexed _id);
    event newAddrRemove(bytes32 indexed _id);
    
    
    /****************** Greeting ****************************/
    function greet() pure public returns(string memory)
    {
        return "Hello you have called the contract accessControl.sol";
    }
    /*******************************************************/
    
    
    function addAccountToRegister(bytes32 id, address account) public
    {
        assert (msg.sender == admin);
        whitePages[id] = account;
        
        emit newAddrRegistered(id);
    }
    
    function removeAccountFromRegister(bytes32 id) public
    {
        assert(msg.sender == admin);
        delete whitePages[id];
        
        emit newAddrRemove(id);
    }
    
    function getAddress(bytes32 id) public view returns(address) 
    {
        return whitePages[id];
    }
    
}
