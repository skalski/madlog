# madlog
summarizes stacktraces and error-logs from all running Docker containers.


## how to use

just start `madlog` on linux or `madlog.exe` on Windows.

Use `-h` to see help.

You can set `-level=1` if you only want to see message that contains errors and stacktraces.


## how to build

linux: `gcc madlog.c -o madlog && ./madlog`

linux crosscompile to Win 64: `x86_64-w64-mingw32-gcc madlog.c -o madlog.exe`



