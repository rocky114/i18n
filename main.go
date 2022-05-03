package main

import (
	"awesomeProject/i18n"
	"fmt"
)

func main()  {
	i18n.New("./i18n/lang").Generate()

	fmt.Println(i18n.Translate("zh", i18n.Address))
}
