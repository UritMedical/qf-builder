#!/bin/bash
output_name="../dist/bin/qf.exe"

# 获取当前系统时间
build_Time=$(date +'%y%m%d')
# 删除当前文件夹下所有文件名为 counter+日期.txt 且日期不为当天的文件
find . -maxdepth 1 -type f -name "counter*" ! -name "counter${build_Time}.txt" -exec rm {} \;
# 存储计数器的文件路径
counter_file="BuildCounter${build_Time}.txt"
# 初始化计数器，默认为1
counter=1
# 如果计数器文件存在，从文件中读取当前计数器的值
if [ -f "$counter_file" ]; then
    counter=$(cat "$counter_file")
fi

# 格式化计数器为两位数字
formatted_counter=$(printf "%02d" "$counter")
# 增加计数器
((counter++))
##在使用git管理版本时
## 获取当前 Git 提交的版本信息
# version=$(git describe --tags --dirty --always)
#手工维护版本时
version="v1.01"

# 将新的计数器值写入文件
echo "$counter" > "$counter_file"

go build -ldflags="-w -s -X main.version=${version}.${build_Time}B${formatted_counter}" -o "$output_name"

if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
pause
fi
