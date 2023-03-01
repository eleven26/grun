# Grun

记录常用命令的命令行小工具。（`macOS`/`Linux` 可用）。相比简单的 `alias`，它可以记录命令的描述，方便查看。

## 安装

```bash
go install github.com/eleven26/grun@v0.0.1
```

## 支持的命令

```shell
grun -h
```

输出：

``` 
使用 "-h" 参数查看所有子命令

Usage:
  grun [command]

Available Commands:
  add         新建命令
  completion  Generate the autocompletion script for the specified shell
  delete      删除命令
  help        Help about any command
  list        列出所有命令
  run         运行命令
  update      更新命令

Flags:
  -c, --command string       命令
  -d, --description string   描述
  -h, --help                 help for grun
  -i, --id string            id
  -n, --name string          名称

Use "grun [command] --help" for more information about a command.
```

## 示例

### 列出所有命令

```shell
grun list
```

输出：

```
+----+------+---------+--------------+
| ID | 名称 |  命令   |     描述     |
+----+------+---------+--------------+
| 1  |  ls  | ls /var | test command |
+----+------+---------+--------------+
```

### 运行命令

```
# 最后一个参数是 id
grun run 1
```

### 添加命令

```shell
grun store -n "ls" -c "ls /var" -d "test command"
```

### 删除命令

```shell
# 最后一个参数是 id
grun delete 1
```

### 更新命令

```shell
grun update 1 -c "ls /var/log"
```
