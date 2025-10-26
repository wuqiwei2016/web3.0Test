pragma solidity ^0.8.20;


contract number {
     // 整数转罗马数字 (1-3999)
    function intToRoman(uint num) public pure returns (string memory) {
        require(num > 0 && num < 4000, "Number must be between 1 and 3999");
        
        // 罗马数字对照表
        string[13] memory symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        uint[13] memory values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        
        string memory result = "";
        
        for (uint i = 0; i < 13; i++) {
            while (num >= values[i]) {
                result = string(abi.encodePacked(result, symbols[i]));
                num -= values[i];
            }
        }
        
        return result;
    }

    // 罗马数字转整数
    function romanToInt(string memory s) public pure returns (uint) {
        uint result = 0;
        uint prevValue = 0;
        
        bytes memory strBytes = bytes(s);
        
        for (uint i = strBytes.length; i > 0; i--) {
            uint currentValue = _getRomanValue(strBytes[i-1]);
            
            if (currentValue < prevValue) {
                result -= currentValue;
            } else {
                result += currentValue;
            }
            
            prevValue = currentValue;
        }
        
        return result;
    }
    
    // 辅助函数：获取单个罗马字符的数值
    function _getRomanValue(bytes1 char) private pure returns (uint) {
        if (char == "I") {
            return 1;
        }
        if (char == "V") {
            return 5;
        }
        if (char == "X") {
            return 10;
        }
        if (char == "L") {
            return 50;
        }
        if (char == "C") {
            return 100;
        }
        if (char == "D") {
            return 500;
        }
        if (char == "M") {
            return 1000;
        }
        return 0;
    }
}