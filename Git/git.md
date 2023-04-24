# git学习

### 文件配置

git初始化

`git init`

查看git配置列表

`git config --list`

配置git用户名和邮箱(去掉global配置当前项目)

`git config --global user.name '名字'`

`git config --global user.email '邮箱'`

创建忽略配置文件(.gitignore)

```textile
*.txt  //忽略所有txt文件
!a.txt //不忽略该文件
文件夹 //忽略该文件夹内所有内容
文件夹/**/*.php //忽略该文件夹下的php文件
```

全局alias命令别名

```textile
git config --global alias.a add
git config --global alias.c commit
git config --global alias.o checkout

//也可以修改.gitconfig文件末尾添加
[alias]
a = add
c = commit
o = checkout
# ...

//或者在受支持的命令行
alias gs='git status'
alias ga='git add'
alias gb='git branch'
alias gc='git commit'
```





### 版本控制

拉取版本

`git clone -m url`

添加项目致暂存区

`git add 文件名`//单个文件或目录

`git add .`//所有文件和目录

查看文件状态

`git status`

提交到版本库

`git comit -m '说明'`

删除工作区文件或文件夹(删除放置暂存区)

`git rm '文件'`

`git rm -rf '文件夹'`

删除工作区和暂存区的文件

`git rm -f '文件'`

只删除暂存区文件，保留工作区文件

`git rm --cached '文件'`

### 日志

`git log`查看所有提交日志信息

`git log -n`限定查看提交日志数量

`git log --oneline`简化log输出

`git log -p`输出每个提交的具体内容

`git show`只显示一个提交的具体内容

`git shortlog`输出汇总消息

`git shortlog -s`统计每个作者的提交数量

`git shortlog -n`对统计量进行倒叙排序

`git log --author="作者"`过滤提交

`git log --after '日期' --before '日期'`限定日期范围

`git log --name-only`显示提交的文件名

`git log --name-status`显示提交的文件状态



### 撤销和修改

修改提交致版本库的信息(包括说明和文件内容)

`git comit --amend`撤销上一次提交  将暂存区文件重新提交

`git checkout --文件名`   拉取暂存区文件 并将其替换成工作区文件

`get reset HEAD --文件名`拉取最近一次提交版本库的文件到暂存区，操作不影响工作区



### 分支操作

`git branch 分支名`创建分支

`git branch [-a]`查看分支,所有分支

 `git checkout dev`切换分支