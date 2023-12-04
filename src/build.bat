@echo off

set output_name=..\dist\bin\qf.exe

:: 获取当前系统日期
for /f "delims=" %%a in ('wmic OS Get localdatetime ^| find "."') do set datetime=%%a
set build_Time=%datetime:~2,6%

:: 删除当前文件夹下所有文件名为 counter+日期.txt 且日期不为当天的文件
for %%f in (counter*) do (
    set filename=%%~nf
    if /i not "!filename:~7,6!"=="%build_Time%" del "%%f"
)

:: 存储计数器的文件路径
set counter_file=BuildCounter%build_Time%.txt

:: 初始化计数器，默认为1
set counter=1

:: 如果计数器文件存在，从文件中读取当前计数器的值
if exist %counter_file% (
    set /p counter=<%counter_file%
    rem 删除末尾的换行符
    set counter=%counter:~0,-1%
)

:: 格式化计数器为两位数字
set "formatted_counter=0%counter%"
set "formatted_counter=%formatted_counter:~-2%"

:: 增加计数器
set /a counter+=1

:: 在使用 Git 管理版本时

:: 获取当前 Git 提交的版本信息

:: for /f %%v in ('git describe --tags --dirty --always') do set version=%%v

:: 手工维护版本时
set version=v1.01

:: 将新的计数器值写入文件
echo %counter% > %counter_file%

:: 生成可执行文件
go build -ldflags="-w -s -X main.version=%version%.%build_Time%B%formatted_counter%" -o "%output_name%"

:: 暂停以保持窗口打开
::pause
