package flights

import (
	"fmt"
	"strings"
	"time"

	pkio "github.com/PassKit/passkit-golang-grpc-sdk/io"
)

const (
	CarrierCode        = "YY"
	DesignatorRevision = 2
)

var (
	runSeed             = time.Now().UnixNano()
	Origin              = "YY4"
	Destination         = "ADP"
	FlightNumber        = fmt.Sprintf("%03d", time.Now().UnixNano()%900+100)
	OperatingCarrierPNR = fmt.Sprintf("P%05d", positiveMod(time.Now().UnixNano(), 100000))
	departure           = time.Now().AddDate(0, 0, 14)
	DepartureDate       = &pkio.Date{Year: int32(departure.Year()), Month: int32(departure.Month()), Day: int32(departure.Day())}
	DepartureDateTime   = departure.Format("2006-01-02") + "T13:00:00"
	CarrierCreated      bool
	CreatedAirports     = map[string]bool{}
)

func positiveMod(value, divisor int64) int64 {
	result := value % divisor
	if result < 0 {
		return result + divisor
	}
	return result
}

func alphaCode(seed int64, width int) string {
	if seed < 0 {
		seed = -seed
	}
	var code strings.Builder
	for index := 0; index < width; index++ {
		code.WriteByte(byte('A' + seed%26))
		seed /= 26
	}
	return code.String()
}
