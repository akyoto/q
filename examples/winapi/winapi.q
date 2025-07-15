main() {
	title := "Title\0"
	message := "Message\0"
	user32.MessageBoxA(0, message.ptr, title.ptr, 0)
}

extern {
	user32 {
		MessageBoxA(window *any, text *byte, title *byte, flags uint) -> (button int)
	}
}