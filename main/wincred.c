#include <windows.h>
#include <stdio.h>
#include <io.h>

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