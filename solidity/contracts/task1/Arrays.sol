pragma solidity ^0.8.20;

// 合并两个有序的数组
contract Arrays{

  //  arr1: [1, 3, 5]
//arr2: [2, 4, 6]
    function mergeSorted(uint[] calldata arr1, uint[] calldata arr2) public pure returns (uint[] memory){
        uint len1 = arr1.length;
        uint len2 = arr2.length;
        uint[] memory merged = new uint[](len1+len2);

        uint i = 0;
        uint j = 0;
        uint k = 0;
        while(i<len1 && j < len2){
            if(arr1[i] <= arr2[k]){
                merged[k]= arr1[i];
                i++;
            }else{
                   merged[k] = arr2[j];
                   j++; 
            }
            k++;
        }
        // 其他元素处理
        while(i<len1){
               merged[k] = arr1[i];
               i++;
               k++; 
        }
    while (j < len2) {
                merged[k] = arr2[j];
                j++;
                k++;
            }

    return merged;

    }
}