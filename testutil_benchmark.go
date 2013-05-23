package qbs

import (
	"fmt"
	"runtime"
	"testing"
)

func doBenchmarkFind(b *testing.B) {
	b.StopTimer()
	bas := new(basic)
	bas.Name = "Basic"
	bas.State = 3
	mg, _ := GetMigration()
	mg.DropTable(bas)
	mg.CreateTableIfNotExists(bas)
	q, _ := GetQbs()
	q.Save(bas)
	closeMigrationAndQbs(mg, q)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ba := new(basic)
		ba.Id = 1
		q, _ = GetQbs()
		err := q.Find(ba)
		if err != nil {
			panic(err)
		}
		q.Close()
	}
	b.StopTimer()
	runtime.GC()
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("alloc:%d, total:%d\n", stats.Alloc, stats.TotalAlloc)
}

func doBenchmarkTransaction(b *testing.B) {
	b.StopTimer()
	bas := new(basic)
	bas.Name = "Basic"
	bas.State = 3
	mg, _ := GetMigration()
	mg.DropTable(bas)
	mg.CreateTableIfNotExists(bas)
	q, _ := GetQbs()
	q.Save(bas)
	closeMigrationAndQbs(mg, q)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ba := new(basic)
		ba.Id = 1
		q, _ = GetQbs()
		q.Begin()
		q.Find(ba)
		q.Close()
	}
	b.StopTimer()
	runtime.GC()
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("alloc:%d, total:%d\n", stats.Alloc, stats.TotalAlloc)
}
