s3cp
==

Utility to easily upload / download content from S3.

This tool is intended to be used as part of a non-interactive install or deployment of software.

## Environment Variables

The environment variables required for s3cp to operate.

|Name|Required|Description|
|:--|:--:|:---|
|AWS_ACCESS_KEY_ID|yes|the aws credentials|
|AWS_SECRET_ACCESS_KEY|yes|the aws credentials|
|S3_BUCKET|yes|the s3 bucket to read/write to|
|S3_REGION|no|the region where the bucket is located.  defaults to us-west-2|

##### Copy single file to S3

Copy the file, sample.txt, to S3 using the same name.

```
s3cp sample.txt s3:
```

##### Copy single file to S3 with different name

Copy argle.txt to S3, renaming it to bargle.txt in the process.

```
s3cp argle.txt s3:bargle.txt
```

##### Copy multiple files S3

Copies file1, file2, and file3 to S3.

```
s3cp file1 file2 file3 s3:
```

##### Copy single file from S3

Copies file, sample.txt, from S3 to current directory

```
s3cp s3:sample.txt .
```

##### Copies multiple files from S3

Copies file1, file2, file3 from S3 to current directory

```
s3cp s3:file1 s3:file2 s3:file3 .
```
