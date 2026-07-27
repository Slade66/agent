---
name: leetcode-solution-standard
description: 为 LeetCode 题目生成标准化的 problem.md、solution.md、solution_template.go 和 solution_test.go，确保题面归档、解法笔记、代码骨架和测试用例结构一致。
---

# leetcode-solution-standard

## SOP

1. 接收用户提供的 LeetCode 题目文本。
2. 确认输出位置：如果用户指定目录，就在指定目录下创建题目文件夹；否则在当前目录下创建。
3. 题目文件夹命名为 `题号-题目标题`，例如 `1071-greatest-common-divisor-of-strings`。
4. 在题目文件夹内创建 `problem.md`、`solution.md`、`solution_template.go` 和 `solution_test.go` 四个文件。
5. 按下方规范分别写入题面归档、解法笔记骨架、Go 函数骨架和 table-driven tests，并在完成前对 Go 文件执行 `gofmt`。

## 目录结构

- **推荐结构：** 每道 LeetCode 题使用一个独立目录，目录内固定放置以下四个文件。

    ```text
    leetcode/
      1071-greatest-common-divisor-of-strings/
        problem.md
        solution.md
        solution_template.go
        solution_test.go
    ```

- **`problem.md`：** 只放题目原文，包括题目描述、输入输出示例、约束条件。不写解法，保持“原题归档”的状态。

- **`solution.md`：** 放解题笔记。新建时只保留“题意理解”和“方法一”两个二级标题。

- **`solution_template.go`：** 只放空函数，不写具体实现；新增解法时复制为 `solution1.go`、`solution2.go` 等。

- **`solution_test.go`：** 只放测试用例代码。

- **Q & A：**

    - Q：为什么题目文件夹的名字是这样？
    - A：用 slug 命名，例如 `1071-greatest-common-divisor-of-strings`，以避免空格和点带来的命令行路径问题。

    - Q：为什么题目、解题笔记和代码分开？
    - A：为了之后复习和重写。先看 `problem.md` 独立思考，再用 `solution.md` 记录思路，最后在独立的 Go 文件中实现不同解法。

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
    
    - **Input:** ...
    
    - **Output:** ...

    - **Explanation:** ...
    
    ### Example 2
    
    - **Input:** ...
    
    - **Output:** ...

    - **Explanation:** ...
    
    ...
    
    ## Constraints
    
    - 约束 1
    
    - 约束 2
    
    - ...
    
    ```

- **整体风格：** 使用大纲笔记格式，每个独立语义段落都写成一个 `-` 列表项。

- **段落间距：** 每个列表项之间空一行，保持清晰的视觉分隔，方便阅读。

- **示例格式：** 每个示例使用 `### Example N` 分隔；`Input`、`Output` 和 `Explanation` 各自作为独立列表项，并用 Markdown 双星号完整加粗标签及冒号，固定写成 `**Input:**`、`**Output:**` 和 `**Explanation:**`。

- **代码标记：** 变量名、函数名、表达式和约束条件使用反引号包裹，例如 `str1`、`str2`、`1 <= str1.length <= 1000`。

## solution.md 结构

- **初始内容：** 新建时只写下面两个 H2 标题，不补充正文。

- **空行规则：** `## 题意理解` 上方保留一个空行，两个标题之间保留三个空行，`## 方法一` 下方保留四个空行。

    ```md

    ## 题意理解



    ## 方法一




    ```

## solution_template.go 结构

- **实现边界：** 只创建可编译的函数骨架，不实现具体解题逻辑，也不添加解法注释，保留给用户自己完成。

- **推荐格式：**

    ```go
    package main

    func gcdOfStrings(str1 string, str2 string) string {
        return ""
    }
    ```

- **格式要求：** 提交前执行 `gofmt`。

## solution_test.go 结构

- **生成原则：** 按照用户提供的题目文本创建测试用例，先完整覆盖题目中的官方示例，再补充边界条件和易错反例。

- **测试风格：** 使用 table-driven tests，统一组织输入、期望输出和用例名称。

- **推荐格式：**

    ```go
    package main
    
    import "testing"
    
    func TestGcdOfStrings(t *testing.T) {
        testCases := []struct {
            name string
            str1 string
            str2 string
            want string
        }{
            {
                name: "official example 1",
                str1: "ABCABC",
                str2: "ABC",
                want: "ABC",
            },
        }
    
        for _, tc := range testCases {
            t.Run(tc.name, func(t *testing.T) {
                got := gcdOfStrings(tc.str1, tc.str2)
                if got != tc.want {
                    t.Errorf("gcdOfStrings(%q, %q) = %q, want %q", tc.str1, tc.str2, got, tc.want)
                }
            })
        }
    }
    ```

## 参考示例

- ./references/1071. Greatest Common Divisor of Strings/
