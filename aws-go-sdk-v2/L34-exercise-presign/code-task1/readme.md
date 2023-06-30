
## Usage

Replace `dateneimer` with the name of your bucket.

```bash
go run main/main.go --file testdata/text.txt --bucket dateneimer
```


## CodeWhisperer

For best results, follow these practices.

Give CodeWhisperer something to work with. The more code your file contains, the more context CodeWhisperer has for generating recommendations.
Write descriptive comments. “Function to upload a file to S3” will get better results than “Upload a file”.
Specify the libraries you prefer by using import statements.
Use descriptive names for variable and functions. A function called “upload_file_to_S3” will get better results than a function called “file_upload”.
Break down complex tasks into simpler tasks.