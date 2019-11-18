package main

func main(){
	println(season(13))
}

func season(m int)(s string){
	switch m {
	case 12,1,2:
		return "冬季"
	case 3,4,5:
		return "春季"
	case 6,7,8:
		return "夏季"
	case 9,10,11:
		return "秋季"
	default:
		return "???"
	}
}