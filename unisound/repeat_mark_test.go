package unisound

import (
	"encoding/json"
	"testing"
)

func addWs1(ws []interface{}) (addWs []interface{}) {
	m1 := make(map[string]interface{})
	m1["text"] = "黄"
	m1["type"] = float64(0)
	ws = append(ws, m1)

	m2 := make(map[string]interface{})
	m2["text"] = "河"
	m2["type"] = float64(0)
	ws = append(ws, m2)

	m3 := make(map[string]interface{})
	m3["text"] = "入"
	m3["type"] = float64(0)
	ws = append(ws, m3)

	m4 := make(map[string]interface{})
	m4["text"] = "sli"
	m4["type"] = float64(4)
	ws = append(ws, m4)

	m5 := make(map[string]interface{})
	m5["text"] = "黄"
	m5["type"] = float64(2)
	ws = append(ws, m5)

	m6 := make(map[string]interface{})
	m6["text"] = "河"
	m6["type"] = float64(2)
	ws = append(ws, m6)

	m7 := make(map[string]interface{})
	m7["text"] = "入"
	m7["type"] = float64(2)
	ws = append(ws, m7)

	m8 := make(map[string]interface{})
	m8["text"] = "海"
	m8["type"] = float64(2)
	ws = append(ws, m8)

	m9 := make(map[string]interface{})
	m9["text"] = "流"
	m9["type"] = float64(2)
	ws = append(ws, m9)

	m10 := make(map[string]interface{})
	m10["text"] = "我"
	m10["type"] = float64(0)
	ws = append(ws, m10)

	m11 := make(map[string]interface{})
	m11["text"] = "我"
	m11["type"] = float64(2)
	ws = append(ws, m11)

	m12 := make(map[string]interface{})
	m12["text"] = "你"
	m12["type"] = float64(0)
	ws = append(ws, m12)

	m13 := make(map[string]interface{})
	m13["text"] = "你"
	m13["type"] = float64(2)
	ws = append(ws, m13)

	m14 := make(map[string]interface{})
	m14["text"] = "们"
	m14["type"] = float64(2)
	ws = append(ws, m14)

	// 多个字重复
	// 都为0的压缩为一个判断
	m15 := make(map[string]interface{})
	m15["text"] = "门"
	m15["type"] = float64(0)
	ws = append(ws, m15)

	m16 := make(map[string]interface{})
	m16["text"] = "门"
	m16["type"] = float64(0)
	ws = append(ws, m16)

	m17 := make(map[string]interface{})
	m17["text"] = "门"
	m17["type"] = float64(2)
	ws = append(ws, m17)

	m18 := make(map[string]interface{})
	m18["text"] = "禁"
	m18["type"] = float64(2)
	ws = append(ws, m18)

	m19 := make(map[string]interface{})
	m19["text"] = "黄"
	m19["type"] = float64(0)
	ws = append(ws, m19)

	m20 := make(map[string]interface{})
	m20["text"] = "黄"
	m20["type"] = float64(0)
	ws = append(ws, m20)

	m21 := make(map[string]interface{})
	m21["text"] = "黄"
	m21["type"] = float64(2)
	ws = append(ws, m21)

	m22 := make(map[string]interface{})
	m22["text"] = "红"
	m22["type"] = float64(0)
	ws = append(ws, m22)

	m23 := make(map[string]interface{})
	m23["text"] = "红"
	m23["type"] = float64(0)
	ws = append(ws, m23)

	m24 := make(map[string]interface{})
	m24["text"] = "红"
	m24["type"] = float64(2)
	ws = append(ws, m24)

	m25 := make(map[string]interface{})
	m25["text"] = "红"
	m25["type"] = float64(2)
	ws = append(ws, m25)

	m26 := make(map[string]interface{})
	m26["text"] = "黄"
	m26["type"] = float64(0)
	ws = append(ws, m26)

	m27 := make(map[string]interface{})
	m27["text"] = "黄"
	m27["type"] = float64(0)
	ws = append(ws, m27)

	m28 := make(map[string]interface{})
	m28["text"] = "黄"
	m28["type"] = float64(2)
	ws = append(ws, m28)

	m29 := make(map[string]interface{})
	m29["text"] = "河"
	m29["type"] = float64(2)
	ws = append(ws, m29)

	m30 := make(map[string]interface{})
	m30["text"] = "入"
	m30["type"] = float64(2)
	ws = append(ws, m30)

	m31 := make(map[string]interface{})
	m31["text"] = "河"
	m31["type"] = float64(0)
	ws = append(ws, m31)

	m32 := make(map[string]interface{})
	m32["text"] = "黄"
	m32["type"] = float64(0)
	ws = append(ws, m32)

	m33 := make(map[string]interface{})
	m33["text"] = "黄"
	m33["type"] = float64(2)
	ws = append(ws, m33)

	m34 := make(map[string]interface{})
	m34["text"] = "河"
	m34["type"] = float64(2)
	ws = append(ws, m34)

	m35 := make(map[string]interface{})
	m35["text"] = "黄"
	m35["type"] = float64(0)
	ws = append(ws, m35)

	m36 := make(map[string]interface{})
	m36["text"] = "黄"
	m36["type"] = float64(2)
	ws = append(ws, m36)

	m37 := make(map[string]interface{})
	m37["text"] = "黄"
	m37["type"] = float64(0)
	ws = append(ws, m37)

	m38 := make(map[string]interface{})
	m38["text"] = "黄"
	m38["type"] = float64(0)
	ws = append(ws, m38)

	m39 := make(map[string]interface{})
	m39["text"] = "黄"
	m39["type"] = float64(2)
	ws = append(ws, m39)

	//m40 := make(map[string]interface{})
	//m40["text"] = "海"
	//m40["type"] = float64(0)
	//ws = append(ws, m40)
	return ws
}

func addWs2(ws []interface{}) (addWs []interface{}) {
	//m1 := make(map[string]interface{})
	//m1["text"] = "黄"
	//m1["type"] = float64(0)
	//ws = append(ws, m1)
	//
	//m2 := make(map[string]interface{})
	//m2["text"] = "河"
	//m2["type"] = float64(0)
	//ws = append(ws, m2)
	//
	//m3 := make(map[string]interface{})
	//m3["text"] = "入"
	//m3["type"] = float64(0)
	//ws = append(ws, m3)
	//
	//m4 := make(map[string]interface{})
	//m4["text"] = "海"
	//m4["type"] = float64(0)
	//ws = append(ws, m4)
	//
	//m5 := make(map[string]interface{})
	//m5["text"] = "流"
	//m5["type"] = float64(0)
	//ws = append(ws, m5)
	//
	//m6 := make(map[string]interface{})
	//m6["text"] = "黄"
	//m6["type"] = float64(2)
	//ws = append(ws, m6)
	//
	//m7 := make(map[string]interface{})
	//m7["text"] = "河"
	//m7["type"] = float64(2)
	//ws = append(ws, m7)
	//
	//m8 := make(map[string]interface{})
	//m8["text"] = "入"
	//m8["type"] = float64(2)
	//ws = append(ws, m8)
	//
	//m9 := make(map[string]interface{})
	//m9["text"] = "海"
	//m9["type"] = float64(2)
	//ws = append(ws, m9)
	//
	//m10 := make(map[string]interface{})
	//m10["text"] = "流"
	//m10["type"] = float64(2)
	//ws = append(ws, m10)
	//
	/////////////////////////
	//m11 := make(map[string]interface{})
	//m11["text"] = "黄"
	//m11["type"] = float64(0)
	//ws = append(ws, m11)
	//
	//m12 := make(map[string]interface{})
	//m12["text"] = "河"
	//m12["type"] = float64(0)
	//ws = append(ws, m12)
	//
	//m13 := make(map[string]interface{})
	//m13["text"] = "黄"
	//m13["type"] = float64(2)
	//ws = append(ws, m13)
	//
	//m14 := make(map[string]interface{})
	//m14["text"] = "河"
	//m14["type"] = float64(2)
	//ws = append(ws, m14)
	//
	//m15 := make(map[string]interface{})
	//m15["text"] = "入"
	//m15["type"] = float64(2)
	//ws = append(ws, m15)
	//
	//m16 := make(map[string]interface{})
	//m16["text"] = "海"
	//m16["type"] = float64(2)
	//ws = append(ws, m16)
	//
	//m17 := make(map[string]interface{})
	//m17["text"] = "流"
	//m17["type"] = float64(2)
	//ws = append(ws, m17)
	//
	////////////////////////
	//m18 := make(map[string]interface{})
	//m18["text"] = "黄"
	//m18["type"] = float64(2)
	//ws = append(ws, m18)
	//
	//m19 := make(map[string]interface{})
	//m19["text"] = "河"
	//m19["type"] = float64(2)
	//ws = append(ws, m19)
	//
	//m20 := make(map[string]interface{})
	//m20["text"] = "入"
	//m20["type"] = float64(0)
	//ws = append(ws, m20)
	//
	//m21 := make(map[string]interface{})
	//m21["text"] = "海"
	//m21["type"] = float64(0)
	//ws = append(ws, m21)
	//
	//m22 := make(map[string]interface{})
	//m22["text"] = "入"
	//m22["type"] = float64(2)
	//ws = append(ws, m22)
	//
	//m23 := make(map[string]interface{})
	//m23["text"] = "海"
	//m23["type"] = float64(2)
	//ws = append(ws, m23)
	//
	//m24 := make(map[string]interface{})
	//m24["text"] = "流"
	//m24["type"] = float64(2)
	//ws = append(ws, m24)
	//////////////////////////
	//m25 := make(map[string]interface{})
	//m25["text"] = "黄"
	//m25["type"] = float64(2)
	//ws = append(ws, m25)
	//
	//m26 := make(map[string]interface{})
	//m26["text"] = "河"
	//m26["type"] = float64(2)
	//ws = append(ws, m26)
	//
	//m27 := make(map[string]interface{})
	//m27["text"] = "入"
	//m27["type"] = float64(2)
	//ws = append(ws, m27)
	//
	//m28 := make(map[string]interface{})
	//m28["text"] = "海"
	//m28["type"] = float64(0)
	//ws = append(ws, m28)
	//
	//m29 := make(map[string]interface{})
	//m29["text"] = "流"
	//m29["type"] = float64(0)
	//ws = append(ws, m29)
	//
	//m30 := make(map[string]interface{})
	//m30["text"] = "海"
	//m30["type"] = float64(2)
	//ws = append(ws, m30)
	//
	//m31 := make(map[string]interface{})
	//m31["text"] = "流"
	//m31["type"] = float64(2)
	//ws = append(ws, m31)
	////////////////////////////
	//
	//m32 := make(map[string]interface{})
	//m32["text"] = "黄"
	//m32["type"] = float64(0)
	//ws = append(ws, m32)
	//
	//m33 := make(map[string]interface{})
	//m33["text"] = "黄"
	//m33["type"] = float64(0)
	//ws = append(ws, m33)
	//
	//m34 := make(map[string]interface{})
	//m34["text"] = "黄"
	//m34["type"] = float64(2)
	//ws = append(ws, m34)
	//
	//m35 := make(map[string]interface{})
	//m35["text"] = "河"
	//m35["type"] = float64(2)
	//ws = append(ws, m35)
	//
	//m36 := make(map[string]interface{})
	//m36["text"] = "入"
	//m36["type"] = float64(2)
	//ws = append(ws, m36)
	//
	//m37 := make(map[string]interface{})
	//m37["text"] = "海"
	//m37["type"] = float64(2)
	//ws = append(ws, m37)
	//
	//m38 := make(map[string]interface{})
	//m38["text"] = "流"
	//m38["type"] = float64(2)
	//ws = append(ws, m38)
	//
	///////////////////////
	//
	//m39 := make(map[string]interface{})
	//m39["text"] = "黄"
	//m39["type"] = float64(2)
	//ws = append(ws, m39)
	//
	//m40 := make(map[string]interface{})
	//m40["text"] = "河"
	//m40["type"] = float64(2)
	//ws = append(ws, m40)
	//
	//m41 := make(map[string]interface{})
	//m41["text"] = "入"
	//m41["type"] = float64(0)
	//ws = append(ws, m41)
	//
	//m42 := make(map[string]interface{})
	//m42["text"] = "入"
	//m42["type"] = float64(0)
	//ws = append(ws, m42)
	//
	//m43 := make(map[string]interface{})
	//m43["text"] = "入"
	//m43["type"] = float64(2)
	//ws = append(ws, m43)
	//
	//m44 := make(map[string]interface{})
	//m44["text"] = "海"
	//m44["type"] = float64(2)
	//ws = append(ws, m44)
	//
	//m45 := make(map[string]interface{})
	//m45["text"] = "流"
	//m45["type"] = float64(2)
	//ws = append(ws, m45)
	//
	///////////////////////
	//
	//m46 := make(map[string]interface{})
	//m46["text"] = "黄"
	//m46["type"] = float64(2)
	//ws = append(ws, m46)
	//
	//m47 := make(map[string]interface{})
	//m47["text"] = "河"
	//m47["type"] = float64(2)
	//ws = append(ws, m47)
	//
	//m48 := make(map[string]interface{})
	//m48["text"] = "入"
	//m48["type"] = float64(2)
	//ws = append(ws, m48)
	//
	//m49 := make(map[string]interface{})
	//m49["text"] = "海"
	//m49["type"] = float64(2)
	//ws = append(ws, m49)
	//
	//m50 := make(map[string]interface{})
	//m50["text"] = "流"
	//m50["type"] = float64(0)
	//ws = append(ws, m50)
	//
	//m51 := make(map[string]interface{})
	//m51["text"] = "流"
	//m51["type"] = float64(0)
	//ws = append(ws, m51)
	//
	//m52 := make(map[string]interface{})
	//m52["text"] = "流"
	//m52["type"] = float64(2)
	//ws = append(ws, m52)

	///////////////
	//m55 := make(map[string]interface{})
	//m55["text"] = "黄"
	//m55["type"] = float64(2)
	//ws = append(ws, m55)
	//
	//m56 := make(map[string]interface{})
	//m56["text"] = "河"
	//m56["type"] = float64(2)
	//ws = append(ws, m56)
	//
	//m57 := make(map[string]interface{})
	//m57["text"] = "黄"
	//m57["type"] = float64(0)
	//ws = append(ws, m57)
	//
	//m58 := make(map[string]interface{})
	//m58["text"] = "河"
	//m58["type"] = float64(0)
	//ws = append(ws, m58)
	//
	//m59 := make(map[string]interface{})
	//m59["text"] = "入"
	//m59["type"] = float64(2)
	//ws = append(ws, m59)
	//
	//m60 := make(map[string]interface{})
	//m60["text"] = "海"
	//m60["type"] = float64(2)
	//ws = append(ws, m60)
	//
	//m61 := make(map[string]interface{})
	//m61["text"] = "流"
	//m61["type"] = float64(2)
	//ws = append(ws, m61)

	//////////////////////////
	//m62 := make(map[string]interface{})
	//m62["text"] = "黄"
	//m62["type"] = float64(2)
	//ws = append(ws, m62)
	//
	//m63 := make(map[string]interface{})
	//m63["text"] = "河"
	//m63["type"] = float64(2)
	//ws = append(ws, m63)
	//
	//m64 := make(map[string]interface{})
	//m64["text"] = "入"
	//m64["type"] = float64(2)
	//ws = append(ws, m64)
	//
	//m65 := make(map[string]interface{})
	//m65["text"] = "海"
	//m65["type"] = float64(2)
	//ws = append(ws, m65)
	//
	//m661 := make(map[string]interface{})
	//m661["text"] = "流"
	//m661["type"] = float64(2)
	//ws = append(ws, m661)
	//
	//m66 := make(map[string]interface{})
	//m66["text"] = "黄"
	//m66["type"] = float64(0)
	//ws = append(ws, m66)
	//
	//m67 := make(map[string]interface{})
	//m67["text"] = "河"
	//m67["type"] = float64(0)
	//ws = append(ws, m67)
	//
	//m68 := make(map[string]interface{})
	//m68["text"] = "入"
	//m68["type"] = float64(0)
	//ws = append(ws, m68)
	//
	//m69 := make(map[string]interface{})
	//m69["text"] = "海"
	//m69["type"] = float64(0)
	//ws = append(ws, m69)
	//
	//m70 := make(map[string]interface{})
	//m70["text"] = "流"
	//m70["type"] = float64(0)
	//ws = append(ws, m70)
	////////////////////////
	m71 := make(map[string]interface{})
	m71["text"] = "黄"
	m71["type"] = float64(2)
	ws = append(ws, m71)
	m72 := make(map[string]interface{})
	m72["text"] = "河"
	m72["type"] = float64(2)
	ws = append(ws, m72)
	m73 := make(map[string]interface{})
	m73["text"] = "入"
	m73["type"] = float64(2)
	ws = append(ws, m73)
	m74 := make(map[string]interface{})
	m74["text"] = "海"
	m74["type"] = float64(2)
	ws = append(ws, m74)
	m75 := make(map[string]interface{})
	m75["text"] = "流"
	m75["type"] = float64(2)
	ws = append(ws, m75)
	m76 := make(map[string]interface{})
	m76["text"] = "流"
	m76["type"] = float64(0)
	ws = append(ws, m76)
	m77 := make(map[string]interface{})
	m77["text"] = "流"
	m77["type"] = float64(0)
	ws = append(ws, m77)
	return ws
}

func TestRepeatMark(t *testing.T) {
	var ws []interface{}
	ws = make([]interface{}, 0)

	ws = addWs2(ws)
	b, _ := json.Marshal(ws)
	t.Log("input:", string(b))

	markWs, markNum := markRepeat(ws)

	t.Log("mark number is:", markNum)
	b, _ = json.Marshal(markWs)
	t.Log("output:", string(b))
}

func isSameSli(in []interface{}) bool {
	if len(in) <= 1 {
		return false
	}
	m := make(map[string]int8)
	for i := range in {
		word := in[i].(map[string]interface{})
		m[word["text"].(string)] = 0
	}
	if len(m) == 1 {
		return true
	}
	return false
}

func markSameZeroSli(zeroSliCache []interface{}, normalSliCache []interface{}) (mark uint32) {
	// multi words is same
	zeroSame := isSameSli(zeroSliCache)
	if zeroSame {
		wordZ := zeroSliCache[0].(map[string]interface{})
		wordN := normalSliCache[len(normalSliCache)-1].(map[string]interface{})
		txtZ := wordZ["text"]
		txtN := wordN["text"]
		if txtZ == txtN {
			wordN["repeat"] = 1
			return mark1
		}
	}
	return mark0
}

const (
	mark0 = iota
	mark1
)

func markRepeat(ws []interface{}) (markWs []interface{}, markNum uint32) {
	zeroSliCache := make([]interface{}, 0)
	normalSliCache := make([]interface{}, 0)

	for j := range ws {
		word := ws[j].(map[string]interface{})
		ty := word["type"].(float64)
		switch ty {
		case 2:
			// pre mark not repeat
			word["repeat"] = 0
			// 0 0 0 2 0
			if len(zeroSliCache) > len(normalSliCache) && (len(zeroSliCache) != 0 && len(normalSliCache) != 0) {
				mark := markSameZeroSli(zeroSliCache, normalSliCache)
				markNum += mark
				zeroSliCache = zeroSliCache[:0]
				normalSliCache = normalSliCache[:0]
			}
			normalSliCache = append(normalSliCache, ws[j]) // 0 0 0
		case 0:
			// pre mark not repeat
			word["repeat"] = 0
			// first increase zeroSliCache
			if len(normalSliCache) == 0 {
				continue
			}
			zeroSliCache = append(zeroSliCache, ws[j])
			nLen := len(normalSliCache)
			zLen := len(zeroSliCache)
			// not equal and not last one
			if nLen != zLen && j+1 != len(ws) {
				continue
			}
			// 0 0 0 2 2 2
			wordGroupRepeat := true
			if nLen == zLen {
				for i := range zeroSliCache {
					txtZ := zeroSliCache[i].(map[string]interface{})
					txtN := normalSliCache[i].(map[string]interface{})
					if txtN["text"] != txtZ["text"] {
						wordGroupRepeat = false
					}
				}
			} else {
				wordGroupRepeat = false
			}

			if wordGroupRepeat {
				for i := range normalSliCache {
					word := normalSliCache[i].(map[string]interface{})
					word["repeat"] = 1
					markNum += mark1
				}
				zeroSliCache = zeroSliCache[:0]
				normalSliCache = normalSliCache[:0]
			} else {
				mark := markSameZeroSli(zeroSliCache, normalSliCache)
				markNum += mark
				zeroSliCache = zeroSliCache[:0]
				normalSliCache = normalSliCache[:0]
			}
		}
	}
	return ws, markNum
}
