pragma solidity ^0.8.20;

contract StringDeal {
    
    // 反转一个字符串。输入 "abcde"，输出 "edcba"
    function dealString(string memory s) public pure returns (string memory) {
        bytes memory bytesStr = bytes(s);
        bytes memory reversedStr = new bytes(bytesStr.length);
        for (uint i = 0; i < bytesStr.length; i++) {
            reversedStr[i] = bytesStr[bytesStr.length - 1 - i];
        }
        return string(reversedStr);
    }


   
}