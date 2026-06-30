---
name: leetcode-solution-standard
description: 为 LeetCode 题目生成标准化的 problem.md、solution.go 和 solution_test.go，确保题面归档、解法实现和测试用例结构一致。
---

# leetcode-solution-standard

## 目录结构

- **推荐结构：** 每道 LeetCode 题使用一个独立目录，目录内固定放置题目、解法和测试三个文件。

    ```text
    leetcode/
      1071-greatest-common-divisor-of-strings/
        problem.md
        solution.go
        solution_test.go
    ```

- **`problem.md`：** 只放题目原文，包括题目描述、输入输出示例、约束条件。

    - 不写解法，保持“原题归档”的状态。

- **`solution.go`：** 放题解和代码实现，包括解题思路、关键观察、算法步骤、可运行代码、复杂度和易错点。

    - 解法逻辑写在函数上方，作为函数注释。
    - 多个解法放在同一个文件中，方便横向对比。
    - 代码负责表达“怎么做”，注释重点解释“为什么这么做”，避免重复描述代码表面动作。

- **`solution_test.go`：** 只放测试用例代码。

- **Q & A：**

    - Q：为什么题目描述和解法代码分开？
    - A：为了之后复习和重写。复习时可以先只看 `problem.md`，确认自己是否还能独立推出解法。

    - Q：为什么多个解法放在一个代码文件？
    - A：为了对比解法。相同题目的不同思路放在一起，更容易比较复杂度。

## problem.md 结构

- **推荐格式：**

    ```md
    # 题号. 题目标题
    
    ## Description
    
    - 题目描述第 1 段。
    
    - 题目描述第 2 段。
    
    - ...
    
    ## Examples
    
    ### Example 1
    
    - Input: ...
    
    - Output: ...
    
    ### Example 2
    
    - Input: ...
    
    - Output: ...
    
    ...
    
    ## Constraints
    
    - 约束 1
    
    - 约束 2
    
    - ...
    
    ```

- **整体风格：** 使用大纲笔记格式，每个独立语义段落都写成一个 `-` 列表项。

- **段落间距：** 每个列表项之间空一行，保持清晰的视觉分隔，方便阅读。

- **示例格式：** 每个示例使用 `### Example N` 分隔，`Input` 和 `Output` 各自作为独立列表项。

- **代码标记：** 变量名、函数名、表达式和约束条件使用反引号包裹，例如 `str1`、`str2`、`1 <= str1.length <= 1000`。
