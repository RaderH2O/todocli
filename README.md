# todocli, a ✔️ todo app written in Go 🔷

## How to use

todocli is very minimalistic. When you start the application, you will be faced with some options
adding todos, toggling todos, deleting todos, and quitting the application

### Adding
To add a task, type `a` and press enter, then you will be prompted to enter the content of your task

### Toggling
Toggling is simply "doing" or "undoing" the task. You will need to specify the task number you want to toggle right after typing `t`. So if you want to toggle a task with the id of `3`, you would write `t3` and then enter

### Deleting
Deleting tasks is similar to toggling, where you have to specificy what item you are going to delete. So for example, `d12` will delete an item with id of `12`

### Quitting
Quitting is very simple. You just type `q`, enter, and now are not running your application!

## How to run

Prerequisite:
- A Go compiler

1. Run :
```shell
go run .
```
2. Enjoy!

## How to build

Prerequisite:
- A Go compiler

1. Run :
```shell
go build .
```
2. There should be an executable file in the same directory named `todocli`!