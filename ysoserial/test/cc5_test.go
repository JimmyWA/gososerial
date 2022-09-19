package test

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	gososerial "github.com/JimmyWA/gososerial"
)

func TestCC5(t *testing.T) {
	ser := gososerial.GetCC5("calc.exe")
	serStr := strings.ToUpper(hex.EncodeToString(ser))
	stdStr := "ACED00057372002E6A617661782E6D616E6167656D656E742E42616441747472696275746556616C7565457870457863657074696F6ED4E7DAAB632D46400200014C000376616C7400124C6A6176612F6C616E672F4F626A6563743B787200136A6176612E6C616E672E457863657074696F6ED0FD1F3E1A3B1CC4020000787200136A6176612E6C616E672E5468726F7761626C65D5C635273977B8CB0300044C000563617573657400154C6A6176612F6C616E672F5468726F7761626C653B4C000D64657461696C4D6573736167657400124C6A6176612F6C616E672F537472696E673B5B000A737461636B547261636574001E5B4C6A6176612F6C616E672F537461636B5472616365456C656D656E743B4C001473757070726573736564457863657074696F6E737400104C6A6176612F7574696C2F4C6973743B787071007E0008707572001E5B4C6A6176612E6C616E672E537461636B5472616365456C656D656E743B02462A3C3CFD22390200007870000000077372001B6A6176612E6C616E672E537461636B5472616365456C656D656E746109C59A2636DD8502000449000A6C696E654E756D6265724C000E6465636C6172696E67436C61737371007E00054C000866696C654E616D6571007E00054C000A6D6574686F644E616D6571007E000578700000005174002679736F73657269616C2E7061796C6F6164732E436F6D6D6F6E73436F6C6C656374696F6E7335740018436F6D6D6F6E73436F6C6C656374696F6E73352E6A6176617400096765744F626A6563747371007E000B0000003371007E000D71007E000E71007E000F7371007E000B0000002874002779736F73657269616C2E7061796C6F6164732E7574696C2E5061796C6F616452756E6E657224317400125061796C6F616452756E6E65722E6A61766174000463616C6C7371007E000B0000002171007E001271007E001371007E00147371007E000B0000004874002C79736F73657269616C2E7365636D67722E45786563436865636B696E6753656375726974794D616E6167657274002045786563436865636B696E6753656375726974794D616E616765722E6A61766174000B63616C6C577261707065647371007E000B0000002174002579736F73657269616C2E7061796C6F6164732E7574696C2E5061796C6F616452756E6E657271007E001374000372756E7371007E000B0000005C71007E000D71007E000E7400046D61696E737200266A6176612E7574696C2E436F6C6C656374696F6E7324556E6D6F6469666961626C654C697374FC0F2531B5EC8E100200014C00046C69737471007E00077872002C6A6176612E7574696C2E436F6C6C656374696F6E7324556E6D6F6469666961626C65436F6C6C656374696F6E19420080CB5EF71E0200014C0001637400164C6A6176612F7574696C2F436F6C6C656374696F6E3B7870737200136A6176612E7574696C2E41727261794C6973747881D21D99C7619D03000149000473697A657870000000007704000000007871007E002478737200346F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E6B657976616C75652E546965644D6170456E7472798AADD29B39C11FDB0200024C00036B657971007E00014C00036D617074000F4C6A6176612F7574696C2F4D61703B7870740003666F6F7372002A6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E6D61702E4C617A794D61706EE594829E7910940300014C0007666163746F727974002C4C6F72672F6170616368652F636F6D6D6F6E732F636F6C6C656374696F6E732F5472616E73666F726D65723B78707372003A6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E66756E63746F72732E436861696E65645472616E73666F726D657230C797EC287A97040200015B000D695472616E73666F726D65727374002D5B4C6F72672F6170616368652F636F6D6D6F6E732F636F6C6C656374696F6E732F5472616E73666F726D65723B78707572002D5B4C6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E5472616E73666F726D65723BBD562AF1D83418990200007870000000057372003B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E66756E63746F72732E436F6E7374616E745472616E73666F726D6572587690114102B1940200014C000969436F6E7374616E7471007E00017870767200116A6176612E6C616E672E52756E74696D65000000000000000000000078707372003A6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E66756E63746F72732E496E766F6B65725472616E73666F726D657287E8FF6B7B7CCE380200035B000569417267737400135B4C6A6176612F6C616E672F4F626A6563743B4C000B694D6574686F644E616D6571007E00055B000B69506172616D54797065737400125B4C6A6176612F6C616E672F436C6173733B7870757200135B4C6A6176612E6C616E672E4F626A6563743B90CE589F1073296C02000078700000000274000A67657452756E74696D65757200125B4C6A6176612E6C616E672E436C6173733BAB16D7AECBCD5A990200007870000000007400096765744D6574686F647571007E003C00000002767200106A6176612E6C616E672E537472696E67A0F0A4387A3BB34202000078707671007E003C7371007E00357571007E003900000002707571007E003900000000740006696E766F6B657571007E003C00000002767200106A6176612E6C616E672E4F626A656374000000000000000000000078707671007E00397371007E0035757200135B4C6A6176612E6C616E672E537472696E673BADD256E7E91D7B4702000078700000000174000863616C632E657865740004657865637571007E003C0000000171007E00417371007E0031737200116A6176612E6C616E672E496E746567657212E2A0A4F781873802000149000576616C7565787200106A6176612E6C616E672E4E756D62657286AC951D0B94E08B020000787000000001737200116A6176612E7574696C2E486173684D61700507DAC1C31660D103000246000A6C6F6164466163746F724900097468726573686F6C6478703F40000000000000770800000010000000007878"
	fmt.Println(serStr)
	fmt.Println(stdStr)
	fmt.Println(serStr == stdStr)
}
