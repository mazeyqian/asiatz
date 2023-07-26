# asiatz

Asiatz is a Golang library for converting timezones and other utilities. It provides a function to convert Shanghai time to UTC time.

## Installation

To use the asiatz library in your Golang project, you can import it using the following command:

```
go get github.com/mazeyqian/asiatz
```

## Usage

To use the `ConvertShanghaiToUTC` function, first import the `asiatz` package:

```go
import "github.com/mazeyqian/asiatz"
```

Then, call the `ConvertShanghaiToUTC` function with a Shanghai time string:

```go
utcTime, err := asiatz.ConvertShanghaiToUTC("08:00")
if err != nil {
    // handle error
}
fmt.Println(utcTime) // Output: 00:00
```

## Contributing

If you find a bug or have an idea for a new feature, please open an issue on the GitHub repository. Pull requests are also welcome!

## License

This library is licensed under the MIT license. See the LICENSE file for more details.
