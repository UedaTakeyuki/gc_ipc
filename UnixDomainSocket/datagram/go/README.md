# Performance
```
pi@raspberrypi:~/github/gc_ipc/UnixDomainSocket/datagram/go/server $ uname -a
Linux raspberrypi 5.10.103-v7+ #1529 SMP Tue Mar 8 12:21:37 GMT 2022 armv7l GNU/Linux
pi@raspberrypi:~/github/gc_ipc/UnixDomainSocket/datagram/go/server $ python -m getrpimodel
2 Model B
pi@raspberrypi:~/github/gc_ipc/UnixDomainSocket/datagram/go/server $ go run *.go
2024/06/07 10:08:08 556 ns
2024/06/07 10:08:08 eraps main.elapse: 1699 μs
2024/06/07 10:08:35 324 ns
2024/06/07 10:08:35 eraps main.elapse: 297 μs
2024/06/07 10:08:53 325 ns
2024/06/07 10:08:53 eraps main.elapse: 356 μs
2024/06/07 10:09:05 325 ns
2024/06/07 10:09:05 eraps main.elapse: 334 μs
2024/06/07 10:09:26 369 ns
2024/06/07 10:09:26 eraps main.elapse: 315 μs
```
