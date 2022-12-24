package main

import (
	"testing"
)

func BenchmarkMap(b *testing.B) {
	b.ReportAllocs()
	list_all_files("/Users/carlos/Downloads/enron_mail_20110402/maildir")
}

// run test benchmark
// https://medium.com/@felipedutratine/profile-your-benchmark-with-pprof-fb7070ee1a94

// run pprof tool
// go tool pprof mem.prof
// >web
