# asiatz

**asiatz** is a Go package for converting timezones and other utilities. It provides functions to convert various time zones to UTC time.

## Installation

To use the asiatz package in your Go project, you can import it using the following command:

```
go get github.com/mazeyqian/asiatz
```

## Usage

To use the timezone conversion functions, first import the `asiatz` package:

```go
import "github.com/mazeyqian/asiatz"
```

Then, call the appropriate function with a time string in the format "HH:mm":

```go
utcTime, err := asiatz.ShanghaiToUTC("08:00")
if err != nil {
    // handle error
}
fmt.Println(utcTime) // Output: 00:00
```

Here is a list of supported time zones and their corresponding functions:

| Time Zone   | Function Name  |
|-------------|----------------|
| Shanghai    | ShanghaiToUTC  |
| Tokyo       | TokyoToUTC     |
| Hong Kong   | HongKongToUTC  |
| Singapore   | SingaporeToUTC |
| Taipei      | TaipeiToUTC    |
| Seoul       | SeoulToUTC     |
| Beijing     | BeijingToUTC   |
| Dubai       | DubaiToUTC     |
| Delhi       | DelhiToUTC     |
| Jakarta     | JakartaToUTC   |
| Bangkok     | BangkokToUTC   |

The `ToUTC` function is a helper function used by the timezone conversion functions in the `time_utils.go` file. It takes two arguments: `offset` and `timeString`. 

The `offset` argument specifies the UTC offset for the given time zone. For example, the UTC offset for Shanghai is +8, so you would pass 8 as the `offset` argument when converting a Shanghai time string to UTC.

The `timeString` argument is a string in the format "HH:mm" that represents the time in the given time zone.

The `ToUTC` function returns a string in the same format ("HH:mm") that represents the equivalent time in UTC.

Here's an example usage of the `ToUTC` function:

```go
utcTime, err := ToUTC(8, "08:00")
if err != nil {
    // handle error
}
fmt.Println(utcTime) // Output: 00:00
```

In this example, we're converting a Shanghai time string ("08:00") to its equivalent UTC time string ("00:00"). We pass 8 as the `offset` argument since Shanghai's UTC offset is +8.

## Contributing

If you find a bug or have an idea for a new feature, please open an issue on the GitHub repository. Pull requests are also welcome!

## License

This package is licensed under the MIT license. See the LICENSE file for more details.
