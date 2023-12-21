module github.com/olcw78/golang/module/hello

go 1.21.5

replace github.com/olcw78/golang/module/greeting => ../greeting

require github.com/olcw78/golang/module/greeting v0.0.0-00010101000000-000000000000
