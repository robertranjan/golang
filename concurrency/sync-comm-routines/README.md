# syncronouse communication between go-routines

```mermaid
flowchart LR
    *[main go-routine] --> B("2. square")
    subgraph "Square Ints"
        B --> r2(go-routine:square Ints)
    end
    *[main go-routine] --> A("1. toChannel")
    subgraph "Int to Channel"
        A --> r1(go-routine:convert to channel)
    end
    * --> D("3. print result")
```

* Channel created by toChannel() passed to go-routine:squareInts through square()
* go-routine:squareInts reads each element from channel and square it and drop to another channel(ch2)
* Main go-routine reads from channel(c2) from square and write to stdout
