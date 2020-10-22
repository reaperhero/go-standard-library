package third

import (
	"github.com/tealeg/xlsx"
	"os"
	"testing"
)

func Test_xlsx(t *testing.T) {

	//
	xFile := xlsx.NewFile()
	sheet, _ := xFile.AddSheet("Sheet1")

	//
	var ce1 *xlsx.Cell
	head := sheet.AddRow()
	ce1 = head.AddCell()
	ce1.Value = "用户名"
	ce1 = head.AddCell()
	ce1.Value = "密码"


	file, _ := os.OpenFile("观众白名单.xlsx", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	xFile.Write(file)
}
