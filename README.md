# LeetCode Code Repository

[仓库介绍-中文](./README_CN.md)

Welcome to my LeetCode code repository! This repository is dedicated to documenting my solutions to various LeetCode problems. As I explore different programming languages, you'll find code implementations in Java, C#, Python, and Go. To make the process of starting a new project for a LeetCode problem easier, I've also included a handy shell script named `init_proj.sh`.

## Project Initialization

The init_proj.sh script allows you to quickly set up a new project for a specific LeetCode question in any of the supported languages. Here's an example of how to use it:
 
``` bash
./init_proj.sh /path/to/code language question_id
```
For instance, running the following command:

``` bash
Copy code
./init_proj.sh /题库/1.两数之和 Csharp 1
```

will create a .NET Core project with the following structure:

- **Solution**: Console application for code implementation.
- **SolutionTest**: MSTest project for writing unit tests.

Feel free to adapt the script to your needs and explore different languages effortlessly!

## Repository Structure
- **题库**: Folder containing my solutions to general LeetCode problems.
- **剑指Offer**: Folder dedicated to solutions for problems found in "剑指Offer" (Sword Offer).

Please note that other files and folders in this repository primarily support the init_proj.sh script and can be ignored if you're solely interested in browsing my LeetCode solutions.

## Preferred Editor: Visual Studio Code

I personally use Visual Studio Code for my coding tasks. It's an excellent editor with a robust plugin system. Here are some plugins I recommend, especially for LeetCode:

| languaue | plugins |
| --- | --- |
| C# | C#, C# dev kit, C# extensions|
| Java | Extension Pack for Java|
| Go | Go |
| Python | Python (shiro) |

These plugins enhance the coding experience and provide debugging capabilities for unit test code (excluding Python, where you can use the built-in unittest for debugging).

Happy coding! 🚀