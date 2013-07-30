//   Copyright 2013 Matt Ho
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"flag"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	gos3 "launchpad.net/goamz/s3"
	"log"
	"mime"
	"os"
	"path"
	"strings"
)

var (
	bucketName = os.Getenv("S3_BUCKET")
	regionName = os.Getenv("S3_REGION")
)

func bucket() *gos3.Bucket {
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}

	// verify bucket
	if bucketName == "" {
		log.Fatalf("unable to find bucket, %s\n", bucketName)
	}

	// verify region
	if regionName == "" {
		regionName = "us-west-2"
	}
	if _, ok := aws.Regions[regionName]; !ok {
		log.Fatalf("unable to find region, %s\n", regionName)
	}

	// obtain a reference
	s3 := gos3.New(auth, aws.Regions[regionName])
	return s3.Bucket(bucketName)
}

func contentType(filename string) string {
	parts := strings.Split(filename, ".")
	suffix := "." + parts[len(parts)-1]
	if mimeType := mime.TypeByExtension(suffix); mimeType == "" {
		return "application/octet-stream"
	} else {
		return mimeType
	}
}

func handleDownload(filename string) {
	data, err := bucket().Get(filename)
	if err != nil {
		panic(err)
	}

	basename := path.Base(filename)
	ioutil.WriteFile(basename, data, 0644)
}

func uploadFile(filename string, target string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	// convert the name of the targetfile to what we expect
	key := target[len("s3:"):len(target)]
	if key == "" || key == "." {
		key = path.Base(filename)
	}

	// determine what mime type the original document is
	mimeType := contentType(filename)

	// upload the file
	log.Printf("uploading %s to s3:%s (%s)\n", filename, key, mimeType)
	bucket().Put(key, data, mimeType, gos3.Private)
}

func uploadFiles(sources []string, target string) {
	for _, source := range sources {
		if strings.HasPrefix(source, "s3:") {
			log.Fatalf("invalid arguments, %s -- s3: can either be part of the source or part of the targets\n", source)
		}

		uploadFile(source, target)
	}
}

func downloadFile(source string, target string) {
	key := source[len("s3:"):len(source)]
	data, err := bucket().Get(key)
	if err != nil {
		log.Fatalln(err)
	}

	filename := target
	if target == "." {
		filename = path.Base(key)
	}
	ioutil.WriteFile(filename, data, 0644)
}

func downloadFiles(sources []string, target string) {
	for _, source := range sources {
		if !strings.HasPrefix(source, "s3:") {
			log.Fatalf("invalid arguments, %s -- s3: can either be part of the source or part of the targets\n", source)
		}

		downloadFile(source, target)
	}
}

func copyFiles(args []string) {
	if len(args) < 2 {
		log.Fatalln("s3cp requires at least 2 arguments; one or more sources and a target")
	}

	target := args[len(args)-1]
	sources := args[0 : len(args)-1]
	if strings.HasPrefix(target, "s3:") {
		uploadFiles(sources, target)
	} else {
		downloadFiles(sources, target)
	}
}

func main() {
	flag.Parse()

	copyFiles(flag.Args())
}
