package utils

import "github.com/wule61/go-f5-soap/global_lb"

func VSPagination(resources []global_lb.VirtualServerDefinition, pageSize int) [][]global_lb.VirtualServerDefinition {
	return splitVSArray(resources, splitArrayCnt(len(resources), pageSize), pageSize)
}

func splitVSArray(sources []global_lb.VirtualServerDefinition, num, pageSize int) [][]global_lb.VirtualServerDefinition {
	max := len(sources)
	if max < num {
		return nil
	}

	var segmens = make([][]global_lb.VirtualServerDefinition, 0)
	quantity := pageSize
	end := 0
	for i := 1; i <= num; i++ {
		qu := i * quantity
		if i != num {
			segmens = append(segmens, sources[i-1+end:qu])
		} else {
			segmens = append(segmens, sources[i-1+end:])
		}
		end = qu - i
	}
	return segmens
}

func GeneralPagination(resources []interface{}, pageSize int) [][]interface{} {
	return GeneralsplitArray(resources, splitArrayCnt(len(resources), pageSize), pageSize)
}

func GeneralsplitArray(sources []interface{}, num, pageSize int) [][]interface{} {
	max := len(sources)
	if max < num {
		return nil
	}

	var segmens = make([][]interface{}, 0)
	quantity := pageSize
	end := 0
	for i := 1; i <= num; i++ {
		qu := i * quantity
		if i != num {
			segmens = append(segmens, sources[i-1+end:qu])
		} else {
			segmens = append(segmens, sources[i-1+end:])
		}
		end = qu - i
	}
	return segmens
}

func Pagination(resources []string, pageSize int) [][]string {
	return splitArray(resources, splitArrayCnt(len(resources), pageSize), pageSize)
}

func splitArray(sources []string, num, pageSize int) [][]string {
	max := len(sources)
	if max < num {
		return nil
	}
	var segmens = make([][]string, 0)
	quantity := pageSize
	end := 0
	for i := 1; i <= num; i++ {
		qu := i * quantity
		if i != num {
			segmens = append(segmens, sources[i-1+end:qu])
		} else {
			segmens = append(segmens, sources[i-1+end:])
		}
		end = qu - i
	}
	return segmens
}

func splitArrayCnt(sourcesLen, pageSize int) int {
	if sourcesLen < pageSize {
		return 1
	}
	s := sourcesLen / pageSize
	y := sourcesLen % pageSize
	if y > 0 {
		return s + 1
	} else {
		return s
	}
}
