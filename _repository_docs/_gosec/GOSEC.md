# Gosec

[IN PROGRESS]

## To Do
- [ ] Showing same/similar things to `golangci-lint`, redundant?

## Install to `./bin/`

`curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s v2.14.0`

## Run with all options:

`bin/gosec ./...` #=>

```
[gosec] 2022/12/05 09:21:33 Including rules: default
[gosec] 2022/12/05 09:21:33 Excluding rules: default
[gosec] 2022/12/05 09:21:33 Import directory: /home/matt/_apps/golangdocker
[gosec] 2022/12/05 09:21:33 Import directory: /home/matt/_apps/golangdocker/api
[gosec] 2022/12/05 09:21:33 Import directory: /home/matt/_apps/golangdocker/common
[gosec] 2022/12/05 09:21:33 Import directory: /home/matt/_apps/golangdocker/config
[gosec] 2022/12/05 09:21:34 Import directory: /home/matt/_apps/golangdocker/docs
[gosec] 2022/12/05 09:21:34 Checking package: common
[gosec] 2022/12/05 09:21:34 Checking file: /home/matt/_apps/golangdocker/common/colorOutput.go
[gosec] 2022/12/05 09:21:34 Checking file: /home/matt/_apps/golangdocker/common/common.go
[gosec] 2022/12/05 09:21:34 Import directory: /home/matt/_apps/golangdocker/sysinfo
[gosec] 2022/12/05 09:21:34 Checking package: config
[gosec] 2022/12/05 09:21:34 Checking file: /home/matt/_apps/golangdocker/config/config.go
[gosec] 2022/12/05 09:21:35 Checking package: docs
[gosec] 2022/12/05 09:21:35 Checking file: /home/matt/_apps/golangdocker/docs/docs.go
[gosec] 2022/12/05 09:21:35 Checking package: sysinfo
[gosec] 2022/12/05 09:21:35 Checking file: /home/matt/_apps/golangdocker/sysinfo/sysinfo.go
[gosec] 2022/12/05 09:21:35 Checking package: api
[gosec] 2022/12/05 09:21:35 Checking file: /home/matt/_apps/golangdocker/api/api.go
[gosec] 2022/12/05 09:21:35 Checking package: main
[gosec] 2022/12/05 09:21:35 Checking file: /home/matt/_apps/golangdocker/main.go
Results:


[/home/matt/_apps/golangdocker/sysinfo/sysinfo.go:38] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    37:         } else {
  > 38:                 rsp.Body.Close()
    39:                 fmt.Printf("     %s Successfully established https connection to: %s\n", common.ConsoleSuccess("[ ✓ SUCCESS ]"), common.ConsoleBold(url))



[/home/matt/_apps/golangdocker/common/colorOutput.go:27] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    26:         out.Write([]byte("\n\n"))
  > 27:         out.WriteTo(os.Stdout)
    28: }



[/home/matt/_apps/golangdocker/api/api.go:120-122] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    119:        loadInfo := sysinfo.GetLoadInfo(c)
  > 120:        c.Status(200).JSON(&fiber.Map{
  > 121:                "loadInfo": loadInfo,
  > 122:        })
    123:        return nil



[/home/matt/_apps/golangdocker/api/api.go:104-106] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    103:        netInfo := sysinfo.GetNetInfo(c)
  > 104:        c.Status(200).JSON(&fiber.Map{
  > 105:                "netInfo": netInfo,
  > 106:        })
    107:        return nil



[/home/matt/_apps/golangdocker/api/api.go:88-90] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    87:         hostInfo := sysinfo.GetHostInfo(c)
  > 88:         c.Status(200).JSON(&fiber.Map{
  > 89:                 "hostInfo": hostInfo,
  > 90:         })
    91:         return nil



[/home/matt/_apps/golangdocker/api/api.go:72-74] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    71:         cpuInfo := sysinfo.GetCPUInfo(c)
  > 72:         c.Status(200).JSON(&fiber.Map{
  > 73:                 "cpuInfo": cpuInfo,
  > 74:         })
    75:         return nil



[/home/matt/_apps/golangdocker/api/api.go:56-58] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    55:         memInfo := sysinfo.GetMemInfo(c)
  > 56:         c.Status(200).JSON(&fiber.Map{
  > 57:                 "memInfo": memInfo,
  > 58:         })
    59:         return nil



[/home/matt/_apps/golangdocker/api/api.go:40-42] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    39:         apiRoutes := sysinfo.GetAPIRoutes(c)
  > 40:         c.Status(200).JSON(&fiber.Map{
  > 41:                 "apiRoutes": apiRoutes,
  > 42:         })
    43:         return nil



[/home/matt/_apps/golangdocker/api/api.go:26] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    25: func apiFalseRoot(c *fiber.Ctx) error {
  > 26:         c.Redirect("/api/v1")
    27:         return nil



Summary:
  Gosec  : 2.14.0
  Files  : 7
  Lines  : 635
  Nosec  : 0
  Issues : 9
```