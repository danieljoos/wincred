#include <windows.h>
#include <stdio.h>
#include <io.h>
#include <fcntl.h>

typedef struct _CREDENTIALW {
	DWORD                  Flags;
	DWORD                  Type;
	LPWSTR                 TargetName;
	LPWSTR                 Comment;
	FILETIME               LastWritten;
	DWORD                  CredentialBlobSize;
	LPBYTE                 CredentialBlob;
	DWORD                  Persist;
	DWORD                  AttributeCount;
	PCREDENTIAL_ATTRIBUTEW Attributes;
	LPWSTR                 TargetAlias;
	LPWSTR                 UserName;
} CREDENTIALW, *PCREDENTIALW;

static void get_credential(void)
{
	CREDENTIALW **creds;
	DWORD num_creds;
	int i;

	if (!CredEnumerateW(NULL, 0, &num_creds, &creds))
		return;

	/* search for the first credential that matches username */
	for (i = 0; i < num_creds; ++i)
        printf("%s\n", creds[i]->UserName);

	CredFree(creds);
}

int main() {
    get_credential();
}