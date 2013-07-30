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
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestContentType(t *testing.T) {
	filename := "sample.txt"
	expected := "text/plain; charset=utf-8"
	if contentType(filename) != expected {
		t.Fatalf("expected content type to be %s; actual was %s\n", expected, contentType(filename))
	}
}

func TestContentTypeForUnknownExtension(t *testing.T) {
	filename := "sample.arglebargle"
	expected := "application/octet-stream"
	if contentType(filename) != expected {
		t.Fatalf("expected content type to be %s; actual was %s\n", expected, contentType(filename))
	}
}

func TestUploadAndDownload(t *testing.T) {
	source := fmt.Sprintf("sample-%d.txt", time.Now().Unix())
	target := "the-" + source

	ioutil.WriteFile(source, []byte(source), 0644)
	uploadFiles([]string{source}, "s3:")

	// remove the file
	os.Remove(source)
	downloadFiles([]string{"s3:" + source}, target)
	data, err := ioutil.ReadFile(target)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != source {
		t.Fatalf("expected data returned to be %s; actual was %s\n", source, string(data))
	}
	os.Remove(target)
}
