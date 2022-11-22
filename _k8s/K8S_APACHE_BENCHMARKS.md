# Apache Benchmarks

#### 1 Replica

`.\bin\abs.exe -n 5000 -c 10 http://golang.0nezer0.com/api/v1/cpu`

```
Server Software:
Server Hostname:        golang.0nezer0.com
Server Port:            80

Document Path:          /api/v1/cpu
Document Length:        3593 bytes

Concurrency Level:      10
Time taken for tests:   28.452 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      18610000 bytes
HTML transferred:       17965000 bytes
Requests per second:    175.73 [#/sec] (mean)
Time per request:       56.905 [ms] (mean)
Time per request:       5.690 [ms] (mean, across all concurrent requests)
Transfer rate:          638.74 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.4      1       3
Processing:     4   56  42.8     86     192
Waiting:        2   53  42.6     84     191
Total:          5   57  42.7     87     192

Percentage of the requests served within a certain time (ms)
  50%     87
  66%     92
  75%     94
  80%     95
  90%     98
  95%    100
  98%    102
  99%    104
 100%    192 (longest request)
```

#### 2 Replicas

`.\bin\abs.exe -n 5000 -c 10 http://golang.0nezer0.com/api/v1/cpu`

```
Server Software:
Server Hostname:        golang.0nezer0.com
Server Port:            80

Document Path:          /api/v1/cpu
Document Length:        3593 bytes

Concurrency Level:      10
Time taken for tests:   15.450 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      18610000 bytes
HTML transferred:       17965000 bytes
Requests per second:    323.62 [#/sec] (mean)
Time per request:       30.900 [ms] (mean)
Time per request:       3.090 [ms] (mean, across all concurrent requests)
Transfer rate:          1176.29 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.4      1       3
Processing:     4   30  32.7      9     182
Waiting:        2   27  32.3      7     182
Total:          5   31  32.7     11     182

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     16
  75%     68
  80%     75
  90%     84
  95%     89
  98%     94
  99%     98
 100%    182 (longest request)
```

#### 3 Replicas

`.\bin\abs.exe -n 5000 -c 10 http://golang.0nezer0.com/api/v1/cpu`

```
Server Software:
Server Hostname:        golang.0nezer0.com
Server Port:            80

Document Path:          /api/v1/cpu
Document Length:        3593 bytes

Concurrency Level:      10
Time taken for tests:   11.339 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      18610000 bytes
HTML transferred:       17965000 bytes
Requests per second:    440.95 [#/sec] (mean)
Time per request:       22.678 [ms] (mean)
Time per request:       2.268 [ms] (mean, across all concurrent requests)
Transfer rate:          1602.77 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.4      1       3
Processing:     4   22  25.0      9     169
Waiting:        3   19  24.6      7     169
Total:          5   23  25.0     10     169

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     13
  75%     20
  80%     43
  90%     70
  95%     79
  98%     87
  99%     91
 100%    169 (longest request)
```
