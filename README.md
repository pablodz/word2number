# word2number

Word 2 numbers.

# USE CASES

| INPUT                                                       | OUTPUT                                    |
| ----------------------------------------------------------- | ----------------------------------------- |
| siete mil pesos                                             | 7000 pesos                                |
| ochenta y nueve dólares                                     | 89 dólares                                |
| seis mil ochenta y cuatro euros                             | 6084 euros                                |
| cuatro mil novecientos noventa y nueve pesos                | 4999 pesos                                |
| cero guion cuarenta y cinco guion dos guion catorce         | 0-45-2-14                                 |
| ABCD tres siete cuatro ocho nueve                           | A B C D 3 7 4 8 9                         |
| cero uno uno cero uno                                       | 1011                                      |
| más cero setenta y seis                                     | 067                                       |
| Calle cuarenta y tres a la cuarenta y siete                 | Calle 43 a la 47                          |
| veinte de agosto                                            | 20 de agosto                              |
| cinco de julio                                              | 5 de julio                                |
| primero de febrero                                          | primero de febrero                        |
| primero del ocho del dos mil diez                           | primero del 8 del 1980                    |
| treinta y dos de febrero de mil novecientos ochenta         | 32 de febrero de 1980                     |
| 31 de diciembre del seis mil doce                           | 31 de dicieembre de 6012                  |
| menos 15                                                    | \-15                                      |
| menos 18 años                                               | \-18 años                                 |
| veinticinco                                                 | 25                                        |
| treinta positivo                                            | 30 positivo                               |
| treinta y cinco negativo                                    | 35 negativo                               |
| menos seis mil                                              | \-6000                                    |
| veinte años                                                 | 20 años                                   |
| cero                                                        | 0                                         |
| siete años tres meses ocho dias quince horas nueve segundos | 7 años 3 meses 8 dias 15 horas 9 segundos |
| mil veintiocho punto tres cuatro                            | 1028.34                                   |


# Install

```bash
go get -v github.com/pablodz/word2number
```

# Spanish

- Tested from 0 to 20000


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
