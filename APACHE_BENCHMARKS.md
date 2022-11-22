# Apache Benchmarks

## GRAIN OF SALT REQUIRED

These results are not meant to be real world examples. The benchmarks below, while on two separate hosts, are on the same intenal network accessing the server by direct IP Address. This is by no means what happens in the real world, as we are skipping DNS resolution, remote server physical distances, and the usual network infrastructure and routing that would normally sit in front of the server. By skipping all of these real-world layers, the results are skewed.

However, what it does show is the rather large traffic, request, and throughput capacity that this tiny Docker contaier has.

## Results

`ab -n 5000 -c 100 http://192.168.0.99:5000/api/v1/cpu`

```
Server Software:
Server Hostname:        192.168.0.99
Server Port:            5000

Document Path:          /api/v1/cpu
Document Length:        6077 bytes

Concurrency Level:      100
Time taken for tests:   3.560 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      31030000 bytes
HTML transferred:       30385000 bytes
Requests per second:    1404.48 [#/sec] (mean)
Time per request:       71.201 [ms] (mean)
Time per request:       0.712 [ms] (mean, across all concurrent requests)
Transfer rate:          8511.92 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.4      0      11
Processing:     5   70   7.5     69     198
Waiting:        3   44  18.1     46     195
Total:          6   70   7.5     70     198
ERROR: The median and mean for the initial connection time are more than twice the standard
       deviation apart. These results are NOT reliable.

Percentage of the requests served within a certain time (ms)
  50%     70
  66%     71
  75%     71
  80%     72
  90%     74
  95%     76
  98%     80
  99%     80
 100%    198 (longest request)
```

Of Note:

Number Of Requests:     5000
Concurrency Level:      100 (Equivalent to 50 batches of 100 simultaneous requests)
Document Length:        6077 bytes

Time taken for tests:   3.560 seconds
Total Data Transferred: 31.03MB
Requests per second:    1404.48 [#/sec] (mean)
Transfer rate:          8511.92 [Kbytes/sec] received