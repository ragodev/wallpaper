package wallpaper

import (
	"syscall"
	"unsafe"
)

const (
	// Windows constants
	SPI_SETDESKWALLPAPER uintptr = 0x0014
	WM_WININICHANGE      uintptr = 0x001A

	LibraryName  string = "user32.dll"
	FunctionName string = "SystemParametersInfoW"
)

func SetWallpaper(path string) error {
	user32, _ := syscall.LoadLibrary(LibraryName)
	systemParameters, _ := syscall.GetProcAddress(user32, FunctionName)
	defer syscall.FreeLibrary(user32)

	ret, _, err := syscall.Syscall9(uintptr(systemParameters), 4, SPI_SETDESKWALLPAPER, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))), WM_WININICHANGE, 0, 0, 0, 0, 0)

	if ret == 0 {
		return err
	}

	return nil
}
