package main

import (
	"testing"
)

func BenchmarkListFiles(b *testing.B) {
	b.ReportAllocs()
	list_all_files("/Users/carlos/Downloads/enron_mail_20110402/maildir")
}

func BenchmarkProgramWith1000LimitMap(b *testing.B) {
	b.ReportAllocs()
	index_mails("/Users/carlos/Downloads/enron_mail_20110402/maildir", 1000)
}

// run test benchmark
// https://medium.com/@felipedutratine/profile-your-benchmark-with-pprof-fb7070ee1a94
// go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out

// run pprof tool
// go tool pprof [profile_file.out]
// >web
