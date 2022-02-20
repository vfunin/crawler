# crawler
[![codecov](https://codecov.io/gh/vfunin/crawler/branch/main/graph/badge.svg?token=BA952ZCZ94)](https://codecov.io/gh/vfunin/crawler)

Crawler is a program for find titles in URLs.

Examples of using:
```shell
./bin/crawler -u https://ya.ru # Finds all links and title up to the second level of nesting and outputs to the console
./bin/crawler -u https://ya.ru -o result.csv # Same but output to file
./bin/crawler -u https://ya.ru -l # Use for change log level (string debug/info/error etc)
./bin/crawler -u https://ya.ru -p # Fires panic and recover in first link
```
