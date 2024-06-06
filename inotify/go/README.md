# Performance

```
pi@raspberrypi:~/github/gc_ipc/inotify/go/srv $ uname -a
Linux raspberrypi 5.10.103-v7+ #1529 SMP Tue Mar 8 12:21:37 GMT 2022 armv7l GNU/Linux
pi@raspberrypi:~/github/gc_ipc/inotify/go/srv $ python -m getrpimodel
2 Model B
pi@raspberrypi:~/github/gc_ipc/inotify/go/srv $ go run srv.go
2024/06/06 21:38:53 event: CREATE        "/tmp/go-build2461811398"
2024/06/06 21:38:56 event: WRITE         "/tmp/testinotify"
2024/06/06 21:38:56 modified file: /tmp/testinotify
2024/06/06 21:38:56 2907
2024/06/06 21:38:56 event: WRITE         "/tmp/testinotify"
2024/06/06 21:38:56 modified file: /tmp/testinotify
2024/06/06 21:38:56 3766
2024/06/06 21:38:56 event: REMOVE        "/tmp/go-build2461811398"
2024/06/06 21:39:11 event: CREATE        "/tmp/go-build3923545553"
2024/06/06 21:39:14 event: WRITE         "/tmp/testinotify"
2024/06/06 21:39:14 modified file: /tmp/testinotify
2024/06/06 21:39:14 2607
2024/06/06 21:39:14 event: WRITE         "/tmp/testinotify"
2024/06/06 21:39:14 modified file: /tmp/testinotify
2024/06/06 21:39:14 3426
2024/06/06 21:39:14 event: REMOVE        "/tmp/go-build3923545553"
```
