# goploader

upload files with golang to http server

test your changes by running Makefile:

```
make
```

### build

```
cd cli && go build -o goploader
```

### run

```
./goploader --file=example/files/testfile.json --url=http://localhost:8080/upload --method=post
```
