# Dates Range
Go module to get a list of dates between two dates


## API

```go
datesrange.New(start_date, end_date)
```


### start_date
**Type**: `string`   
Start date (e.g. `2020-01-01`)


### end_date
**Type**: `string`   
End date (e.g. `2020-03-07`). Can be empty if you need a range up to today.



### @return
**Type**: `[]time.Time, error`   



## Usage

```go
list, err := datesrange.New("2020-06-26", "2020-09-26")
if err != nil {
   fmt.Println("Error:", err)
}

for _, d := range list {
  fmt.Println(d)
}
```
