package main

import "fmt"

func main() {
	// 1. Underlying array တစ်ခု ဖန်တီးခြင်း
	numbersArray := [6]int{10, 20, 30, 40, 50, 60}
	fmt.Println("Underlying Array:", numbersArray)

	// 2. Array မှ slice တစ်ခု ဖန်တီးခြင်း (index 2 မှ 4 မတိုင်ခင်အထိ)
	mySlice := numbersArray[2:4] // [30, 40]
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", mySlice, len(mySlice),cap(mySlice))

	// Capacity: index 2 မှ array အဆုံးထိရေတွက်မည် (30, 40, 50, 60) -> 4 ခု
	// 3. Append within Capacity
	mySlice = append(mySlice, 70)
	fmt.Println("\n--- After appending 70 (within capacity) ---")
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", mySlice, len(mySlice), cap(mySlice))

	// Underlying array ပါ ပြောင်းလဲသွားသည်ကို သတိပြုပါ
	fmt.Println("Underlying Array is now:", numbersArray) // Output: [10 20 30 40 70 60]

	// 4. Append exceeding Capacity
	// လက်ရှိ cap = 4, len = 3 ဖြစ်နေသေးသည်
	mySlice = append(mySlice, 80) // cap ပြည့်သွားပြီ
	mySlice = append(mySlice, 90) // cap ကျော်လွန်သွားပြီ
	fmt.Println("\n--- After appending 80 and 90 (exceeding capacity) ---")

	// Go က array အသစ်တစ်ခု (size 8) ကို တည်‌ဆောက်ပြီး data တွေ copy ကူးထည့်ပါမည်။
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", mySlice, len(mySlice), cap(mySlice))

	// array အသစ်ကို ညွှန်းဆိုထားသောကြောင့် မူလ array ကို ထိခိုက်တော့မည် မဟုတ်ပါ။
	fmt.Println("Underlying Array remains unchanged:", numbersArray)
}