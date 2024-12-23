package greedy

/**
134. 加油站
level: medium

在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/

/**
这个问题的核心是通过遍历环路判断是否有一个加油站可以作为出发点，绕一圈后油量足够。通过分析这个问题，可以得出以下几点结论：

油量和消耗的差值：对于每个加油站 i，定义 net[i] = gas[i] - cost[i]，即从加油站 i 出发，油量与消耗之间的差值。如果 net[i] 为正，说明从该加油站出发到下一站油量是足够的；如果为负，说明从该加油站出发不能到达下一站。

环路的总和：如果从某个加油站出发的总油量差 net[i] 为负数（累加），那么绕一圈一共消耗的油量会超过加油站提供的油量，显然无法绕圈。所以，只要从某个起点到终点的总油量差是非负的，就一定能成功绕圈。

思路：
首先计算每个加油站的 net[i]，表示从该加油站出发的油量差。
如果所有加油站的总油量差小于 0，直接返回 -1，因为不可能绕圈。
如果有可能绕圈，尝试从每个加油站开始，判断是否能绕一圈。
采用贪心策略，从一个加油站出发，逐站遍历，如果在某一站油量为负，说明无法从当前起点出发绕圈，可以尝试从下一个加油站开始。
*/
func canCompleteCircuit(gas []int, cost []int) int {
    totalGas, totalCost := 0, 0
    for _, g := range gas {
        totalGas += g
    }
    for _, c := range cost {
        totalCost += c
    }
    if totalCost > totalGas {
        return -1
    }
    currentGas, start := 0, 0
    for i := 0; i < len(gas); i++ {
        currentGas += gas[i] - cost[i]
        if currentGas < 0 {
            start = i + 1
            currentGas = 0
        }
    }
    return start
}
