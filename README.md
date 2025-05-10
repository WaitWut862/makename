# makename
Generates random names in string form.

Here is a demonstration of how makename can be used:

``` go

package main

import(
  "fmt"
  "github.com/WaitWut862/makename"
)

func main () {
  for i := 0; i < 100; i++ {
    name := makename.CompleteRand()
    fmt.Println(name)
    }
}
```
The functions can be cattegorized as such:
* A function starting with the word *Full* will include a first name, middle name, and sirname.
* A function staring with the words *FirstLast* will inlude a first name, and a sirname, while excluding a middle name.
* A function starting with the word *First* will only be a first name.
* A function staring with the word *Last* will only be a sirname.
* A function ending with the word *Neutral* will have a randomized gender.
* A function ending with either *Male* or *Female* will have the respective gender.
* If *Initial* follows a word in the function name, that part of the name will be initialized, e.g. Morris R may be produced by FirstLastInitialNeutral()
* Finaly, CompleteRand() will produce names of varied length, varied gender, and varied initialization.

Feel free to add more names to the lists.
