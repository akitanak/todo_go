# TODO

This is TODO CLI App written by Go.


## Create TODO

Create a TODO.

### Usage

```bash
$ todo add [todo description] [flags]
```

### Flags

|flag|description|
|:---|:----------|
|-d  |Due date.  |

```bash
$ todo add "build todo app with golang." -d 2022/01/31
```

## List TODO

List TODOs.

### Usage

```bash
$ todo list [flags]
No. Todo                         Due Date
 1  build todo app with golang.  2022/01/31
 2  buy a Macbook Pro.
 3  call Dental clinic.          2022/01/22
```

### Flags

|flag|description|
|:---|:----------|
|-a  |List all TODOs include finished TODO. |

