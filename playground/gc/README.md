

```shell
GODEBUG=gctrace=1 go run main.go

# gc即将运行时大小 ——>gc结束操作时堆大小 ——>活动堆大小
gc 1 @0.001s 15%: 0.004+1.9+0.003 ms clock, 0.034+1.3/2.8/1.9+0.029 ms cpu, 5->6->6 MB, 6 MB goal, 8 P
gc 2 @0.009s 8%: 0.003+5.8+0.014 ms clock, 0.025+0.42/6.6/1.0+0.11 ms cpu, 14->14->13 MB, 15 MB goal, 8 P
gc 1 @0.002s 0%: 0.006+1.3+0.004 ms clock, 0.051+0/0.011/0.81+0.036 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 2 @0.003s 24%: 0.043+1.8+1.8 ms clock, 0.34+0/0.009/0.27+14 ms cpu, 5->10->7 MB, 6 MB goal, 8 P
gc 3 @0.010s 15%: 0.044+2.6+0.027 ms clock, 0.35+0/0.024/2.7+0.22 ms cpu, 14->15->4 MB, 15 MB goal, 8 P

```

