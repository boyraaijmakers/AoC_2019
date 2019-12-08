package sixteen

import (
	"fmt"
	"strconv"
)

func Sixteen() {
	input := `222222222222212222222121222022222222222222222222222222222202222222222222222222022022022220202221202022220102202222212212202222222221112222220221222222221222222222212222222120222122222222222222222222222222222212222222222222222222122222022221202220202022222212222222202202202222222221002222221220222222220222222222212222222020222022222222222222222222222222222202222222222222222222002022022220202221202222222112202222202202202222222221022222222222222222222222222222222222222120222222222222222222222222222222222202222222222222222222112222122220222220212022222222212222222222222222222222222222220222222222221222222222202222222122222122222222222222222222222222222202222222222222222222102122022220212220202022021212222222202212212222222222122222220222222222222222222222222222222221222022222222222222222222222222222222222222222222222222202122022220222221222122022212222222222202202222222220002222222221222222221222212222222222222022222222222222222222222222222222222212222222222222222222222022222220202220222222021102222222202202222222222221012222220222222222221222212222212222222221222022222222222222222222222222222222222222222222222222022222222222212220202022220102202222212212212222222222022222222222222222220222202222202222222221222122222222022222222222222222222222222222222222222222102222022220222220212022121122202222212212202222222221112222220220222222222222212222202222222022222222222222122222222222222222222202222222222222222222022022022222222222222022022202212222222202212222222222112222222221222222221222202222212222222120222022222222022222222222222222222222122222222222222222022122222222212220212222221002212222202202202222222220012222221222222222222222212222212222222220222122222222122222222222202222212202222221222222222222122222122220212222222022220102222222212202202222222220222222220222222222221222212222212222222121222222222222122222222222202222222202022220222222222222022022222221222222212022020112212222202202222222222222012222222220222222222222212222212221222222222122222222122222222222212222222222122222222222222222212122022221222221222022222112112222222202222222222222212222222221222222222222212222222221222222222022222222122222222222222222222212122220222222202222102022222222202222212022122122022222202222222222222220012222220222222222221222202222222220222120222022222222022222222222222222202202022222222221202222102222022222221220212022120022102222202212222222222221212222200220222222222222202222212222222120222222222222122222222222222222212202122200222222212222222222122221221220222022122102122222222212202222222220002222200220222222221222222222212221222022222022222222222222222222212222202202022210222221222222002122122222211220212022020212102220202222102222222221122222201221222222222222222222222222222020222022202222022222222222202222222212222221222220202222222222122220212221212122120222002221202222012222222222002222212222222222221222202222202220222221222022220222222222222222212222222212122220222221212222002122122222222220222022222002122221222202012222222220102222200220222222222222212222222220222122222022200222022222222222222222202222022222202222212222012022122220200222222222222202022221212212202222222220002222200222222222222222222222212222222021222022221222222222222222202222222222222200222222222222222022222220212222222122220012212221212202022220222222022222202221222222221222222222212221222121222022200222122222222222202222212212022212222021222222222122222220220220202112121202212222222222012222222221102222210220222222222222202222102220222220222122201222222222222222202222202212122211202120212222202122022220210222222112022012122221222222012222222220122222210221222222220222212222222221222220222022201222222222222222202222202222222212222221212222002022222220222222212002120202122222202212112220222222022222212122222222221222202222212222222020222122220222022222022222202222222222022200222222212222112022022220211221212102122002102221222222102222222220022222210020222222221222202222222221222122222222222222222222022222222222202212122212222222222222212022022221220221212122221002112221202222222222222221012122222121222220221222202222012220222121222122220222122222222222212222212212222222212122202222202222222220220222212202020122122220222202212222222220202022210120222222222222212222010220222120222222202222022222122222212222212202022212212122222222102022222221212222202002022122002222202202202222222220212022222222222221211222202222021221222022222222201222022222122222202222222212022221212022202222102122022222210222212112222002122220222222112221222221002122212121222221201222212222202220222120222122222222222222222222212222222212222202202121202222002122122221200220012112122012222222202212012221222221022022220022222222210222222222020221222121222122212222022222122222202222212222122202222121222222102222122221210220112122220202122222202212222222222222002122210222222222221222212222212220222021222122201222222222222222212222212222222222222222212222102222122222201221002002122202202222222212002220222222012122212122222212211222222222000221222020222222212222022222122222212222222212222201222220212222102122022220201221122212120212102221202222200222222220122022211022222200220222222222121222222021222222210222222222222222202222212212122222202022202222022112022220202222222102122022002222212222110222222220012022221222222202222222222222012220222120222022112122022222022222202222202222222220212122212222002112022221212221122012122112022220202202122221222222012022222121222201202222212222202221222221222122202122022222122222212220202212122200212222202222202222122220202221112122221022122221212212021221222222002222221122222211222222222222011222202222222022100122122222122222222220202212022221212121222222102212122220202222212112021202012222222222111222222222222022211121222220201222202222212222222022222222220222122222122222202222212212222210212121212222102022122220220220212102120022122220202202121221122221122022200122222221200222212222011222202120222122000122022222122222222221202222122211212221202222012012222220210222102002012222112222012202200222022221002222202122222202221222212222001220202020222122002022122222122222222220212212122220202220212222102112122221211221022022200202102220022202121222122221012122221222222202201222202222121220202020222022002122222222222222222221222222122202222221202222112102122221200221002012212012012220102212220221122220012122212121222211211222202222212220202220222122210022122222022222202220222212221200222221210222202212022222200222012112002102012222122222012222222222212022200121222201222222212222210220222121222022211022220222022222222220212212220222222222202222102012122221221222202022220202102221122222012220122221222022222121222202220222212222101222212121222122000122220222222222222220222202020221202020211222202102122222222222102102112212012220012222112222022221012122200220222221210222202222222222212020222122212222022222222222202222202202120220222020221220122012122221220221222022210102022221112212001220022221212122211120222202202222202222221220212021222122101022020222122222222222222212021202222122221221002002022220201221222122010212222221002222222200222220022122201121222222212222212022112220212222222222212122220222022222212220212212020201212020212222122212222221221222000022112122002222102202200221122222212122211122222211200222212122212221222221222022020202022222022222222221212212221201202120211220002202122220220220210102020122012222002222111212022221002122210122222201222222222222212221202221222022220002221222220222212222212202020222212122201221122022122222202220120212000202012222112222022201122222012122221220222202202222222021011220212020222022121122120222220222212221202222021200212221212222222212022220201220000222000122102220222212101211122020012022210221222212210222202020120220212220222022120202121222022222212200212212021202202122222222002112122221202221012102020122202221122202012212222221102122200021222211200222222221200220202121212122202122220222122222212201212212120221202020222221012102022221200221221112202012002222222212201200122120102022210021222202220222202222200221012221222022221222222222222222212202012202120210202221201221202102022220222220211022022202222220112202010220022220102012212220222202201222222020222221002022212222122012022222222222222201102202221222222222221220212012022222210220011002011202212221212202200200022221002022200221222220210222202022000220112020222122021002220222121222212022022212120212202022221121212222022222221220111222100002122220122212112210122122102222220220222221210222222020101221202120212022201002121222120202022012012222122222222120212021022012022221201220020102210022012222112202002212122221212222211220222222222222202221102222102220212122012012020222120222212122022202120220212222201021202212122222211222002222002222222222122222102200220021122222201120022210220222222221201221202222202022212102122222021222102222002202122201212222222222011012222222201220210222201112212220002222220211220222222112222022122200200222222221220221022021202022211222221222021202102221112202120222212122221122200022022222212221201212212202212221022222000211221221022202210020222212211222222021121222022221212022200122122222021202222120202222120210202121212121200002222221201221021212121202112222022202022220222122202212200121022201211222212221202220222222212222001122020222022222012211202212221221012121200120002002122221220220201202210102212222222222010221222221102102201021022200201222212020120221022221222122001002122222021222002210002212222220122222221221100112022220211220022212222020122222002202012201122020102012211221122221212222222020112222222220212122100102020222021222012112102212021222122120210122010202122222200221120122001201112221012222120212021120102012202221022220202222222221112221122221220122200012220222221212022000202202121210212222221121100022122221221220012212111000122221012220211222221222222122222221122202210222212221100221202020210022200112102222120202012221022002022201002020222121122222222221211222210112220211212221012222202202122020202222200220222202221222202222212221202021221222121012010222221202122202202202021201022020220122112222122220222221020022211110112220202212220202020221002022210222122210210222222022002222212022212122020012211222121222222202002012222220212222221221000202022222112220022212212010102221122211220221121021212222222020022211212222202022002222122221202222202222002222222202102001102002220220222020201022121002122221100221202012221102212221022202222202121220002002200222122211220222212122201221002120200022102102120222122222212012202202020220122220220122011022122222021220112002122112202221122212021212220220012212211222222202220222222120022220112021222222022112211222120212222121222022120202122020202022202122222220200221001112202022112221122210222211222222222022210022122221212222212122122221002122200122112102012222020202222122002222122220112220221120101222122220200220100112212121102220202212101222222222012012220120022200222222202122100220202221222022220122120222020202012002112222022222212221201021020102222220101221022102112010012220102220202201220122222002202120022220200222212122210220102021220222110122021222222210222111002222022222222121202221111002220220122220212022020021002221112201200211221120012002200011222202211222212012022020002221221022211022111222121200202202222202121210122020210021100212020221101220200212021211112220012220000220020221122102220222122211210222222100212221222222220022212012120222122212102101202122221202102211200222111102021221111222222012201211222222222200021212122021012202212002022211211222212010121220222020200022112112122222121212102120112002122220012110222022100202122202000222122022122002022221222201010212021021022002221222022220222222222020011121222020212222200012210222021201112012222022022222222101200120202222020212111220212102110112222222222211221220120122222222212201222210211222202102201122022021210022010212202222222210102221122202221201022210221222102112022220011221020122010221012220112221010210222221022102222021022212200222212221201220012121212122112222011222222222122201112022222200102210220020120112120221002221220002202012212222022220002210020022010121212122222200201222222212001221102221220222110202100222220200022021012002222202112000222122100202121201220220121102120211102222202220001201021022222011222020222212210222222102201221222122220012002112202222020202022210122112020222122022201222212202221222201222112102012200212221222211221202122021200000220002112211211222202102211122022220202222021102220222021201022011102012020202022122200022000012121202100221220112021022222222112222110221222020101120212000212202220222222020102220002120211012020002200222220200012121212212122211022120201021222122221221201220101002022112022221122220112220222122102021210002112222201222222201221020222222202002110012102222120222202111122122220212212021212220122202220202101222021110202110202221212210022210221021021111200211202221222222202020121022012122202102121022111222222210012000000222220211002222220021101112021201002222020221021201012221222211021210122122211012222022022211202222202102011021112021210212110112020222220211112211121202020221222001220221112002020210010222121220011200012222102210001202221220210111221210002201201222212201222222212011212122222011222222020211121121102222121220222221220221211012122201110221201021020021102221112212112222121120112222200221112221212222212001020121002110200020201021100220221200222100001102020202222021202021020022222201110220220102100021012220112202210221220121211020222001102222200222202022020121002110210220110001122220120221111010201022122211012211211022212012122211000221211022201210202220222222110200122222200200201211122212212022222101022221022222211202010210100222222202221121011202020211002212221222111222120202120221120011122221212222002222020002221221021112222022222202211122212111010221112101201002221010210222020200111011121002021022102010202121111202021201211222010120100001202220212201211000120220212221212121222220222222212021011222122122202010020200001220122201010120221002122111102212201222222212022221212222120111100010122221212212122100122121210120211100122211201122222220020022102220211011101101221220120200212120022112222211102222202020222222021212022221201100210210012222022210201120022021102222200001002202201022222012111221222201211122001102010220021200002012202022110102122221222121220222120202022220122010220101102220022211111120121021211010220111212200221122222121102221122101210101111111221220022202122121020212001120002110211220020212222222002221010001010121102220022212120010021022102110210100112210021022202010222221112200200220221001102220111200011101200112012211212221200220101022121222122222202020110012202221112222200121022221122101212202202201000222212012002222002201212110210102011221202201011000201222110021022100220200002022021222210222110010200201202221012211021002121221121102202002112212202022212101200110022021212102120101020222212211121202200012212022212202222110120122222220000222021020222001112221000212022202222220001022220210112200000000110220122121011002020100202220011012211120111210011120101202210120101110211001210011001100011220120121001002022020111001112001110212121122111100`

	var outputImage [6][25]int

	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			outputImage[i][j] = -1
		}
	}

	for i := 0; i < 100; i++ {
		layer := input[150*i : 150*i+150]

		for j := 0; j < 6; j++ {
			for k := 0; k < 25; k++ {
				if outputImage[j][k] == -1 && string(layer[j*25+k]) != "2" {
					outputImage[j][k], _ = strconv.Atoi(string(layer[j*25+k]))
				}
			}
		}

	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			if outputImage[i][j] == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("▮")
			}

		}
		fmt.Printf("\n")
	}

}
