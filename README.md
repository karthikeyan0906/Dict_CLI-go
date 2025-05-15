# Dict_CLI-go
Dictionary CLI written in Go

This is a simple CLI dictionary tool written in Go that fetches word definitions from a free dictionary API. It processes a single word entered by the user and prints its definition, part of speech, and pronunciation.

Below packages are used 

net/http - for making API requests
encoding/json - parsing unstructured json 
os.Args - to work with CLI args passed by user
io - to read response body

**Prerequisites**
 Install Go on your machine

** Installation **
  git clone https://github.com/karthikeyan0906/Dict_CLI-python.git 

**Run the script**
    go run dict.go args     
Example 
    go run dict.go pen
Output
     go run dict.go pen
    pen (noun): An enclosure (enclosed area) used to contain domesticated animals, especially sheep or cattle.
    Pronunciation: /pɛn/
if the word doesn't exist it will return "No definitions found"
