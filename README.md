wincred
=======

Go wrapper around the Windows Credential Manager API functions.

![Go](https://github.com/danieljoos/wincred/workflows/Go/badge.svg)
[![GoDoc](https://godoc.org/github.com/danieljoos/wincred?status.svg)](https://godoc.org/github.com/danieljoos/wincred)


Installation
------------

```Go
go get github.com/danieljoos/wincred
```


Usage
-----

See the following examples:

### Create and store a new generic credential object
```Go
package main

import (
    "fmt"
    "github.com/danieljoos/wincred"
)

func main() {
    cred := wincred.NewGenericCredential("myGoApplication")
    cred.CredentialBlob = []byte("my secret")
    err := cred.Write()
    
    if err != nil {
        fmt.Println(err)
    }
} 
```

### Retrieve a credential object
```Go
package main

import (
    "fmt"
    "github.com/danieljoos/wincred"
)

func main() {
    cred, err := wincred.GetGenericCredential("myGoApplication")
    if err == nil {
        fmt.Println(string(cred.CredentialBlob))
    }
} 
```

### Remove a credential object
```Go
package main

import (
    "fmt"
    "github.com/danieljoos/wincred"
)

func main() {
    cred, err := wincred.GetGenericCredential("myGoApplication")
    if err != nil {
        fmt.Println(err)
        return
    }
    cred.Delete()
} 
```

### List all available credentials
```Go
package main

import (
    "fmt"
    "github.com/danieljoos/wincred"
)

func main() {
    creds, err := wincred.List()
    if err != nil {
        fmt.Println(err)
        return
    }
    for i := range(creds) {
        fmt.Println(creds[i].TargetName)
    }
}
```

Hints
-----

### Encoding

The credential objects simply store byte arrays without specific meaning or encoding.
For sharing between different applications, it might make sense to apply an explicit string encoding - for example **UTF-16 LE** (used nearly everywhere in the Win32 API).

```Go
package main

import (
	"fmt"
	"os"

	"github.com/danieljoos/wincred"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	cred := wincred.NewGenericCredential("myGoApplication")

	encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	blob, _, err := transform.Bytes(encoder, []byte("mysecret"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cred.CredentialBlob = blob
	err = cred.Write()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

```

### Limitations

The size of a credential blob is limited to **2560 Bytes** by the Windows API.
