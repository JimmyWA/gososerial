package gososerial

import (
	"github.com/JimmyWA/Gososerial/ysoserial/gadget"
)

func GetCC1(cmd string) []byte {
	return gadget.GetCommonsCollections1(cmd)
}

func GetCC2(cmd string) []byte {
	return gadget.GetCommonsCollections2(cmd)
}

func GetCC3(cmd string) []byte {
	return gadget.GetCommonsCollections3(cmd)
}

func GetCC4(cmd string) []byte {
	return gadget.GetCommonsCollections4(cmd)
}

func GetCC5(cmd string) []byte {
	return gadget.GetCommonsCollections5(cmd)
}

func GetCC6(cmd string) []byte {
	return gadget.GetCommonsCollections6(cmd)
}

func GetCC7(cmd string) []byte {
	return gadget.GetCommonsCollections7(cmd)
}

func GetAllNames() []string {
	return []string{
		gadget.CC1,
		gadget.CC2,
		gadget.CC3,
		gadget.CC4,
		gadget.CC5,
		gadget.CC6,
		gadget.CC7,
	}
}
