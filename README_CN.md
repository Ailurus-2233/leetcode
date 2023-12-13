# LeetCode 代码仓库

欢迎来到我的 LeetCode 代码仓库！这个仓库专注于记录我对各种 LeetCode 问题的题解。随着我学习不同的编程语言，你将在这里找到 Java、C#、Python 和 Go 四种语言的代码实现。为了方便启动新项目，我还添加了一个方便的 shell 脚本，名为 `init_proj.sh`。

## 项目初始化

`init_proj.sh` 脚本允许你快速为特定的 LeetCode 问题在支持的任何语言中设置新项目。使用方法如下：

``` bash
./init_proj.sh /path/to/code language question_id
```

例如，运行以下命令：

``` bash
./init_proj.sh /题库/1.两数之和 Csharp 1
```

将创建一个 .NET Core 项目，结构如下：

- **Solution**：用于代码实现的控制台应用程序。
- **SolutionTest**：MSTest 项目，用于编写单元测试。

目前这个脚本只支持四种语言（Java,Python,C#,Go）,请随意根据自己的需求调整脚本，并轻松探索不同的编程语言！

## 仓库结构

- **题库**：包含我对一般 LeetCode 问题的题解。
- **剑指Offer**：专门用于“剑指Offer”问题的题解。

请注意，该仓库中的其他文件和文件夹主要支持 `init_proj.sh` 脚本，如果你只对浏览我的 LeetCode 题解感兴趣，可以忽略它们。

## 首选编辑器：Visual Studio Code

我个人使用 Visual Studio Code 进行编码任务。它是一款功能强大的编辑器，具有强大的插件系统。以下是我推荐的插件，特别适用于 LeetCode：

| languaue | plugins |
| --- | --- |
| C# | C#, C# dev kit, C# extensions |
| Java | Extension Pack for Java |
| Go | Go |
| Python | Python (shiro) |

这些插件提升了编码体验，并为单元测试代码提供了调试功能（Python 除外，你可以使用内置的 unittest 进行调试）。

## 测试用例示例

1. Java测试用例：[Q1.两数之和-Java](./题库/1.两数之和/Java/q1/SolutionTest.java)
2. Go测试用例：[Q1.两数之和-Go](./题库/1.两数之和/Go/main_test.go)
3. C#测试用例：[Q1.两数之和-C#](./题库/1.两数之和/Csharp/SolutionTest/UnitTest1.cs)
4. Python测试用例：[Q1.两数之和-Python](./题库/1.两数之和/Python/Solution.py)

祝编码愉快！ 🚀