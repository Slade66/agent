---
name: leetcode-what-solution
description: 用工程视角判断 LeetCode 解法是否值得写和讲，强调算法要易于理解、能维护、能扩展、能迁移到业务场景。Use when：需要选择、解释、评价 LeetCode 解法，尤其是在 AC trick 与通用工程解法之间做取舍。
---

# 你要给出什么样的算法？

## 一、核心判断

- **核心：** 业务需求永远在变，通用框架才是工程思维。

- **取舍：** 当一道题有很多解法时，我更关心有没有一种算法能普适于很多业务场景。

- **边界：** 如果某种算法的应用很局限，只是为了通过刷题 AC，那它就没有被我看到的意义——因为我不是做题家，我讲究实用主义。

## 二、例题：Increasing Triplet Subsequence

- **原题：** 判断数组里是否存在三个下标 `i < j < k`，并且满足 `nums[i] < nums[j] < nums[k]`。

    ```md
    # 334. Increasing Triplet Subsequence

    ## Description

    - Given an integer array `nums`, return `true` if there exist indices `i`, `j`, and `k` satisfying `i < j < k` and `nums[i] < nums[j] < nums[k]`.

    - If no such indices exist, return `false`.

    ## Examples

    ### Example 1

    - Input: `nums = [1,2,3,4,5]`

    - Output: `true`

    - Explanation: Any triplet where `i < j < k` is valid.

    ### Example 2

    - Input: `nums = [5,4,3,2,1]`

    - Output: `false`

    - Explanation: No triplet exists.

    ### Example 3

    - Input: `nums = [2,1,5,0,4,6]`

    - Output: `true`

    - Explanation: One valid triplet is `(1, 4, 5)` because `nums[1] == 1 < nums[4] == 4 < nums[5] == 6`.

    ## Constraints

    - `1 <= nums.length <= 5 * 10^5`

    - `-2^31 <= nums[i] <= 2^31 - 1`

    ## Follow up

    - Implement the solution with `O(n)` time complexity and `O(1)` space complexity.
    ```

## 三、不好的算法：双变量贪心

### 代码

- **写法：** 这个写法的时间复杂度是 `O(n)`，空间复杂度是 `O(1)`，但它是一个针对特定题目的 trick。

    ```go
    func increasingTriplet(nums []int) bool {
        first, second := math.MaxInt64, math.MaxInt64

        for _, num := range nums {
            if num <= first {
                first = num        // 找到更小的开头
            } else if num <= second {
                second = num       // 找到中间的数
            } else {
                return true        // num > second，找到了！
            }
        }
        return false
    }
    ```

### 为什么它能工作

- **核心问题：** 为什么可以更新 `first`？

    - 这是最容易困惑的地方。假设当前状态是 `first = 3`，`second = 5`，后面来了一个 `1`。

    - 你更新 `first = 1`，此时 `first` 这个 `1` 在 `second` 这个 `5` 的后面。

    - 但这没关系！因为 `second = 5` 已经证明前面存在过一个比 `5` 小的数，也就是那个 `3`。

    - 现在 `first = 1` 虽然出现在后面，但只要后面再出现一个 `num > 5`，就一定满足“某个比 `5` 小的数 < `5` < `num`”。

    - 换句话说，`second` 本身就已经携带了“前面有比它小的数”这个信息。`first` 更新得更小，只会让后续更容易找到第三个数。

### 为什么它不好

- **可扩展性：** 如果业务需求从“找长度为 3 的递增序列”变成“找 5 个或 10 个...”，这个双变量写法就不够用了，你要重写整个逻辑，再加 `third`、`fourth` 之类的变量。

- **理解成本：** `first` 被替换的逻辑非常反直觉。半年后你或者同事来排查 bug，看到 `first = 1`、`second = 8`，第一反应肯定是去数组里找这两个值的位置。结果发现 `1` 出现在 `8` 的后面，整个人会陷入迷茫：“这代码到底在算什么？”

- **面试说服力：** 面试官会认可你的代码能力，但也会担心你写“无法扩展的 trick code”。

- **总结：** 它是“针对特定输入的 hack”，不是工程解法。

## 四、更好的解法：状态真实，可迁移

- **总览：** 至少还有两种常见思路，而且它们比贪心法更容易理解，只是效率不同。

### 解法 1：预处理

- **核心思路：** 对于每个位置 `j`，只要它左边有一个更小的数，右边有一个更大的数，就找到了递增三元组。

- **代码：** `leftMin[i]` 表示 `nums[0..i]` 的最小值，`rightMax[i]` 表示 `nums[i..n-1]` 的最大值。

    ```go
    func increasingTriplet(nums []int) bool {
        n := len(nums)
        if n < 3 {
            return false
        }

        leftMin := make([]int, n)
        rightMax := make([]int, n)

        leftMin[0] = nums[0]
        for i := 1; i < n; i++ {
            leftMin[i] = min(leftMin[i-1], nums[i])
        }

        rightMax[n-1] = nums[n-1]
        for i := n - 2; i >= 0; i-- {
            rightMax[i] = max(rightMax[i+1], nums[i])
        }

        for j := 1; j < n-1; j++ {
            if leftMin[j-1] < nums[j] && nums[j] < rightMax[j+1] {
                return true
            }
        }
        return false
    }
    ```

- **复杂度：** 时间复杂度是 `O(n)`，空间复杂度是 `O(n)`。

- **可读性：** 逻辑直观，不用理解反直觉的贪心状态。可读性即正义，维护成本比 CPU 周期更贵；这个逻辑半年后你自己回来维护，或者给非算法背景的同事 review，都没有压力。

- **空间取舍：** 缺点是需要额外 `O(n)` 空间，但它换来的是更真实、更容易解释的状态；业务里通常可以接受。

### 解法 2：动态规划

- **核心思路：** `dp[i]` 表示以 `nums[i]` 结尾的最长递增子序列长度。

- **代码：** 只要某个位置的最长递增子序列长度达到 `3`，就说明找到了答案。

    ```go
    func increasingTriplet(nums []int) bool {
        n := len(nums)
        dp := make([]int, n)
        for i := range dp {
            dp[i] = 1
        }

        for i := 0; i < n; i++ {
            for j := 0; j < i; j++ {
                if nums[j] < nums[i] {
                    dp[i] = max(dp[i], dp[j]+1)
                    if dp[i] >= 3 {
                        return true
                    }
                }
            }
        }
        return false
    }
    ```

- **复杂度取舍：** 时间复杂度是 `O(n²)`，空间复杂度是 `O(n)`。

- **可扩展性：** 这是标准框架，极其通用。业务从“找 3 个”变成“找 5 个、7 个”，或者变成“找最长是多少”，都可以沿用这套逻辑；把 `>= 3` 改成 `>= k`，或者把目标长度做成配置项，整体逻辑不需要重写。面试时可以说：“如果这是线上业务，我会先用这个框架实现，把目标长度做成配置项。这样产品说从 3 个改成 5 个时，不需要改代码发版。”这展示的是你考虑通用业务问题，而不是只考虑一道题。

- **可维护性：** 这是标准 DP 模板，新人接手不用读大量注释。

- **业务可观测性：** 业务里更推荐这个，是因为 `dp[i]` 是真实状态。你可以打印出每个位置的 `dp` 值，清楚地看到用户在第几步达到了长度 2、在第几步达到了长度 3；如果最终没找到，也能看到用户最长走到了哪一步，比如只走到长度 2。相反，贪心法里的 `first` / `second` 是幽灵状态，你打印出来也不知道对应哪个真实位置，排查问题非常痛苦。

## 五、总结

- **判断标准：** DP 或预处理解法之所以好，是因为它们的状态是真实的：`dp[i]` 真的就是以第 `i` 个数结尾的序列长度，`leftMin[i]` 真的就是左边区间的最小值。你看到什么，它就是什么；状态越真实，代码越容易维护，也越容易迁移到新的业务需求。

- **最终结论：** 我要给出的算法，不只是能 AC 的算法，而是易于理解、能维护、能扩展、能迁移到业务场景里的算法。
