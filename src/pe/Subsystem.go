package pe

type Subsystem uint16

const (
	IMAGE_SUBSYSTEM_UNKNOWN                  Subsystem = 0
	IMAGE_SUBSYSTEM_NATIVE                   Subsystem = 1
	IMAGE_SUBSYSTEM_WINDOWS_GUI              Subsystem = 2
	IMAGE_SUBSYSTEM_WINDOWS_CUI              Subsystem = 3
	IMAGE_SUBSYSTEM_OS2_CUI                  Subsystem = 5
	IMAGE_SUBSYSTEM_POSIX_CUI                Subsystem = 7
	IMAGE_SUBSYSTEM_NATIVE_WINDOWS           Subsystem = 8
	IMAGE_SUBSYSTEM_WINDOWS_CE_GUI           Subsystem = 9
	IMAGE_SUBSYSTEM_EFI_APPLICATION          Subsystem = 10
	IMAGE_SUBSYSTEM_EFI_BOOT_SERVICE_DRIVER  Subsystem = 11
	IMAGE_SUBSYSTEM_EFI_RUNTIME_DRIVER       Subsystem = 12
	IMAGE_SUBSYSTEM_EFI_ROM                  Subsystem = 13
	IMAGE_SUBSYSTEM_XBOX                     Subsystem = 14
	IMAGE_SUBSYSTEM_WINDOWS_BOOT_APPLICATION Subsystem = 16
)