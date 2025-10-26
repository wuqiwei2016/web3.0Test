pragma solidity ^0.8.20;


//创建一个名为Voting的合约，包含以下功能：
//一个mapping来存储候选人的得票数
//一个vote函数，允许用户投票给某个候选人
//一个getVotes函数，返回某个候选人的得票数
//一个resetVotes函数，重置所有候选人的得票数
contract Voting {
    // 存储候选人的得票数（候选人地址 => 得票数）
    mapping(address =>uint256) private _votes;
    address[] private _candidates; // 记录候选人地址


     // 注册候选人
    function registerCandidate(address candidate) external{
        _candidates.push(candidate);
    }

   // 投票
    function vote(address candidate) external {
        _votes[candidate] += 1;
    } 

    // 查询票数
    function getVotes(address candidate) external view returns(uint256) {
        return _votes[candidate];
    }

    function resetVoTes() external{
        for(uint i = 0; i<_candidates.length; i++){
            _votes[_candidates[i]] = 0;
        }
    }
}