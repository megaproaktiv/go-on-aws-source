# Copy a single file to s3




## Unit test

```bash
go test 
```

## Integration test with AWS connection

### Configuration in `Taskfile.yml`

```yml
vars:
  BUCKET: dateneimer
  PREFIX: upload/
  FILE: testdata/test.txt
```

```bash
task run
```

## Benchmark vs AWS cli

```bash
 time ./dist/s3cp --file testdata/test.txt --bucket dateneimer --prefix upload/
Uploading testdata/test.txt to dateneimer/upload/
./dist/s3cp --file testdata/test.txt --bucket dateneimer --prefix upload/  0,01s user 0,02s system 15% cpu 0,196 total
```

CLI

```bash
aws s3 cp testdata/test.txt s3://dateneimer/upload/test.txt
upload: testdata/test.txt to s3://dateneimer/upload/test.txt
aws s3 cp testdata/test.txt s3://dateneimer/upload/test.txt  0,39s user 0,09s system 67% cpu 0,716 total
```