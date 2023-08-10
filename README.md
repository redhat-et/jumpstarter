# jumpstarter
Jumpstarter agent and cmdline tools

## Usage

### command line tool

```
$ jumpstarter help
$ jumpstarter set-disk-image <id> <image file>
$ jumpstarter attach-storage <id>
$ jumpstarter detach-storage <id>

$ jumpstarter list-drivers
$ jumpstarter list-devices [-d driver]

$ jumpstarter console <id> [-d driver]
$ jumpstarter power-on <id> [-c] [--console] [--cycle] [--attach-storage]
$ jumpstarter power-off <id>

$ jumpstarter set-control <id> <signal> <status>

$ jumpstarter set-name <id> <name>
$ jumpstarter set-tags <id> tag1 [tag2] [tag3] ...

$ jumpstarter run <id> <jumpstarter-playbook.yaml>
```
