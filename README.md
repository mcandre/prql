# prql ("Prequel"): a command line SQL syntax checker

# EXAMPLE

```console
$ prql examples/
examples/empty-statement.sql: syntax error at position 1
```

# DOWNLOAD

https://github.com/mcandre/prql/releases

# API DOCUMENTATION

https://godoc.org/github.com/mcandre/prql

# MAJOR FEATURES

* Validates SQL syntax, such as common PostgreSQL and MySQL statements.
* Scans multi-statement .SQL scripts.
* Recurses along large folder trees.
* No dependency on live SQL servers or clients.
* Ops-friendly exit code for CI, script chaining.

# COMPLEX SCRIPT MATCHING

```console
$ find examples -type f -name '*.sql' -print0 |
    while IFS= read -r -d '' f; do
        prql "$f" || exit 1
    done
examples/empty-statement.sql: syntax error at position 1
```

# RUNTIME REQUIREMENTS

* N/A

# CONTRIBUTING

See [DEVELOPMENT.md](DEVELOPMENT.md).

# LICENSE

FreeBSD

# CREDITS

* [xwb1989/sqlparser](https://github.com/xwb1989/sqlparser) does the heavy lifting.
