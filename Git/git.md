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

//修改.bash_profile或者在受支持的命令行运行
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

`git restore`把文件从缓存区撤销，回到未被追踪的状态

### 分支操作

创建分支

`git branch 分支名`

`git checkout -b 分支名`

`git branch [-a,-r]`查看所有分支,远程分支

 `git checkout dev`切换分支

第一种改名方法

`git checkout -b 新分支 原分支`

`git branch -d 原分支`//删除原分支

第二种改名方法

`git branch -m 分支名`为当前分支改名

`git branch -m 原分支名 新分支名`为指定分支改名

`git branch -d 分支名`删除分支

查看merged情况

`git branch --merged`//合并过的分支

`git branch --no-merged`//未合并的分支

`git merge 分支名`分支合并(当前分支与其他分支合并)

`git stash`将所有未提交的修改（包括暂存的和非暂存的）都保存起来，用于后续恢复当前工作目录。

`git stash apply stash@{1}`恢复某个暂时保存的工作

`git stash pop`恢复最近一次缓存的工作

`git stash list`查看所有的stash

`git stash drop [stash@{0}]`移除stash

`git stash clear`删除所有的stash

`git archive 分支名 --prefix='hngy/' --forma=zip > hngy.zip`将文件打包配置路径和格式

`git rebase master[主分支]`主分支已经改变，无法合并分支，需要先同步主分支内容然后在合并

### 远程操作

`git clone -o 主机名 url`克隆时指定主机名

`git remote add 主机名 url`添加远程主机

`git remote rm 主机名`删除远程主机

`git remote renmae 原主机名 新主机名`修改主机名

`git remote [-v]`列出所有远程主机[查看主机网址]

`git remote show 主机名`查看该主机详细信息

`git fetch 主机名 [分支名]`远程主机更新，全部[分支]拉取回本地

### 拉取和推送

`git pull [--rebase] 远程主机名 远程分支名:本地分支名`取回远程主机某个分支的更新，再与本地的指定分支合并。等同于先 git fetch 再 git merge。

注:与当前分支合并冒号内容可以省略,已经与当前分支存在追踪关系，只需写明远程主机名即可`git pull origin`，`git pull`当前分支自动与唯一一个追踪分支进行合并(慎用)。

`git branch --set-upstream 本地分支 主机名/远程分支`手动建立追踪关系

`git pull -p`如果远程主机删除某个分支，在`git pull`时并不会删除本地的分支，此命令可以同步删除分支

`git push [-f,--force,--force-with-lease] 远程主机名 本地分支名:远程分支名`用法同上，将分支推送到远程主机上，如果远程没有该分支将会自动创建。

注：<mark>省略本地分支名将会删除远程分支,force命令极不安全</mark>,等同于`git push 远程主机名 --delete 分支名`

`git push -u origin master`如果当前分支与多个主机存在追踪关系，可以使用该命令指定一个默认主机，以便`git push`推送当前分支致远程主机。

### 实例

1. gitignore文件配置失效,在配置之前部分文件已经被git仓库进行追踪，需要删除本地的缓存，进行重新提交。`git rm -r --cached .或填写路径`然后在添加到暂存区进行重新提交。