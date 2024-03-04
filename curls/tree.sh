#!/bin/bash
export PATH="/mingw64/bin:/usr/local/bin:/usr/bin:/bin:$USERPROFILE/Downloads"

function get_curriculum_tree () {
    /C/Windows/System32/curl.exe "https://localhost/api/curriculum-tree"  -b cookie.txt --insecure
}

get_curriculum_tree