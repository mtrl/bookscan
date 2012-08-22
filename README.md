bookscan
========

My book scan repo for use with gphoto2

Connect the cameras. Unmount them. Run gphoto2 --auto-detect to get the port names

cd to the root of the repo then run the automatic photo taker and image rotator like this:

Set the camera up with gphoto2 --port [port] --config
- Macro
- Manual
- Focuslock?

The new "go" script launches both camera capuring scripts at once with

"go [left camera port] [start page] [right camera port] [start page] [interval]

Or run one at a time with:

"capture [left/right] [port name] [starting page number] [interval between loops]

e.g. capture right usb:001,005 1 1
