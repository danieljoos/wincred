package wincred

import (
	"syscall"
	"unsafe"
	"fmt"
)

var (
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")

	procCredRead   = modadvapi32.NewProc("CredReadW")
	procCredWrite  = modadvapi32.NewProc("CredWriteW")
	procCredDelete = modadvapi32.NewProc("CredDeleteW")
	procCredFree   = modadvapi32.NewProc("CredFree")
	procCredList   = modadvapi32.NewProc("CredEnumerateW")
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa374788(v=vs.85).aspx
type nativeCREDENTIAL struct {
	Flags              uint32
	Type               uint32
	TargetName         *uint16
	Comment            *uint16
	LastWritten        syscall.Filetime
	CredentialBlobSize uint32
	CredentialBlob     uintptr
	Persist            uint32
	AttributeCount     uint32
	Attributes         uintptr
	TargetAlias        *uint16
	UserName           *uint16
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa374790(v=vs.85).aspx
type nativeCREDENTIAL_ATTRIBUTE struct {
	Keyword   *uint16
	Flags     uint32
	ValueSize uint32
	Value     uintptr
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa374788(v=vs.85).aspx
type nativeCRED_TYPE uint32

const (
	naCRED_TYPE_GENERIC                 nativeCRED_TYPE = 0x1
	naCRED_TYPE_DOMAIN_PASSWORD         nativeCRED_TYPE = 0x2
	naCRED_TYPE_DOMAIN_CERTIFICATE      nativeCRED_TYPE = 0x3
	naCRED_TYPE_DOMAIN_VISIBLE_PASSWORD nativeCRED_TYPE = 0x4
	naCRED_TYPE_GENERIC_CERTIFICATE     nativeCRED_TYPE = 0x5
	naCRED_TYPE_DOMAIN_EXTENDED         nativeCRED_TYPE = 0x6
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa374804(v=vs.85).aspx
func nativeCredRead(targetName string, typ nativeCRED_TYPE) (*Credential, error) {
	var pcred uintptr
	targetNamePtr, _ := syscall.UTF16PtrFromString(targetName)
	ret, _, err := procCredRead.Call(
		uintptr(unsafe.Pointer(targetNamePtr)),
		uintptr(typ),
		0,
		uintptr(unsafe.Pointer(&pcred)),
	)
	if ret == 0 {
		return nil, err
	}
	defer procCredFree.Call(pcred)

	return nativeToCredential((*nativeCREDENTIAL)(unsafe.Pointer(pcred))), nil
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa375187(v=vs.85).aspx
func nativeCredWrite(cred *Credential, typ nativeCRED_TYPE) error {
	ncred := nativeFromCredential(cred)
	ncred.Type = uint32(typ)
	ret, _, err := procCredWrite.Call(
		uintptr(unsafe.Pointer(ncred)),
		0,
	)
	if ret == 0 {
		return err
	}

	return nil
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa374787(v=vs.85).aspx
func nativeCredDelete(cred *Credential, typ nativeCRED_TYPE) error {
	targetNamePtr, _ := syscall.UTF16PtrFromString(cred.TargetName)
	ret, _, err := procCredDelete.Call(
		uintptr(unsafe.Pointer(targetNamePtr)),
		uintptr(typ),
		0,
	)
	if ret == 0 {
		return err
	}

	return nil
}

// https://msdn.microsoft.com/en-us/library/windows/desktop/aa374794(v=vs.85).aspx
func nativeCredList() error {
	fmt.Println("in listing function-----------")
	var count uint
	var credList []*nativeCREDENTIAL
	ret, _, err := procCredList.Call(
		uintptr(0),
		uintptr(0),
		uintptr(unsafe.Pointer(&count)),
		uintptr(unsafe.Pointer(&credList[0])),
	)
	fmt.Println(ret)
	fmt.Println(err)
	fmt.Println("Number of items in the keychain:")
	fmt.Println(count)
	fmt.Println("Keychain items:")
	fmt.Println("This is a uintptr- an integer type that is large enough to hold the bit pattern of any pointer:")
	fmt.Println(uintptr(unsafe.Pointer(&credList)))
	fmt.Println("this should be list of pointers to credentials")
	fmt.Println(unsafe.Pointer(&credList))
	fmt.Println(credList)
	fmt.Println(utf16PtrToString(((*nativeCREDENTIAL)(unsafe.Pointer(&credList[0]))).UserName))
	// Trying in cgo
	return nil
}