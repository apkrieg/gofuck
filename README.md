# GoFuck - An AOT AIO Brainfuck Toolchain in Go

GoFuck is a Brainfuck interpreter, compiler, and runtime with Ahead-of-Time (*AOT*) compilation and debugging support.

`go get github.com/apkrieg/gofuck`

## Spec

### Language

| Mnemonic | Description |
| -------- | ----------- |
| > | ++Pointer |
| < | --Pointer |
| + | ++*Pointer |
| - | --*Pointer |
| . | print(*Pointer) |
| , | *Pointer = scan() |
| [ | while(*ptr) { |
| ] | } |

## Usage

### Interpret

```
gofuck
>>>: ++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
Hello World!
>>>:
```

### Compile

```
gofuck build hello_world.bf
```

### Decompile (Coming Soon!)
```
gofuck unbuild hello_world.bfc
```

### Run

```
gofuck run hello_world.bf
```

### Debug

Debug mode will start the runtime in debug mode.

| Event | debug | run |
| ----- | ----- | --- |
| A `d` instruction is encountered | debug console is launched | no op |
| Data index out of bounds | debug console is launched | error |

Start in debugging mode:
```
gofuck debug hello_world.bf
```

Resume normal execution
```
dbg: resume
```

Dump everything (Coming Soon!)
```
dbg: dump dump_file.dmp
```

Set cell #0 to 100
```
dbg: set_cell 0 100
```

Set pointer to 2
```
dbg: set_pointer 2
```

Print cell #3
```
dbg: get_cell 3
```

Print pointer
```
dbg: get_pointer
```

## Future Plans

### Cleanup/Optimization

I wrote most of the initial release at 3:00 AM because I couldn't sleep. There are probably many improvements that could happen, but I'll get to that some other day. The important thing is, it works.

### Threading

This is still in early planning, but there would be two new instructions, `(` and `)`. Any code inside of `()` would start on a new thread and receive its own data pointer initialized to the calling threads data pointer.

### Compiling to Native

Compiling to native executables should be implemented within six months or sooner if my schedule allows.

## Authors
- Andrew Krieg

## License
- MIT License
