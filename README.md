# ipl

Generate a IP List from CIDR only if necessary.
Else, it prints the IP.

## Setup

```
go install github.com/shoxxdj/ipl@latest
```

## Usage

Using argv

```
ipl <CIDR>
```

Using stdin

```
echo "10.10.10.0/24" | ipl
```

