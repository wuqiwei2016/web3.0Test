pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title MyNFT
 * @dev ERC721标准NFT合约，支持铸造功能并关联元数据
 */
contract MyNFT is ERC721, Ownable {
    // 当前NFT的最大ID
    uint256 private _tokenIdCounter;

    /**
     * @dev 构造函数，初始化NFT的名称和符号
     */
    constructor() ERC721("MyNFT", "MNFT") Ownable(msg.sender) {
        // 初始计数器设为1，从ID 1开始铸造NFT
        _tokenIdCounter = 1;
    }

    /**
     * @dev 铸造NFT并关联元数据
     * @param recipient 接收NFT的地址
     * @param tokenURI NFT的元数据URI
     * @return 新铸造的NFT的tokenId
     */
    function mintNFT(address recipient, string memory tokenURI) public onlyOwner returns (uint256) {
        uint256 newItemId = _tokenIdCounter;
        _mint(recipient, newItemId);
        _setTokenURI(newItemId, tokenURI);
        _tokenIdCounter++;
        return newItemId;
    }

    /**
     * @dev 获取下一个将被铸造的NFT的ID
     * @return 下一个NFT的ID
     */
    function getNextTokenId() public view returns (uint256) {
        return _tokenIdCounter;
    }
}