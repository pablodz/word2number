# word2number

Word 2 numbers.

"one" -> 1

# Install

```bash
go get -v github.com/pablodz/word2number
```

# Spanish

- Tested from 0 to 5000


```golang
text := "Mi código es cuatro mil novecientos noventa y nueve , y llueve cada cuatro mil novecientos noventa y ocho"
textFixed, err := word2number.Text2NumES(text)
if err != nil {
    fmt.Println(err)
}
fmt.Println("text:\t\t", text)
fmt.Println("textFixed:\t", textFixed)
```

```bash
text:            Mi código es cuatro mil novecientos noventa y nueve , y llueve cada cuatro mil novecientos noventa y ocho
textFixed:       Mi código es 4999 , y llueve cada 4998
```

# English

- Tested from 0 to 1000


```golang
text := "My code is one thousand"
textFixed, err := word2number.Text2NumEN(text)
if err != nil {
    fmt.Println(err)
}
fmt.Println("text:\t\t", text)
fmt.Println("textFixed:\t", textFixed)
```

```bash
text:            My code is one thousand
textFixed:       My code is 1000
```

# Portuguese


- Tested from 0 to 100


```golang
text := "My code is cinquenta e nove"
textFixed, err := lang.Text2NumPR(text)
if err != nil {
    fmt.Println(err)
}
fmt.Println("text:\t\t", text)
fmt.Println("textFixed:\t", textFixed)
```

```bash
text:            My code is cinquenta e nove
textFixed:       My code is 59
```
