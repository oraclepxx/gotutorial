package main;

import "fmt";

func main() {
	for i := 1; i <= 4; i++ {
		var num = i * 100;
		for j := 1; j <= 4; j++ {
			if (j == i) {
				continue;
			}
			var num2 = num + j * 10;
			for k := 1; k <= 4; k++ {
				if (k == i || k == j) {
					continue;
				}
				var num3 = num2 + k;
				fmt.Println(num3);
			}
		}
	}
}
