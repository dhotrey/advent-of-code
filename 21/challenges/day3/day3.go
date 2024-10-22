package day3

import (
	"21/utils"
	"strconv"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(3, mode)
	defer file.Close()
	const zeroPos = 0
	const onePos = 1
	diagnostics := []string{}

	bitArr := [][]int{} // [ [zeroCounter,oneCounter],[19,2] . . . ]
	for data.Scan() {
		log.Debug(data.Text())
		diagnostics = append(diagnostics, data.Text())
		for i, bit := range data.Text() {
			log.Debug("rune literal", "rune", bit)

			if i == len(bitArr) {
				bitArr = append(bitArr, []int{0, 0})
				log.Debugf("creating zero and one counter at %d : %d", i, bitArr[i])
			}

			// bitVal := strconv.QuoteRuneToASCII(bit)
			// log.Debugf("got bitVal %s at %d", string(bitVal), i)
			log.Debugf("got bit as rune %v", bit)

			switch bit {
			case 49:
				bitArr[i][onePos]++
			case 48:
				bitArr[i][zeroPos]++
			default:
				log.Debug("Switch case not working")
			}
			log.Debugf("BitVal at %d post increment :%v", i, bitArr[i])
			log.Debug("")
		}
	}
	log.Infof("final bitArr %v", bitArr)
	gammaRateStr := ""
	epsilonRateStr := ""

	for _, val := range bitArr {
		if val[zeroPos] > val[onePos] {
			gammaRateStr += "0"
			epsilonRateStr += "1"
		} else {
			gammaRateStr += "1"
			epsilonRateStr += "0"
		}
	}
	log.Infof("Gamma rate %s", gammaRateStr)

	gammaRate, _ := strconv.ParseInt(gammaRateStr, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateStr, 2, 64)
	log.Info("Got gamma and epsilon rates", "gamma", gammaRate, "epslion rate", epsilonRate)
	powerConsumption := gammaRate * epsilonRate
	log.Infof("Power Consumption is : %d", powerConsumption)

	var o2Ratingb, co2Ratingb string
	if bitArr[0][zeroPos] > bitArr[0][onePos] {
		o2Ratingb = getRating(diagnostics, "o2", "0", 0, 1)
		co2Ratingb = getRating(diagnostics, "co2", "1", 0, 1)
	} else {
		o2Ratingb = getRating(diagnostics, "o2", "1", 0, 1)
		co2Ratingb = getRating(diagnostics, "co2", "0", 0, 1)
	}

	o2Rating, _ := strconv.ParseInt(o2Ratingb, 2, 64)
	co2Rating, _ := strconv.ParseInt(co2Ratingb, 2, 64)
	log.Info("Got o2 and co2 rating", "o2Rating", o2Rating, "co2Rating", co2Rating)
	lifeSupportRating := o2Rating * co2Rating
	log.Infof("Got lifesupport rating : %d", lifeSupportRating)

}

func getRating(data []string, ratingType string, chosenBit string, currBitIdx int, recursionTracker int) string {
	bitFreq := map[string]int{"0": 0, "1": 0}
	getDefaultBit := map[string]string{"o2": "1", "co2": "0"}
	refinedDiagnostics := []string{}

	for _, binStr := range data {
		if string(binStr[currBitIdx]) == chosenBit {
			refinedDiagnostics = append(refinedDiagnostics, binStr)
		}
	}
	if len(refinedDiagnostics) == 1 {
		log.Infof("function called %d times for %s rating", recursionTracker, ratingType)
		return refinedDiagnostics[0] // all the values filtered out
	}

	for _, refinedBinStr := range refinedDiagnostics {
		if currBitIdx < len(refinedBinStr)-1 {
			nextChar := string(refinedBinStr[currBitIdx+1])
			bitFreq[nextChar]++
		} else if currBitIdx == len(refinedBinStr)-1 {
			log.Infof("function called %d times for %s rating", recursionTracker, ratingType)
			return refinedBinStr // Final return
		}
	}

	recursionTracker++
	switch ratingType {
	case "o2":
		if bitFreq["0"] > bitFreq["1"] {
			return getRating(refinedDiagnostics, ratingType, "0", currBitIdx+1, recursionTracker)
		} else if bitFreq["1"] > bitFreq["0"] {
			return getRating(refinedDiagnostics, ratingType, "1", currBitIdx+1, recursionTracker)
		} else { // when it's equal
			return getRating(refinedDiagnostics, ratingType, getDefaultBit[ratingType], currBitIdx+1, recursionTracker)
		}
	default: // co2
		if bitFreq["0"] > bitFreq["1"] {
			return getRating(refinedDiagnostics, ratingType, "1", currBitIdx+1, recursionTracker)
		} else if bitFreq["1"] > bitFreq["0"] {
			return getRating(refinedDiagnostics, ratingType, "0", currBitIdx+1, recursionTracker)
		} else { // when it's equal
			return getRating(refinedDiagnostics, ratingType, getDefaultBit[ratingType], currBitIdx+1, recursionTracker)
		}
	}
}
