
# LOGIN:
## GET http://192.168.100.1/VmLogin.html
 - `$("#password").value( password )`
 - `#password` csrf prevention via name changing.
 No user field.

## POST http://192.168.100.1/cgi-bin/VmLoginCgi
 - DATA {`#password.name`}=password
 - Response contains javascript validation.
    - `var res="0";`
        - if 1 wrong password, show dialog.
        - if 3 locked out, show dialog.
    - `var defpass="0";`
        - if 1 using default password, naughty, show dialog.
    - `var singleUser="0";`
        - if 1 means someone else is logged in, show dialog.
    - `var lanAccess="1";`
        - 1 if on local network, else 0 and show dialog.

 - Log in when `res=0`, `defpass=0`, `singleUser=0`.
 - When DefaultPasswordDialog is closed, redirectTo called.
 - Log in is `redirectTo()` just redirects to "../home.html"
 - No cookies or tokens, session is per ip.
 - Settings can be adjusted at http://192.168.100.1/VmRgUserInterface.html