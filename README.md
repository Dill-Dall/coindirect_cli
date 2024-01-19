# coindirect_cli
---
Simple cmdline go app using cobra.  
Can be download and installed as binaries via github releases. current v0.0.1

For local development.
```bash
go mod tidy
go run main.go -h
```

The app uses the following  sandbox api "https://api.sandbox.coindirect.com/api/" to answer the following spec.

Solves the following spec:
- A list of all countries
  - With the following attributes: 
    - Name 
    - Currency 
    - Required documentation 
    - Max withdrawal 
  - Sorting (Ascending - Descending) on Name and Currency 
- A list of all currencies, with information in which countries each currency is used


Can be interacted with via cmd functions. .exe is not signed so will trigger antivirus alert.

```bash
./cli-linux -h
A cli for retrieving data from the .coindirect.com api

Usage:
  cmd [flags]

Flags:
  -c, --currencymap      return a map of which currencies are used by which lands
  -d, --descending       descending sort by sortkey, default is ascending order
  -h, --help             help for cmd
  -s, --sortkey string   id|name|currency (default "id")
```
