#!/bin/bash

# 记录当前目录
current_path=$(pwd)

# 获取输入参数
folder_path=$1
programming_language=$2
question_id=$3

# 检测输入参数是否合理
if [ -z "$folder_path" ] || [ -z "$programming_language" ] || [ -z "$question_id" ] ; then
    echo "Usage: $0 <folder_path> <programming_language> <question_id>"
    exit 1
fi

# 检测语言是否合法
case $programming_language in
    "Go" | "Java" | "Python" | "Csharp")
    ;;
    *)
        echo "Invalid programming language. Please choose Go, Java, Python, or Csharp."
        exit 1
    ;;
esac

# 检测目录是否存在，如果不存在则创建
if [ ! -d "$folder_path" ]; then
    mkdir -p "$folder_path"
fi

# 检测是否存在 question.md 文件，如果不存在则创建
question_file="$folder_path/question.md"
if [ ! -f "$question_file" ]; then
    touch "$question_file"
    echo "# Questions" >> "$question_file"
    echo "Please add your questions here." >> "$question_file"
    echo "" >> "$question_file"
fi

init_flag="n"

# 检测是否存在对应语言的文件夹
language_folder="$folder_path/$programming_language"
if [ -d "$language_folder" ]; then
    read -p "The folder for $programming_language already exists. Do you want to recreate it? (y/n): " recreate
    if [ "$recreate" = "y" ]; then
        rm -rf "$language_folder"
        mkdir "$language_folder"
        echo "Folder recreated."
        init_flag="y"
    else
        echo "Script ended."
    fi
else
    mkdir "$language_folder"
    init_flag="y"
    echo "Folder for $programming_language created."
fi

if [ "$init_flag" = "y" ]; then
    # 执行特定于语言的初始化步骤
    case $programming_language in
        "Go")
            cd "$language_folder"
            go mod init solution"$question_id"
            touch main.go
            echo "package main" >> main.go
            cd "$current_path"
            go work use "$language_folder"
        ;;
        "Java")
            mkdir "$language_folder/q$question_id"
            cp Solution.java "$language_folder/q$question_id"
            cp SolutionTest.java "$language_folder/q$question_id"
            cd "$language_folder/q$question_id"
            # 在文件头部添加 package q$question_id 语句
            sed -i "1i package q$question_id;" Solution.java
            sed -i "1i package q$question_id;" SolutionTest.java
        ;;
        "Python")
            cp solution.py "$language_folder"
            cp solution_test.py "$language_folder"
        ;;
        "Csharp")
            cd "$language_folder"
            dotnet new sln -n Csharp
            dotnet new console -n Solution
            dotnet new mstest -n SolutionTest
            dotnet add SolutionTest/SolutionTest.csproj reference Solution/Solution.csproj
            dotnet sln add Solution/Solution.csproj
            dotnet sln add SolutionTest/SolutionTest.csproj
        ;;
    esac
fi

echo "Initialization for $programming_language completed."