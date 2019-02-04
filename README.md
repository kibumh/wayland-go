# What is this?
Yet another, just for fun, wayland library for go.

# Directory Structure
 directory | description
----|-----
wire | implements [wire format of wayland protocol](https://wayland.freedesktop.org/docs/html/ch04.html).
scanner | a program that generates go source files from  wayland's XML files.
client/protocol, server/protocol | The generated go source files live here.
client, server | implements client or server API based on the protocol package.

# End goal
- explorers several edges of the world of wayland. (screen shot tools, AutoHotKey-like automators, etc)
