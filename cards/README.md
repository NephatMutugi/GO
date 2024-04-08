# DATA STRUCTURES

> GO has 2 basic data structures for handling a list of records:
>1. Arrays: Fixed length list of things. 
>2. Slice: An array that can grow and shrink.
> 
> Arrays and slices must have an identical data type and only stores identical data types.
> 
## Slice

> 1. Slice declaration
> variableName:=[] <data-type> {}
> 2. Example
> cards := [] string {"Ace of spades", "Five of Diamonds"}
> 3. Add new item to Slice
> cards = append(cards, "Six of spades")
> 4. Note: The append() function does not modify the existing slice. Instead, it returns a new slice that we assign back to the variable of cards. 
> 
### Iterate through records

> for index, item := range items {
> fmt.Println(item)
> }
> 
#### var _int

> The blank identifier ('_') can be assigned to or declared with any value of any type. 
> Blank identifiers are discarded and not used by the compiler. 
> For example, you might consider using blank identifiers when you want to catch a return value but you have no intent to use this value afterward.