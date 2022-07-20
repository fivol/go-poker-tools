package equity

import (
	"go-poker-tools/pkg/types"
	"math"
	"testing"
)

func TestCalculateEquity(t *testing.T) {
	table := []struct {
		board      string
		myRange    string
		ranges     []string
		iterations int
		equity     map[string]float32
	}{{
		"6s9c4hQcKd",
		"KsKc",
		[]string{"2s3h"},
		1,
		map[string]float32{"KsKc": 1},
	},
		{
			"6s9c4hQcKd",
			"2s3h",
			[]string{"6h3s"},
			1,
			map[string]float32{"2s3h": 0},
		},
		{
			"6s9c4hQcKd",
			"AsQh,AdAc",
			[]string{"AhQd,AdAh", "2c3c,5d4d"},
			10000,
			map[string]float32{"AdAc": 1, "AsQh": 0.25},
		},
		{
			"6s9c4hQcKd",
			"AsQh,AdAc,2h2s",
			[]string{"AhQd,AdAh,Js8s", "2c3c,5d4d,JsJd"},
			1000000,
			map[string]float32{"2s2h": 0.125, "AdAc": 1, "AsQh": 0.43715},
		},
		{
			"4d8h9s2cKh",
			"9h2c,9h2d,9h2s,9c2h,9c2d,9c2s,9d2h,9d2c,9d2s,9s2h,9s2c,9s2d",
			[]string{"9h2h,9c2c,9d2d,9s2s", "Kh5h,Kc5c,Kd5d,Ks5s,Kh4h,Kc4c,Kd4d,Ks4s,Kh3h,Kc3c,Kd3d,Ks3s,Kh2h,Kc2c,Kd2d,Ks2s"},
			100000,
			map[string]float32{
				"9d2s": 0.335,
				"9c2s": 0.3526,
				"9h2s": 0.375,
			},
		},
		{
			"4d8h9s2cKh",
			"6c5d:1.0,6c5h:1.0,6c5s:1.0,6d5c:1.0,6d5h:1.0,6d5s:1.0,6h5c:1.0,6h5d:1.0,6h5s:1.0,6s5c:1.0,6s5d:1.0,6s5h:1.0,7c4d:1.0,7c4h:1.0,7c4s:1.0,7d4c:1.0,7d4h:1.0,7d4s:1.0,7h4c:1.0,7h4d:1.0,7h4s:1.0,7s4c:1.0,7s4d:1.0,7s4h:1.0,8c3d:1.0,8c3h:1.0,8c3s:1.0,8c4d:1.0,8c4h:1.0,8c4s:1.0,8c5d:1.0,8c5h:1.0,8c5s:1.0,8c6d:1.0,8c6h:1.0,8c6s:1.0,8c7d:1.0,8c7h:1.0,8c7s:1.0,8d3c:1.0,8d3h:1.0,8d3s:1.0,8d4c:1.0,8d4h:1.0,8d4s:1.0,8d5c:1.0,8d5h:1.0,8d5s:1.0,8d6c:1.0,8d6h:1.0,8d6s:1.0,8d7c:1.0,8d7h:1.0,8d7s:1.0,8h3c:1.0,8h3d:1.0,8h3s:1.0,8h4c:1.0,8h4d:1.0,8h4s:1.0,8h5c:1.0,8h5d:1.0,8h5s:1.0,8h6c:1.0,8h6d:1.0,8h6s:1.0,8h7c:1.0,8h7d:1.0,8h7s:1.0,8s3c:1.0,8s3d:1.0,8s3h:1.0,8s4c:1.0,8s4d:1.0,8s4h:1.0,8s5c:1.0,8s5d:1.0,8s5h:1.0,8s6c:1.0,8s6d:1.0,8s6h:1.0,8s7c:1.0,8s7d:1.0,8s7h:1.0,9c3c:1.0,9c4d:1.0,9c4h:1.0,9c4s:1.0,9c6c:1.0,9c8c:1.0,9d3d:1.0,9d4c:1.0,9d4h:1.0,9d4s:1.0,9d6d:1.0,9d8d:1.0,9d9c:1.0,9h3h:1.0,9h4c:1.0,9h4d:1.0,9h4s:1.0,9h6h:1.0,9h8h:1.0,9h9c:1.0,9h9d:1.0,9s3s:1.0,9s4c:1.0,9s4d:1.0,9s4h:1.0,9s6s:1.0,9s8s:1.0,9s9c:1.0,9s9d:1.0,9s9h:1.0,Tc3c:1.0,Tc5c:1.0,Tc5d:1.0,Tc5h:1.0,Tc5s:1.0,Tc6c:1.0,Tc9d:1.0,Tc9h:1.0,Tc9s:1.0,Td3d:1.0,Td5c:1.0,Td5d:1.0,Td5h:1.0,Td5s:1.0,Td6d:1.0,Td9c:1.0,Td9h:1.0,Td9s:1.0,TdTc:1.0,Th3h:1.0,Th5c:1.0,Th5d:1.0,Th5h:1.0,Th5s:1.0,Th6h:1.0,Th9c:1.0,Th9d:1.0,Th9s:1.0,ThTc:1.0,ThTd:1.0,Ts3s:1.0,Ts5c:1.0,Ts5d:1.0,Ts5h:1.0,Ts5s:1.0,Ts6s:1.0,Ts9c:1.0,Ts9d:1.0,Ts9h:1.0,TsTc:1.0,TsTd:1.0,TsTh:1.0,Jc3c:1.0,Jc4c:1.0,Jc6c:1.0,Jc8c:1.0,Jc9c:1.0,JcTc:1.0,Jd3d:1.0,Jd4d:1.0,Jd6d:1.0,Jd8d:1.0,Jd9d:1.0,JdTd:1.0,Jh3h:1.0,Jh4h:1.0,Jh6h:1.0,Jh8h:1.0,Jh9h:1.0,JhTh:1.0,Js3s:1.0,Js4s:1.0,Js6s:1.0,Js8s:1.0,Js9s:1.0,JsTs:1.0,Qc3c:1.0,Qc6c:1.0,QcTc:1.0,Qd3d:1.0,Qd6d:1.0,QdTd:1.0,Qh3h:1.0,Qh6h:1.0,QhTh:1.0,Qs3s:1.0,Qs6s:1.0,QsTs:1.0,Kc3c:1.0,Kc6c:1.0,Kc8c:1.0,Kc9c:1.0,KcTc:1.0,KcQc:1.0,Kd3d:1.0,Kd6d:1.0,Kd8d:1.0,Kd9d:1.0,KdTd:1.0,KdQd:1.0,KdKc:1.0,Kh3h:1.0,Kh6h:1.0,Kh8h:1.0,Kh9h:1.0,KhTh:1.0,KhQh:1.0,KhKc:1.0,KhKd:1.0,Ks3s:1.0,Ks6s:1.0,Ks8s:1.0,Ks9s:1.0,KsTs:1.0,KsQs:1.0,KsKc:1.0,KsKd:1.0,KsKh:1.0,Ac4c:0.55,Ac5c:0.55,Ac9d:1.0,Ac9h:1.0,Ac9s:1.0,AcTd:1.0,AcTh:1.0,AcTs:1.0,AcJd:1.0,AcJh:1.0,AcJs:1.0,AcQd:1.0,AcQh:1.0,AcQs:1.0,AcKd:1.0,AcKh:1.0,AcKs:1.0,Ad4d:0.55,Ad5d:0.55,Ad9c:1.0,Ad9h:1.0,Ad9s:1.0,AdTc:1.0,AdTh:1.0,AdTs:1.0,AdJc:1.0,AdJh:1.0,AdJs:1.0,AdQc:1.0,AdQh:1.0,AdQs:1.0,AdKc:1.0,AdKh:1.0,AdKs:1.0,Ah4h:0.55,Ah5h:0.55,Ah9c:1.0,Ah9d:1.0,Ah9s:1.0,AhTc:1.0,AhTd:1.0,AhTs:1.0,AhJc:1.0,AhJd:1.0,AhJs:1.0,AhQc:1.0,AhQd:1.0,AhQs:1.0,AhKc:1.0,AhKd:1.0,AhKs:1.0,As4s:0.55,As5s:0.55,As9c:1.0,As9d:1.0,As9h:1.0,AsTc:1.0,AsTd:1.0,AsTh:1.0,AsJc:1.0,AsJd:1.0,AsJh:1.0,AsQc:1.0,AsQd:1.0,AsQh:1.0,AsKc:1.0,AsKd:1.0,AsKh:1.0",
			[]string{"3d3c:1.0,3h3c:1.0,3h3d:1.0,3s3c:1.0,3s3d:1.0,3s3h:1.0,4d4c:1.0,4h4c:1.0,4h4d:1.0,4s4c:1.0,4s4d:1.0,4s4h:1.0,5d5c:1.0,5h5c:1.0,5h5d:1.0,5s5c:1.0,5s5d:1.0,5s5h:1.0,6d6c:0.25,6h6c:0.25,6h6d:0.25,6s6c:0.25,6s6d:0.25,6s6h:0.25,7c5d:0.25,7c5h:0.25,7c5s:0.25,7c6c:0.25,7c6d:0.25,7c6h:0.25,7c6s:0.25,7d5c:0.25,7d5h:0.25,7d5s:0.25,7d6c:0.25,7d6d:0.25,7d6h:0.25,7d6s:0.25,7d7c:1.0,7h5c:0.25,7h5d:0.25,7h5s:0.25,7h6c:0.25,7h6d:0.25,7h6h:0.25,7h6s:0.25,7h7c:1.0,7h7d:1.0,7s5c:0.25,7s5d:0.25,7s5h:0.25,7s6c:0.25,7s6d:0.25,7s6h:0.25,7s6s:0.25,7s7c:1.0,7s7d:1.0,7s7h:1.0,8c4d:0.25,8c4h:0.25,8c4s:0.25,8c5c:0.25,8c5d:0.25,8c5h:0.25,8c5s:0.25,8c6d:0.25,8c6h:0.25,8c6s:0.25,8c7c:1.0,8d4c:0.25,8d4h:0.25,8d4s:0.25,8d5c:0.25,8d5d:0.25,8d5h:0.25,8d5s:0.25,8d6c:0.25,8d6h:0.25,8d6s:0.25,8d7d:1.0,8d8c:1.0,8h4c:0.25,8h4d:0.25,8h4s:0.25,8h5c:0.25,8h5d:0.25,8h5h:0.25,8h5s:0.25,8h6c:0.25,8h6d:0.25,8h6s:0.25,8h7h:1.0,8h8c:1.0,8h8d:1.0,8s4c:0.25,8s4d:0.25,8s4h:0.25,8s5c:0.25,8s5d:0.25,8s5h:0.25,8s5s:0.25,8s6c:0.25,8s6d:0.25,8s6h:0.25,8s7s:1.0,8s8c:1.0,8s8d:1.0,8s8h:1.0,9c3d:0.25,9c3h:0.25,9c3s:0.25,9c4c:0.25,9c4d:0.25,9c4h:0.25,9c4s:0.25,9c5d:0.25,9c5h:0.25,9c5s:0.25,9c7c:1.0,9c7d:0.25,9c7h:0.25,9c7s:0.25,9c8c:1.0,9d3c:0.25,9d3h:0.25,9d3s:0.25,9d4c:0.25,9d4d:0.25,9d4h:0.25,9d4s:0.25,9d5c:0.25,9d5h:0.25,9d5s:0.25,9d7c:0.25,9d7d:1.0,9d7h:0.25,9d7s:0.25,9d8d:1.0,9d9c:1.0,9h3c:0.25,9h3d:0.25,9h3s:0.25,9h4c:0.25,9h4d:0.25,9h4h:0.25,9h4s:0.25,9h5c:0.25,9h5d:0.25,9h5s:0.25,9h7c:0.25,9h7d:0.25,9h7h:1.0,9h7s:0.25,9h8h:1.0,9h9c:1.0,9h9d:1.0,9s3c:0.25,9s3d:0.25,9s3h:0.25,9s4c:0.25,9s4d:0.25,9s4h:0.25,9s4s:0.25,9s5c:0.25,9s5d:0.25,9s5h:0.25,9s7c:0.25,9s7d:0.25,9s7h:0.25,9s7s:1.0,9s8s:1.0,9s9c:1.0,9s9d:1.0,9s9h:1.0,Tc3d:0.25,Tc3h:0.25,Tc3s:0.25,Tc4c:0.25,Tc4d:0.25,Tc4h:0.25,Tc4s:0.25,Tc5d:0.25,Tc5h:0.25,Tc5s:0.25,Tc8c:1.0,Tc9c:1.0,Td3c:0.25,Td3h:0.25,Td3s:0.25,Td4c:0.25,Td4d:0.25,Td4h:0.25,Td4s:0.25,Td5c:0.25,Td5h:0.25,Td5s:0.25,Td8d:1.0,Td9d:1.0,TdTc:1.0,Th3c:0.25,Th3d:0.25,Th3s:0.25,Th4c:0.25,Th4d:0.25,Th4h:0.25,Th4s:0.25,Th5c:0.25,Th5d:0.25,Th5s:0.25,Th8h:1.0,Th9h:1.0,ThTc:1.0,ThTd:1.0,Ts3c:0.25,Ts3d:0.25,Ts3h:0.25,Ts4c:0.25,Ts4d:0.25,Ts4h:0.25,Ts4s:0.25,Ts5c:0.25,Ts5d:0.25,Ts5h:0.25,Ts8s:1.0,Ts9s:1.0,TsTc:1.0,TsTd:1.0,TsTh:1.0,Jc2c:0.95,Jc3c:0.95,Jc4c:0.95,Jc4d:0.25,Jc4h:0.25,Jc4s:0.25,Jc5c:0.95,Jc8d:0.25,Jc8h:0.25,Jc8s:0.25,Jc9c:1.0,JcTc:1.0,Jd2d:0.95,Jd3d:0.95,Jd4c:0.25,Jd4d:0.95,Jd4h:0.25,Jd4s:0.25,Jd5d:0.95,Jd8c:0.25,Jd8h:0.25,Jd8s:0.25,Jd9d:1.0,JdTd:1.0,JdJc:1.0,Jh2h:0.95,Jh3h:0.95,Jh4c:0.25,Jh4d:0.25,Jh4h:0.95,Jh4s:0.25,Jh5h:0.95,Jh8c:0.25,Jh8d:0.25,Jh8s:0.25,Jh9h:1.0,JhTh:1.0,JhJc:1.0,JhJd:1.0,Js2s:0.95,Js3s:0.95,Js4c:0.25,Js4d:0.25,Js4h:0.25,Js4s:0.95,Js5s:0.95,Js8c:0.25,Js8d:0.25,Js8h:0.25,Js9s:1.0,JsTs:1.0,JsJc:1.0,JsJd:1.0,JsJh:1.0,Qc5c:0.95,Qc5d:0.25,Qc5h:0.25,Qc5s:0.25,Qc6c:0.95,Qc8d:0.25,Qc8h:0.25,Qc8s:0.25,Qc9c:1.0,QcTc:1.0,QcJc:1.0,Qd5c:0.25,Qd5d:0.95,Qd5h:0.25,Qd5s:0.25,Qd6d:0.95,Qd8c:0.25,Qd8h:0.25,Qd8s:0.25,Qd9d:1.0,QdTd:1.0,QdJd:1.0,QdQc:1.0,Qh5c:0.25,Qh5d:0.25,Qh5h:0.95,Qh5s:0.25,Qh6h:0.95,Qh8c:0.25,Qh8d:0.25,Qh8s:0.25,Qh9h:1.0,QhTh:1.0,QhJh:1.0,QhQc:1.0,QhQd:1.0,Qs5c:0.25,Qs5d:0.25,Qs5h:0.25,Qs5s:0.95,Qs6s:0.95,Qs8c:0.25,Qs8d:0.25,Qs8h:0.25,Qs9s:1.0,QsTs:1.0,QsJs:1.0,QsQc:1.0,QsQd:1.0,QsQh:1.0,Kc6d:0.25,Kc6h:0.25,Kc6s:0.25,Kc7d:0.25,Kc7h:0.25,Kc7s:0.25,Kc8c:1.0,Kc9c:1.0,KcTc:1.0,KcJc:1.0,KcJd:1.0,KcJh:1.0,KcJs:1.0,KcQc:1.0,KcQd:1.0,KcQh:1.0,KcQs:1.0,Kd6c:0.25,Kd6h:0.25,Kd6s:0.25,Kd7c:0.25,Kd7h:0.25,Kd7s:0.25,Kd8d:1.0,Kd9d:1.0,KdTd:1.0,KdJc:1.0,KdJd:1.0,KdJh:1.0,KdJs:1.0,KdQc:1.0,KdQd:1.0,KdQh:1.0,KdQs:1.0,KdKc:1.0,Kh6c:0.25,Kh6d:0.25,Kh6s:0.25,Kh7c:0.25,Kh7d:0.25,Kh7s:0.25,Kh8h:1.0,Kh9h:1.0,KhTh:1.0,KhJc:1.0,KhJd:1.0,KhJh:1.0,KhJs:1.0,KhQc:1.0,KhQd:1.0,KhQh:1.0,KhQs:1.0,KhKc:1.0,KhKd:1.0,Ks6c:0.25,Ks6d:0.25,Ks6h:0.25,Ks7c:0.25,Ks7d:0.25,Ks7h:0.25,Ks8s:1.0,Ks9s:1.0,KsTs:1.0,KsJc:1.0,KsJd:1.0,KsJh:1.0,KsJs:1.0,KsQc:1.0,KsQd:1.0,KsQh:1.0,KsQs:1.0,KsKc:1.0,KsKd:1.0,KsKh:1.0,Ac2c:1.0,Ac3c:1.0,Ac4c:1.0,Ac5c:1.0,Ac6c:1.0,Ac7c:1.0,Ac8c:1.0,Ac9c:1.0,Ac9d:1.0,Ac9h:1.0,Ac9s:1.0,AcTc:1.0,AcTd:1.0,AcTh:1.0,AcTs:1.0,AcJc:1.0,AcJd:1.0,AcJh:1.0,AcJs:1.0,AcQc:1.0,AcQd:1.0,AcQh:1.0,AcQs:1.0,AcKc:1.0,AcKd:1.0,AcKh:1.0,AcKs:1.0,Ad2d:1.0,Ad3d:1.0,Ad4d:1.0,Ad5d:1.0,Ad6d:1.0,Ad7d:1.0,Ad8d:1.0,Ad9c:1.0,Ad9d:1.0,Ad9h:1.0,Ad9s:1.0,AdTc:1.0,AdTd:1.0,AdTh:1.0,AdTs:1.0,AdJc:1.0,AdJd:1.0,AdJh:1.0,AdJs:1.0,AdQc:1.0,AdQd:1.0,AdQh:1.0,AdQs:1.0,AdKc:1.0,AdKd:1.0,AdKh:1.0,AdKs:1.0,AdAc:1.0,Ah2h:1.0,Ah3h:1.0,Ah4h:1.0,Ah5h:1.0,Ah6h:1.0,Ah7h:1.0,Ah8h:1.0,Ah9c:1.0,Ah9d:1.0,Ah9h:1.0,Ah9s:1.0,AhTc:1.0,AhTd:1.0,AhTh:1.0,AhTs:1.0,AhJc:1.0,AhJd:1.0,AhJh:1.0,AhJs:1.0,AhQc:1.0,AhQd:1.0,AhQh:1.0,AhQs:1.0,AhKc:1.0,AhKd:1.0,AhKh:1.0,AhKs:1.0,AhAc:1.0,AhAd:1.0,As2s:1.0,As3s:1.0,As4s:1.0,As5s:1.0,As6s:1.0,As7s:1.0,As8s:1.0,As9c:1.0,As9d:1.0,As9h:1.0,As9s:1.0,AsTc:1.0,AsTd:1.0,AsTh:1.0,AsTs:1.0,AsJc:1.0,AsJd:1.0,AsJh:1.0,AsJs:1.0,AsQc:1.0,AsQd:1.0,AsQh:1.0,AsQs:1.0,AsKc:1.0,AsKd:1.0,AsKh:1.0,AsKs:1.0,AsAc:1.0,AsAd:1.0,AsAh:1.0", "4c3c:0.95,4d3d:0.95,4d4c:0.95,4h3h:0.95,4h4c:0.95,4h4d:0.95,4s3s:0.95,4s4c:0.95,4s4d:0.95,4s4h:0.95,5c3c:0.95,5c3d:0.95,5c3h:0.95,5c3s:0.95,5d3c:0.95,5d3d:0.95,5d3h:0.95,5d3s:0.95,5h3c:0.95,5h3d:0.95,5h3h:0.95,5h3s:0.95,5s3c:0.95,5s3d:0.95,5s3h:0.95,5s3s:0.95,6c2c:0.95,6c3c:0.95,6c3d:0.95,6c3h:0.95,6c3s:0.95,6c5d:0.15,6c5h:0.15,6c5s:0.15,6d2d:0.95,6d3c:0.95,6d3d:0.95,6d3h:0.95,6d3s:0.95,6d5c:0.15,6d5h:0.15,6d5s:0.15,6d6c:0.15,6h2h:0.95,6h3c:0.95,6h3d:0.95,6h3h:0.95,6h3s:0.95,6h5c:0.15,6h5d:0.15,6h5s:0.15,6h6c:0.15,6h6d:0.15,6s2s:0.95,6s3c:0.95,6s3d:0.95,6s3h:0.95,6s3s:0.95,6s5c:0.15,6s5d:0.15,6s5h:0.15,6s6c:0.15,6s6d:0.15,6s6h:0.15,7c2c:0.95,7c2d:0.95,7c2h:0.95,7c2s:0.95,7c3d:0.95,7c3h:0.95,7c3s:0.95,7c4d:0.95,7c4h:0.95,7c4s:0.95,7c5c:0.15,7c5d:0.15,7c5h:0.15,7c5s:0.15,7d2c:0.95,7d2d:0.95,7d2h:0.95,7d2s:0.95,7d3c:0.95,7d3h:0.95,7d3s:0.95,7d4c:0.95,7d4h:0.95,7d4s:0.95,7d5c:0.15,7d5d:0.15,7d5h:0.15,7d5s:0.15,7h2c:0.95,7h2d:0.95,7h2h:0.95,7h2s:0.95,7h3c:0.95,7h3d:0.95,7h3s:0.95,7h4c:0.95,7h4d:0.95,7h4s:0.95,7h5c:0.15,7h5d:0.15,7h5h:0.15,7h5s:0.15,7s2c:0.95,7s2d:0.95,7s2h:0.95,7s2s:0.95,7s3c:0.95,7s3d:0.95,7s3h:0.95,7s4c:0.95,7s4d:0.95,7s4h:0.95,7s5c:0.15,7s5d:0.15,7s5h:0.15,7s5s:0.15,8c2c:0.15,8c2d:0.95,8c2h:0.95,8c2s:0.95,8c3c:0.15,8c3d:0.95,8c3h:0.95,8c3s:0.95,8c4c:0.15,8c5c:0.95,8c5d:0.15,8c5h:0.15,8c5s:0.15,8d2c:0.95,8d2d:0.15,8d2h:0.95,8d2s:0.95,8d3c:0.95,8d3d:0.15,8d3h:0.95,8d3s:0.95,8d4d:0.15,8d5c:0.15,8d5d:0.95,8d5h:0.15,8d5s:0.15,8h2c:0.95,8h2d:0.95,8h2h:0.15,8h2s:0.95,8h3c:0.95,8h3d:0.95,8h3h:0.15,8h3s:0.95,8h4h:0.15,8h5c:0.15,8h5d:0.15,8h5h:0.95,8h5s:0.15,8s2c:0.95,8s2d:0.95,8s2h:0.95,8s2s:0.15,8s3c:0.95,8s3d:0.95,8s3h:0.95,8s3s:0.15,8s4s:0.15,8s5c:0.15,8s5d:0.15,8s5h:0.15,8s5s:0.95,9c2d:0.95,9c2h:0.95,9c2s:0.95,9c3c:0.15,9c4c:0.15,9c8c:0.75,9c8d:0.15,9c8h:0.15,9c8s:0.15,9d2c:0.95,9d2h:0.95,9d2s:0.95,9d3d:0.15,9d4d:0.15,9d8c:0.15,9d8d:0.75,9d8h:0.15,9d8s:0.15,9d9c:0.15,9h2c:0.95,9h2d:0.95,9h2s:0.95,9h3h:0.15,9h4h:0.15,9h8c:0.15,9h8d:0.15,9h8h:0.75,9h8s:0.15,9h9c:0.15,9h9d:0.15,9s2c:0.95,9s2d:0.95,9s2h:0.95,9s3s:0.15,9s4s:0.15,9s8c:0.15,9s8d:0.15,9s8h:0.15,9s8s:0.75,9s9c:0.15,9s9d:0.15,9s9h:0.15,Tc2c:0.15,Tc2d:0.95,Tc2h:0.95,Tc2s:0.95,Tc3c:0.15,Tc3d:0.95,Tc3h:0.95,Tc3s:0.95,Tc4c:0.15,Tc5c:0.15,Tc6c:0.95,Tc7c:0.95,Tc7d:0.15,Tc7h:0.15,Tc7s:0.15,Tc8c:0.75,Tc8d:0.15,Tc8h:0.15,Tc8s:0.15,Tc9d:0.95,Tc9h:0.95,Tc9s:0.95,Td2c:0.95,Td2d:0.15,Td2h:0.95,Td2s:0.95,Td3c:0.95,Td3d:0.15,Td3h:0.95,Td3s:0.95,Td4d:0.15,Td5d:0.15,Td6d:0.95,Td7c:0.15,Td7d:0.95,Td7h:0.15,Td7s:0.15,Td8c:0.15,Td8d:0.75,Td8h:0.15,Td8s:0.15,Td9c:0.95,Td9h:0.95,Td9s:0.95,Th2c:0.95,Th2d:0.95,Th2h:0.15,Th2s:0.95,Th3c:0.95,Th3d:0.95,Th3h:0.15,Th3s:0.95,Th4h:0.15,Th5h:0.15,Th6h:0.95,Th7c:0.15,Th7d:0.15,Th7h:0.95,Th7s:0.15,Th8c:0.15,Th8d:0.15,Th8h:0.75,Th8s:0.15,Th9c:0.95,Th9d:0.95,Th9s:0.95,Ts2c:0.95,Ts2d:0.95,Ts2h:0.95,Ts2s:0.15,Ts3c:0.95,Ts3d:0.95,Ts3h:0.95,Ts3s:0.15,Ts4s:0.15,Ts5s:0.15,Ts6s:0.95,Ts7c:0.15,Ts7d:0.15,Ts7h:0.15,Ts7s:0.95,Ts8c:0.15,Ts8d:0.15,Ts8h:0.15,Ts8s:0.75,Ts9c:0.95,Ts9d:0.95,Ts9h:0.95,Jc2d:0.95,Jc2h:0.95,Jc2s:0.95,Jc3c:0.95,Jc3d:0.95,Jc3h:0.95,Jc3s:0.95,Jc6c:0.15,Jc6d:0.15,Jc6h:0.15,Jc6s:0.15,Jc7c:0.15,Jc7d:0.15,Jc7h:0.15,Jc7s:0.15,Jc8c:0.75,Jc8d:0.95,Jc8h:0.95,Jc8s:0.95,Jc9c:0.15,Jc9d:0.95,Jc9h:0.95,Jc9s:0.95,Jd2c:0.95,Jd2h:0.95,Jd2s:0.95,Jd3c:0.95,Jd3d:0.95,Jd3h:0.95,Jd3s:0.95,Jd6c:0.15,Jd6d:0.15,Jd6h:0.15,Jd6s:0.15,Jd7c:0.15,Jd7d:0.15,Jd7h:0.15,Jd7s:0.15,Jd8c:0.95,Jd8d:0.75,Jd8h:0.95,Jd8s:0.95,Jd9c:0.95,Jd9d:0.15,Jd9h:0.95,Jd9s:0.95,Jh2c:0.95,Jh2d:0.95,Jh2s:0.95,Jh3c:0.95,Jh3d:0.95,Jh3h:0.95,Jh3s:0.95,Jh6c:0.15,Jh6d:0.15,Jh6h:0.15,Jh6s:0.15,Jh7c:0.15,Jh7d:0.15,Jh7h:0.15,Jh7s:0.15,Jh8c:0.95,Jh8d:0.95,Jh8h:0.75,Jh8s:0.95,Jh9c:0.95,Jh9d:0.95,Jh9h:0.15,Jh9s:0.95,Js2c:0.95,Js2d:0.95,Js2h:0.95,Js3c:0.95,Js3d:0.95,Js3h:0.95,Js3s:0.95,Js6c:0.15,Js6d:0.15,Js6h:0.15,Js6s:0.15,Js7c:0.15,Js7d:0.15,Js7h:0.15,Js7s:0.15,Js8c:0.95,Js8d:0.95,Js8h:0.95,Js8s:0.75,Js9c:0.95,Js9d:0.95,Js9h:0.95,Js9s:0.15,Qc2d:0.95,Qc2h:0.95,Qc2s:0.95,Qc3c:0.95,Qc3d:0.95,Qc3h:0.95,Qc3s:0.95,Qc4c:0.95,Qc6c:0.15,Qc6d:0.95,Qc6h:0.95,Qc6s:0.95,Qc7c:0.15,Qc8c:0.75,Qc8d:0.15,Qc8h:0.15,Qc8s:0.15,Qc9c:0.15,Qc9d:0.15,Qc9h:0.15,Qc9s:0.15,QcTc:0.15,QcTd:0.15,QcTh:0.15,QcTs:0.15,QcJc:0.15,QcJd:0.15,QcJh:0.15,QcJs:0.15,Qd2c:0.95,Qd2h:0.95,Qd2s:0.95,Qd3c:0.95,Qd3d:0.95,Qd3h:0.95,Qd3s:0.95,Qd4d:0.95,Qd6c:0.95,Qd6d:0.15,Qd6h:0.95,Qd6s:0.95,Qd7d:0.15,Qd8c:0.15,Qd8d:0.75,Qd8h:0.15,Qd8s:0.15,Qd9c:0.15,Qd9d:0.15,Qd9h:0.15,Qd9s:0.15,QdTc:0.15,QdTd:0.15,QdTh:0.15,QdTs:0.15,QdJc:0.15,QdJd:0.15,QdJh:0.15,QdJs:0.15,QdQc:0.15,Qh2c:0.95,Qh2d:0.95,Qh2s:0.95,Qh3c:0.95,Qh3d:0.95,Qh3h:0.95,Qh3s:0.95,Qh4h:0.95,Qh6c:0.95,Qh6d:0.95,Qh6h:0.15,Qh6s:0.95,Qh7h:0.15,Qh8c:0.15,Qh8d:0.15,Qh8h:0.75,Qh8s:0.15,Qh9c:0.15,Qh9d:0.15,Qh9h:0.15,Qh9s:0.15,QhTc:0.15,QhTd:0.15,QhTh:0.15,QhTs:0.15,QhJc:0.15,QhJd:0.15,QhJh:0.15,QhJs:0.15,QhQc:0.15,QhQd:0.15,Qs2c:0.95,Qs2d:0.95,Qs2h:0.95,Qs3c:0.95,Qs3d:0.95,Qs3h:0.95,Qs3s:0.95,Qs4s:0.95,Qs6c:0.95,Qs6d:0.95,Qs6h:0.95,Qs6s:0.15,Qs7s:0.15,Qs8c:0.15,Qs8d:0.15,Qs8h:0.15,Qs8s:0.75,Qs9c:0.15,Qs9d:0.15,Qs9h:0.15,Qs9s:0.15,QsTc:0.15,QsTd:0.15,QsTh:0.15,QsTs:0.15,QsJc:0.15,QsJd:0.15,QsJh:0.15,QsJs:0.15,QsQc:0.15,QsQd:0.15,QsQh:0.15,Kc2d:0.95,Kc2h:0.95,Kc2s:0.95,Kc5c:0.95,Kc6d:0.95,Kc6h:0.95,Kc6s:0.95,Kc7d:0.95,Kc7h:0.95,Kc7s:0.95,Kd2c:0.95,Kd2h:0.95,Kd2s:0.95,Kd5d:0.95,Kd6c:0.95,Kd6h:0.95,Kd6s:0.95,Kd7c:0.95,Kd7h:0.95,Kd7s:0.95,KdKc:0.95,Kh2c:0.95,Kh2d:0.95,Kh2s:0.95,Kh5h:0.95,Kh6c:0.95,Kh6d:0.95,Kh6s:0.95,Kh7c:0.95,Kh7d:0.95,Kh7s:0.95,KhKc:0.95,KhKd:0.95,Ks2c:0.95,Ks2d:0.95,Ks2h:0.95,Ks5s:0.95,Ks6c:0.95,Ks6d:0.95,Ks6h:0.95,Ks7c:0.95,Ks7d:0.95,Ks7h:0.95,KsKc:0.95,KsKd:0.95,KsKh:0.95,Ac2d:0.95,Ac2h:0.95,Ac2s:0.95,Ac7c:0.95,Ac8c:0.95,Ac8d:0.95,Ac8h:0.95,Ac8s:0.95,Ac9c:0.95,Ac9d:0.95,Ac9h:0.95,Ac9s:0.95,AcTc:0.95,AcJc:0.95,AcJd:0.95,AcJh:0.95,AcJs:0.95,AcQc:0.95,AcKc:0.95,AcKd:0.95,AcKh:0.95,AcKs:0.95,Ad2c:0.95,Ad2h:0.95,Ad2s:0.95,Ad7d:0.95,Ad8c:0.95,Ad8d:0.95,Ad8h:0.95,Ad8s:0.95,Ad9c:0.95,Ad9d:0.95,Ad9h:0.95,Ad9s:0.95,AdTd:0.95,AdJc:0.95,AdJd:0.95,AdJh:0.95,AdJs:0.95,AdQd:0.95,AdKc:0.95,AdKd:0.95,AdKh:0.95,AdKs:0.95,Ah2c:0.95,Ah2d:0.95,Ah2s:0.95,Ah7h:0.95,Ah8c:0.95,Ah8d:0.95,Ah8h:0.95,Ah8s:0.95,Ah9c:0.95,Ah9d:0.95,Ah9h:0.95,Ah9s:0.95,AhTh:0.95,AhJc:0.95,AhJd:0.95,AhJh:0.95,AhJs:0.95,AhQh:0.95,AhKc:0.95,AhKd:0.95,AhKh:0.95,AhKs:0.95,As2c:0.95,As2d:0.95,As2h:0.95,As7s:0.95,As8c:0.95,As8d:0.95,As8h:0.95,As8s:0.95,As9c:0.95,As9d:0.95,As9h:0.95,As9s:0.95,AsTs:0.95,AsJc:0.95,AsJd:0.95,AsJh:0.95,AsJs:0.95,AsQs:0.95,AsKc:0.95,AsKd:0.95,AsKh:0.95,AsKs:0.95"},
			1000,
			map[string]float32{
				"KcAs": 0.855,
				"KsAc": 0.785,
			},
		},
	}

	for _, testCase := range table {
		board := types.ParseBoard(testCase.board)
		var ranges []types.Range
		for _, rangeStr := range testCase.ranges {
			range_ := types.ParseRange(rangeStr)
			range_.RemoveCards(board...)
			ranges = append(ranges, range_)
		}
		myRange := types.ParseRange(testCase.myRange)
		myRange.RemoveCards(board...)
		params := RequestParams{
			Board:      board,
			MyRange:    myRange,
			OppRanges:  ranges,
			Iterations: uint32(testCase.iterations),
			Timeout:    1,
		}
		result := CalculateEquity(&params)
		for hand, equity := range testCase.equity {
			calculatedEquity := float32(result.Equity[types.ParseHand(hand)])
			if math.Abs(float64(calculatedEquity-equity)) > 0.01 {
				t.Error("Equity not match", hand, calculatedEquity, "!=", equity)
			}
		}
	}
}

func BenchmarkCalculateEquity(b *testing.B) {
	table := []struct {
		board      string
		myRange    string
		ranges     []string
		iterations int
	}{
		{
			"7hAsQdJs8h",
			"9d9c:0.209,9h9c:0.212,9h9d:0.201,9s9c:0.209,9s9d:0.186,9s9h:0.201,AcJc:0.095,AcJd:0.095,AcJh:0.095,AcKc:0.044,AcKd:0.044,AcKh:0.056,AcKs:0.045,AcQc:0.045,AcQh:0.035,AcQs:0.046,AcTc:0.076,AcTd:0.077,AcTh:0.091,AcTs:0.077,AdAc:0.009,AdJc:0.095,AdJd:0.095,AdJh:0.095,AdKc:0.045,AdKd:0.045,AdKh:0.056,AdKs:0.045,AdQc:0.045,AdQh:0.035,AdQs:0.046,AdTc:0.076,AdTd:0.077,AdTh:0.091,AdTs:0.077,AhAc:0.015,AhAd:0.015,AhJc:0.1,AhJd:0.1,AhJh:0.2,AhKc:0.06,AhKd:0.06,AhKh:0.1,AhKs:0.06,AhQc:0.059,AhQh:0.097,AhQs:0.06,AhTc:0.104,AhTd:0.105,AhTh:0.1,AhTs:0.105,AsAc:0.009,AsAd:0.009,AsAh:0.015,AsJc:0.095,AsJd:0.095,AsJh:0.095,AsKc:0.045,AsKd:0.045,AsKh:0.056,AsKs:0.045,AsQc:0.045,AsQh:0.035,AsQs:0.046,AsTc:0.076,AsTd:0.077,AsTh:0.091,AsTs:0.077,JdJc:0.086,JhJc:0.068,JhJd:0.068,KcQc:0.04,KcQh:0.031,KcQs:0.04,KdKc:0.002,KdQc:0.04,KdQh:0.031,KdQs:0.04,KhKc:0.012,KhKd:0.012,KhQc:0.048,KhQh:0.08,KhQs:0.051,KsKc:0.002,KsKd:0.002,KsKh:0.012,KsQc:0.04,KsQh:0.031,KsQs:0.04,QhQc:0.065,Qs5c:0.274,QsQc:0.094,QsQh:0.063,TdTc:0.261,ThTc:0.301,ThTd:0.3,TsTc:0.261,TsTd:0.259,TsTh:0.301",
			[]string{"3h2h:0.619,4h2h:0.693,4h3h:0.75,5h2h:0.883,5h3h:0.98,5h4h:1,6c5c:0.098,6c5d:0.203,6c5h:0.204,6c5s:0.206,6d5c:0.207,6d5d:0.102,6d5h:0.205,6d5s:0.207,6h2h:0.796,6h3h:0.869,6h4h:1,6h5c:0.396,6h5d:0.204,6h5s:0.207,6s5c:0.206,6s5d:0.203,6s5h:0.204,6s5s:0.103,7c2c:1,7c3c:0.792,7c4c:0.432,7c5c:0.487,7c6h:0.672,7d2d:1,7d3d:0.797,7d4d:0.435,7d5c:0.489,7d6h:0.674,7s2s:1,7s3s:0.802,7s4s:0.435,7s5c:0.491,7s6h:0.675,8c2c:0.8,8c3c:0.8,8c4c:0.444,8c5c:0.513,8c6c:0.003,8c6d:0.008,8c6h:0.659,8c6s:0.004,8c7c:0.113,8c7d:0.232,8c7s:0.234,8d2d:0.8,8d3d:0.8,8d4d:0.448,8d6c:0.015,8d6d:0.009,8d6h:0.666,8d6s:0.014,8d7c:0.224,8d7d:0.114,8d7s:0.231,8s2s:0.8,8s3s:0.8,8s4s:0.445,8s6c:0.008,8s6d:0.011,8s6h:0.661,8s6s:0.003,8s7c:0.218,8s7d:0.224,8s7s:0.113,9c5c:0.2,9c6c:0.316,9c6d:0.326,9c6h:0.342,9c6s:0.325,9c7c:0.35,9c7d:0.694,9c7s:0.696,9c8c:0.369,9c8d:0.363,9c8s:0.363,9d5d:0.411,9d6c:0.291,9d6d:0.299,9d6h:0.336,9d6s:0.299,9d7c:0.658,9d7d:0.337,9d7s:0.665,9d8c:0.365,9d8d:0.359,9d8s:0.358,9h2h:0.548,9h3h:0.598,9h4h:0.792,9h5h:0.8,9h6c:0.403,9h6d:0.405,9h6h:0.794,9h6s:0.405,9h7c:0.693,9h7d:0.699,9h7s:0.7,9h8c:0.383,9h8d:0.378,9h8s:0.375,9s5s:0.412,9s6c:0.291,9s6d:0.299,9s6h:0.336,9s6s:0.299,9s7c:0.661,9s7d:0.666,9s7s:0.338,9s8c:0.364,9s8d:0.358,9s8s:0.358,Ah2h:0.232,Ah3h:0.238,Ah4h:0.254,Ah5h:0.279,Ah6h:0.276,Ah7c:0.064,Ah7d:0.067,Ah7s:0.067,Ah8c:0.11,Ah8d:0.112,Ah8s:0.111,Jc2c:0.3,Jc3c:0.3,Jc3d:0.3,Jc3h:0.3,Jc3s:0.3,Jc4c:0.3,Jc4d:0.3,Jc4h:0.3,Jc4s:0.3,Jc5c:0.061,Jc5d:0.3,Jc5h:0.197,Jc5s:0.3,Jc6c:0.3,Jc6d:0.3,Jc6h:0.225,Jc6s:0.3,Jd2d:0.3,Jd3c:0.3,Jd3d:0.3,Jd3h:0.3,Jd3s:0.3,Jd4c:0.3,Jd4d:0.3,Jd4h:0.3,Jd4s:0.3,Jd5c:0.062,Jd5d:0.3,Jd5h:0.199,Jd5s:0.3,Jd6c:0.3,Jd6d:0.3,Jd6h:0.227,Jd6s:0.3,Jh3c:0.246,Jh3d:0.247,Jh3s:0.249,Jh4c:0.288,Jh4d:0.289,Jh4s:0.288,Jh5c:0.02,Jh5d:0.154,Jh5s:0.155,Jh6c:0.179,Jh6d:0.176,Jh6s:0.174,Kh4h:0.029,Kh5h:0.155,Kh6h:0.114,Qc2c:0.051,Qc2d:0.05,Qc2h:0.088,Qc2s:0.062,Qc3c:0.025,Qc3d:0.025,Qc3s:0.027,Qc4c:0.038,Qc4d:0.038,Qc4h:0.004,Qc4s:0.037,Qc5d:0.022,Qc5s:0.024,Qc6c:0.006,Qc6d:0.005,Qc6s:0.003,Qh2c:0.101,Qh2d:0.1,Qh2s:0.11,Qh3c:0.09,Qh3d:0.09,Qh3s:0.091,Qh4c:0.09,Qh4d:0.091,Qh4s:0.09,Qh5d:0.083,Qh5s:0.083,Qh6c:0.075,Qh6d:0.072,Qh6s:0.071,Qs2c:0.033,Qs2d:0.032,Qs2h:0.07,Qs2s:0.047,Qs3c:0.005,Qs3d:0.005,Qs3s:0.008,Qs4c:0.034,Qs4d:0.034,Qs4s:0.033,Qs5d:0.001,Qs5s:0.004,Qs6c:0.038,Qs6d:0.034,Qs6s:0.033,Tc7c:0.327,Tc7d:0.645,Tc7s:0.642,Tc8c:0.315,Tc8d:0.312,Tc8s:0.311,Td7c:0.64,Td7d:0.331,Td7s:0.645,Td8c:0.315,Td8d:0.311,Td8s:0.31,Th2h:0.619,Th3h:0.646,Th4h:0.737,Th5h:0.8,Th6h:0.8,Th7c:0.609,Th7d:0.611,Th7s:0.613,Th8c:0.348,Th8d:0.342,Th8s:0.343,Ts7c:0.638,Ts7d:0.647,Ts7s:0.33,Ts8c:0.315,Ts8d:0.311,Ts8s:0.31"},
			1000,
		},
	}

	for _, testCase := range table {
		board := types.ParseBoard(testCase.board)
		var ranges []types.Range
		for _, rangeStr := range testCase.ranges {
			r := types.ParseRange(rangeStr)
			r.RemoveCards(board...)
			ranges = append(ranges, r)
		}
		myRange := types.ParseRange(testCase.myRange)
		myRange.RemoveCards(board...)
		params := RequestParams{
			Board:      board,
			MyRange:    myRange,
			OppRanges:  ranges,
			Iterations: uint32(testCase.iterations),
			Timeout:    1,
		}
		for i := 0; i < b.N; i++ {
			CalculateEquity(&params)
		}
	}
}
