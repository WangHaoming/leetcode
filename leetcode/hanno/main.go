package hano

func hano(tray1, tray2, tray3 string, layers int) {
	// if layers == 1 {
	// 	println(tray1, " -> ", tray3, ":", layers)
	// 	return
	// }
	if layers > 0 {
		hano(tray1, tray3, tray2, layers-1)
		println(tray1, " -> ", tray3, ":", layers)
		hano(tray2, tray1, tray3, layers-1)
	}

}

func Test() {
	k := 4
	hano("A", "B", "C", k)
}
