# makename
Generates names.

## Usage

``` go
package main

import(
    "fmt"
    "github.com/WaitWut862/makename"
)

func main () {
    ng := makename.NewNameGenerator()

    // Random name generation
    for range 5 {
        fmt.Println(ng.Rand())
    }

    fmt.Println()

    // Custom name generation
    male := true
    fmt.Println(ng.Generate(
        makename.NameParam{
            Include: true,
        },
        makename.NameParam{
            Include: false,
        },
        makename.NameParam{
            Include: true,
            Initial: true,
        },
        male
    ))
}
```


> Output:
> Caden T
> Larry Bennett Foster
> Warren Gregory Ross
> Abigail A
> Ava Edith Martin
>
> Chase B

Feel free to add more names to the lists.
