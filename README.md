# cronparser

cronparser is a package to parse a cron string and output it to the command line.

## Installation

Requires golang >= 1.21

```
git clone https://github.com/warrenb95/cronparser
cd cronparser
go build
```

## Usage

```
go run . "*/15 0 1,15 * 1-5 /usr/bin/find"
```
or
```
go build
./cronparser"*/15 0 1,15 * 1-5 /usr/bin/find"
```

### Output

```
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
``` 

## License

[MIT](https://choosealicense.com/licenses/mit/)
