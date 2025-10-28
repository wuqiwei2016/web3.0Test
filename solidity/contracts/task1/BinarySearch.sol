pragma solidity ^0.8.20;

contract BinarySearch{


    // 二分查找
    function searcht(uint[] calldata arr,uint target) public pure returns (int){
        uint left = 0;
        uint right = arr.length-1;
        
        while(left <= right){
             // 计算中间索引（使用这种方式避免整数溢出）
            uint mid = left + (right - left) / 2;
            
            // 找到目标值
            if (arr[mid] == target) {
                return int(mid);
            }
            // 目标值在右半部分
            else if (arr[mid] < target) {
                left = mid + 1;
            }
            // 目标值在左半部分
            else {
                // 避免uint下溢
                if (mid == 0) {
                    break;
                }
                right = mid - 1;
            }
        }
        
        // 未找到目标值
        return -1;
        }


       // 辅助函数：验证数组是否为升序排列
    function isSorted(uint[] calldata arr) public pure returns (bool) {
        if (arr.length <= 1) {
            return true;
        }
        
        for (uint i = 0; i < arr.length - 1; i++) {
            if (arr[i] > arr[i + 1]) {
                return false;
            }
        }
        
        return true;
    }
}