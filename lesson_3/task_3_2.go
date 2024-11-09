package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', -1, 64)
}

func Counting(sls string) (string, error) {
	// sls := expression[st:end]
	total_sls := sls
	s_indexes := []int{0}
	for i := 0; i < len(sls); i++ {
		if string(sls[i]) == "*" || string(sls[i]) == "/" || string(sls[i]) == "+" || string(sls[i]) == "-" {
			s_indexes = append(s_indexes, i)
		}
	}
	if string(sls[len(sls)-1]) == "*" || string(sls[len(sls)-1]) == "/" || string(sls[len(sls)-1]) == "-" || string(sls[len(sls)-1]) == "+" {
		return "0.0", fmt.Errorf("UNCORRECT")
	}
	// fmt.Println(s_indexes)
	if string(sls[0]) == "-" && len(s_indexes) == 2 {
		return sls, nil
	}
	if string(sls[0]) == "-" {
		s_indexes = s_indexes[1:]
	}
	s_indexes = append(s_indexes, len(sls))
	// fmt.Println(s_indexes)
	if len(s_indexes) > 2 {
		s_indexes[0] = -1
		for i := 1; i < len(s_indexes)-1; i++ {
			if string(sls[s_indexes[i]]) == "*" {
				if string(sls[s_indexes[i]+1]) == "-" {
					a1, err1 := strconv.ParseFloat(sls[s_indexes[i-1]+1:s_indexes[i]], 64)
					a2, err2 := strconv.ParseFloat(sls[s_indexes[i]+1:s_indexes[i+2]], 64)
					if err1 == nil && err2 == nil {
						// // fmt.Println(sls[s_indexes[i-1]+1 : s_indexes[i+1]])
						total_sls = strings.Replace(total_sls, sls[s_indexes[i-1]+1:s_indexes[i+2]], FloatToString(a1*a2), 1)
						// fmt.Println(total_sls, sls, "*")
						return Counting(total_sls)
					} else {
						return "0.0", fmt.Errorf("Uncorrect")
					}
				} else {
					a1, err1 := strconv.ParseFloat(sls[s_indexes[i-1]+1:s_indexes[i]], 64)
					a2, err2 := strconv.ParseFloat(sls[s_indexes[i]+1:s_indexes[i+1]], 64)
					if err1 == nil && err2 == nil {
						// // fmt.Println(sls[s_indexes[i-1]+1 : s_indexes[i+1]])
						total_sls = strings.Replace(total_sls, sls[s_indexes[i-1]+1:s_indexes[i+1]], FloatToString(a1*a2), 1)
						// fmt.Println(total_sls, sls, "*")
						return Counting(total_sls)
					} else {
						return "0.0", fmt.Errorf("Uncorrect")
					}
				}

			} else if string(sls[s_indexes[i]]) == "/" {
				if string(sls[s_indexes[i]+1]) == "-" {
					a1, err1 := strconv.ParseFloat(sls[s_indexes[i-1]+1:s_indexes[i]], 64)
					a2, err2 := strconv.ParseFloat(sls[s_indexes[i]+1:s_indexes[i+2]], 64)
					if err1 == nil && err2 == nil {
						if a2 != 0 {
							total_sls = strings.Replace(total_sls, sls[s_indexes[i-1]+1:s_indexes[i+1]], FloatToString(a1/a2), 1)
							// fmt.Println(total_sls, sls, "/")
							return Counting(total_sls)
						} else if a2 == 0 {
							return "0.0", errors.New("/ 0")
						}
					}
				} else {
					a1, err1 := strconv.ParseFloat(sls[s_indexes[i-1]+1:s_indexes[i]], 64)
					a2, err2 := strconv.ParseFloat(sls[s_indexes[i]+1:s_indexes[i+1]], 64)
					if err1 == nil && err2 == nil {
						if a2 != 0 {
							// // fmt.Println(sls[s_indexes[i-1]+1:s_indexes[i+1]])
							total_sls = strings.Replace(total_sls, sls[s_indexes[i-1]+1:s_indexes[i+1]], FloatToString(a1/a2), 1)
							// fmt.Println(total_sls, sls, "/")
							return Counting(total_sls)

						} else if a2 == 0 {
							return "0.0", errors.New("/ 0")
						}
					} else {
						return "0.0", fmt.Errorf("Uncorrect")
					}
				}
			}
		}

		for i := 1; i < len(s_indexes)-1; i++ {
			if string(sls[s_indexes[i]]) == "+" {
				if string(sls[s_indexes[i]+1]) == "-" {
					a1, err1 := strconv.ParseFloat(sls[s_indexes[i-1]+1:s_indexes[i]], 64)
					a2, err2 := strconv.ParseFloat(sls[s_indexes[i]+2:s_indexes[i+2]], 64)
					if err1 == nil && err2 == nil {
						// // fmt.Println(sls[s_indexes[i-1]+1 : s_indexes[i+1]])
						total_sls = strings.Replace(total_sls, sls[s_indexes[i-1]+1:s_indexes[i+2]], FloatToString(a1-a2), 1)
						// fmt.Println(total_sls, sls, "+-")
						return Counting(total_sls)
					} else {
						return "0.0", fmt.Errorf("Uncorrect")
					}
				} else {
					a1, err1 := strconv.ParseFloat(sls[s_indexes[i-1]+1:s_indexes[i]], 64)
					a2, err2 := strconv.ParseFloat(sls[s_indexes[i]+1:s_indexes[i+1]], 64)
					if err1 == nil && err2 == nil {
						// // fmt.Println(sls[s_indexes[i-1]+1 : s_indexes[i+1]])
						total_sls = strings.Replace(total_sls, sls[s_indexes[i-1]+1:s_indexes[i+1]], FloatToString(a1+a2), 1)
						// fmt.Println(total_sls, sls, "+")
						return Counting(total_sls)
					} else {
						return "0.0", fmt.Errorf("Uncorrect")
					}
				}
			} else if string(sls[s_indexes[i]]) == "-" {
				if string(sls[s_indexes[i]+1]) == "-" {
					a1, err1 := strconv.ParseFloat(sls[s_indexes[i-1]+1:s_indexes[i]], 64)
					a2, err2 := strconv.ParseFloat(sls[s_indexes[i]+2:s_indexes[i+2]], 64)
					if err1 == nil && err2 == nil {
						// // fmt.Println(sls[s_indexes[i-1]+1 : s_indexes[i+1]])
						total_sls = strings.Replace(total_sls, sls[s_indexes[i-1]+1:s_indexes[i+2]], FloatToString(a1+a2), 1)
						// fmt.Println(total_sls, sls, "--")
						return Counting(total_sls)
					} else {
						return "0.0", fmt.Errorf("Uncorrect")
					}
				} else {
					a1, err1 := strconv.ParseFloat(sls[s_indexes[i-1]+1:s_indexes[i]], 64)
					a2, err2 := strconv.ParseFloat(sls[s_indexes[i]+1:s_indexes[i+1]], 64)
					if err1 == nil && err2 == nil {
						// // fmt.Println(sls[s_indexes[i-1]+1 : s_indexes[i+1]])
						total_sls = strings.Replace(total_sls, sls[s_indexes[i-1]+1:s_indexes[i+1]], FloatToString(a1-a2), 1)
						// fmt.Println(total_sls, sls, "-")
						return Counting(total_sls)
					} else {
						return "0.0", fmt.Errorf("Uncorrect")
					}
				}
			}
		}
	} else {
		return sls, nil
	}
	return "0.0", nil
}

func Brackets_processing(expression string) (string, error) {
	st_ind := []int{}
	end_ind := []int{}
	s := ""

	for i := 0; i < len(expression); i++ {
		s = string(expression[i])
		if s == "(" {
			st_ind = append(st_ind, i+1)
		}
		if s == ")" {
			end_ind = append(end_ind, i)
			// // fmt.Println(Counting(expression[st_ind:end_ind]))
		}
	}
	if len(st_ind) != len(end_ind) {
		return "0.0", fmt.Errorf("UNCORRECT INPUT")
	} else if len(st_ind) == 0 {
		return expression, nil
	} else {
		total_exp := expression
		// fmt.Println(st_ind, end_ind)
		// fmt.Println(total_exp)

		if len(st_ind) == 1 {
			s, err := Counting(expression[st_ind[0]:end_ind[0]])
			if err == nil {
				total_exp = strings.Replace(total_exp, total_exp[st_ind[0]-1:end_ind[0]+1], s, 1)
				// fmt.Println(total_exp)
				return total_exp, nil
			} else {
				return "0.0", fmt.Errorf("COUNTING ERROR")
			}
		} else {
			// fmt.Println(expression)
			c := 0
			for st_ind[len(st_ind)-1] > end_ind[c] {
				c++
			}
			// fmt.Println(expression[st_ind[len(st_ind)-1]-1 : end_ind[c]+1])
			s, err := Counting(expression[st_ind[len(st_ind)-1]:end_ind[c]])
			if err == nil {
				total_exp = strings.Replace(total_exp, total_exp[st_ind[len(st_ind)-1]-1:end_ind[c]+1], s, 1)
				// fmt.Println(total_exp)
				return Brackets_processing(total_exp)
			}
		}
	}
	return "0.0", nil
}

func Calc(expression string) (float64, error) {
	if expression == "" {
		return 0.0, fmt.Errorf("UNCORRRECT")
	}
	s1, err := Brackets_processing(expression)
	if err == nil {
		t_s, err1 := Counting(s1)
		if err1 == nil {
			t, err2 := strconv.ParseFloat(t_s, 64)
			if err2 == nil {
				return t, err2
			} else {
				return 0.0, fmt.Errorf("UNCORRECT")
			}
		} else {
			return 0.0, fmt.Errorf("UNCORRECT")
		}
	}
	return 0.0, fmt.Errorf("UNCORRRECT")
}

func main() {
	// fmt.Println(Calc("1+(2+2)/4"))
	// fmt.Println(Calc("1+((2*3)/1/8.2-(2*(3-7)))/0.5"))
	// fmt.Println(Calc("12*(10+12+4*10-20*(2/2))/1-4"))
	fmt.Println(Calc(""))
}
