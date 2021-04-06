go test -v  -run ^TestQuickSort$ algorithm/go/sort
go test -v   algorithm/go/sort

如果使用 -bench 或 -v 标志，则 go test 会输出完整的输出，甚至是通过包测试，以显示所请求的基准测试结果或详细日志记录。

使用-run标志，

https://www.cnblogs.com/sunsky303/p/11818480.html