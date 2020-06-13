#!/usr/bin/osascript
tell application "iTerm2"
    tell current session of current tab of current window
        write text "cd ~/go/src/github.com/kyleu/npn"
        write text "npn"
        split vertically with default profile
    end tell
    tell second session of current tab of current window
        write text "cd ~/go/src/github.com/kyleu/npn/client"
        write text "../bin/build-client-watch.sh"
        split horizontally with default profile
    end tell
    tell third session of current tab of current window
        write text "cd ~/go/src/github.com/kyleu/npn"
        write text "bin/build-css-watch.sh"
    end tell
end tell
