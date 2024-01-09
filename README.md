# cronparser

cronparser is a package to parse a cron string and output it to the command line e.g

## Installation

Requires golang >= 1.21

```
git clone https://github.com/warrenb95/cronparser
cd cronparser
go build
```

## Usage

```
$ ./cronparser "*/15 0 1,15 * 1-5 /usr/bin/find"
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```

## Testing

This package is unit tested with coverage of `88.4%`

## Contributing

Feel free to create pull requests. Ensure that you update/add unit tests.

## License

[MIT](https://choosealicense.com/licenses/mit/)
