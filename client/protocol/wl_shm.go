/*
 * THIS CODE IS GENERATED. DO NOT EDIT BY HANDS.
 */

package protocol

import (
	"fmt"

	"github.com/kibumh/wayland-go/wire"
)

/*
 * ENUMS
 */

// WlShmError is for wl_shm error values.
//
// These errors can be emitted in response to wl_shm requests.
type WlShmError int

const (
	WlShmErrorInvalidFormat WlShmError = 0 // buffer format is not known
	WlShmErrorInvalidStride WlShmError = 1 // invalid size or stride during pool or buffer creation
	WlShmErrorInvalidFd     WlShmError = 2 // mmapping the file descriptor failed
)

// WlShmFormat is for pixel formats.
//
// This describes the memory layout of an individual pixel.
//
// All renderers should support argb8888 and xrgb8888 but any other
// formats are optional and may not be supported by the particular
// renderer in use.
//
// The drm format codes match the macros defined in drm_fourcc.h.
// The formats actually supported by the compositor will be
// reported by the format event.
type WlShmFormat int

const (
	WlShmFormatArgb8888    WlShmFormat = 0          // 32-bit ARGB format, [31:0] A:R:G:B 8:8:8:8 little endian
	WlShmFormatXrgb8888    WlShmFormat = 1          // 32-bit RGB format, [31:0] x:R:G:B 8:8:8:8 little endian
	WlShmFormatC8          WlShmFormat = 0x20203843 // 8-bit color index format, [7:0] C
	WlShmFormatRgb332      WlShmFormat = 0x38424752 // 8-bit RGB format, [7:0] R:G:B 3:3:2
	WlShmFormatBgr233      WlShmFormat = 0x38524742 // 8-bit BGR format, [7:0] B:G:R 2:3:3
	WlShmFormatXrgb4444    WlShmFormat = 0x32315258 // 16-bit xRGB format, [15:0] x:R:G:B 4:4:4:4 little endian
	WlShmFormatXbgr4444    WlShmFormat = 0x32314258 // 16-bit xBGR format, [15:0] x:B:G:R 4:4:4:4 little endian
	WlShmFormatRgbx4444    WlShmFormat = 0x32315852 // 16-bit RGBx format, [15:0] R:G:B:x 4:4:4:4 little endian
	WlShmFormatBgrx4444    WlShmFormat = 0x32315842 // 16-bit BGRx format, [15:0] B:G:R:x 4:4:4:4 little endian
	WlShmFormatArgb4444    WlShmFormat = 0x32315241 // 16-bit ARGB format, [15:0] A:R:G:B 4:4:4:4 little endian
	WlShmFormatAbgr4444    WlShmFormat = 0x32314241 // 16-bit ABGR format, [15:0] A:B:G:R 4:4:4:4 little endian
	WlShmFormatRgba4444    WlShmFormat = 0x32314152 // 16-bit RBGA format, [15:0] R:G:B:A 4:4:4:4 little endian
	WlShmFormatBgra4444    WlShmFormat = 0x32314142 // 16-bit BGRA format, [15:0] B:G:R:A 4:4:4:4 little endian
	WlShmFormatXrgb1555    WlShmFormat = 0x35315258 // 16-bit xRGB format, [15:0] x:R:G:B 1:5:5:5 little endian
	WlShmFormatXbgr1555    WlShmFormat = 0x35314258 // 16-bit xBGR 1555 format, [15:0] x:B:G:R 1:5:5:5 little endian
	WlShmFormatRgbx5551    WlShmFormat = 0x35315852 // 16-bit RGBx 5551 format, [15:0] R:G:B:x 5:5:5:1 little endian
	WlShmFormatBgrx5551    WlShmFormat = 0x35315842 // 16-bit BGRx 5551 format, [15:0] B:G:R:x 5:5:5:1 little endian
	WlShmFormatArgb1555    WlShmFormat = 0x35315241 // 16-bit ARGB 1555 format, [15:0] A:R:G:B 1:5:5:5 little endian
	WlShmFormatAbgr1555    WlShmFormat = 0x35314241 // 16-bit ABGR 1555 format, [15:0] A:B:G:R 1:5:5:5 little endian
	WlShmFormatRgba5551    WlShmFormat = 0x35314152 // 16-bit RGBA 5551 format, [15:0] R:G:B:A 5:5:5:1 little endian
	WlShmFormatBgra5551    WlShmFormat = 0x35314142 // 16-bit BGRA 5551 format, [15:0] B:G:R:A 5:5:5:1 little endian
	WlShmFormatRgb565      WlShmFormat = 0x36314752 // 16-bit RGB 565 format, [15:0] R:G:B 5:6:5 little endian
	WlShmFormatBgr565      WlShmFormat = 0x36314742 // 16-bit BGR 565 format, [15:0] B:G:R 5:6:5 little endian
	WlShmFormatRgb888      WlShmFormat = 0x34324752 // 24-bit RGB format, [23:0] R:G:B little endian
	WlShmFormatBgr888      WlShmFormat = 0x34324742 // 24-bit BGR format, [23:0] B:G:R little endian
	WlShmFormatXbgr8888    WlShmFormat = 0x34324258 // 32-bit xBGR format, [31:0] x:B:G:R 8:8:8:8 little endian
	WlShmFormatRgbx8888    WlShmFormat = 0x34325852 // 32-bit RGBx format, [31:0] R:G:B:x 8:8:8:8 little endian
	WlShmFormatBgrx8888    WlShmFormat = 0x34325842 // 32-bit BGRx format, [31:0] B:G:R:x 8:8:8:8 little endian
	WlShmFormatAbgr8888    WlShmFormat = 0x34324241 // 32-bit ABGR format, [31:0] A:B:G:R 8:8:8:8 little endian
	WlShmFormatRgba8888    WlShmFormat = 0x34324152 // 32-bit RGBA format, [31:0] R:G:B:A 8:8:8:8 little endian
	WlShmFormatBgra8888    WlShmFormat = 0x34324142 // 32-bit BGRA format, [31:0] B:G:R:A 8:8:8:8 little endian
	WlShmFormatXrgb2101010 WlShmFormat = 0x30335258 // 32-bit xRGB format, [31:0] x:R:G:B 2:10:10:10 little endian
	WlShmFormatXbgr2101010 WlShmFormat = 0x30334258 // 32-bit xBGR format, [31:0] x:B:G:R 2:10:10:10 little endian
	WlShmFormatRgbx1010102 WlShmFormat = 0x30335852 // 32-bit RGBx format, [31:0] R:G:B:x 10:10:10:2 little endian
	WlShmFormatBgrx1010102 WlShmFormat = 0x30335842 // 32-bit BGRx format, [31:0] B:G:R:x 10:10:10:2 little endian
	WlShmFormatArgb2101010 WlShmFormat = 0x30335241 // 32-bit ARGB format, [31:0] A:R:G:B 2:10:10:10 little endian
	WlShmFormatAbgr2101010 WlShmFormat = 0x30334241 // 32-bit ABGR format, [31:0] A:B:G:R 2:10:10:10 little endian
	WlShmFormatRgba1010102 WlShmFormat = 0x30334152 // 32-bit RGBA format, [31:0] R:G:B:A 10:10:10:2 little endian
	WlShmFormatBgra1010102 WlShmFormat = 0x30334142 // 32-bit BGRA format, [31:0] B:G:R:A 10:10:10:2 little endian
	WlShmFormatYuyv        WlShmFormat = 0x56595559 // packed YCbCr format, [31:0] Cr0:Y1:Cb0:Y0 8:8:8:8 little endian
	WlShmFormatYvyu        WlShmFormat = 0x55595659 // packed YCbCr format, [31:0] Cb0:Y1:Cr0:Y0 8:8:8:8 little endian
	WlShmFormatUyvy        WlShmFormat = 0x59565955 // packed YCbCr format, [31:0] Y1:Cr0:Y0:Cb0 8:8:8:8 little endian
	WlShmFormatVyuy        WlShmFormat = 0x59555956 // packed YCbCr format, [31:0] Y1:Cb0:Y0:Cr0 8:8:8:8 little endian
	WlShmFormatAyuv        WlShmFormat = 0x56555941 // packed AYCbCr format, [31:0] A:Y:Cb:Cr 8:8:8:8 little endian
	WlShmFormatNv12        WlShmFormat = 0x3231564e // 2 plane YCbCr Cr:Cb format, 2x2 subsampled Cr:Cb plane
	WlShmFormatNv21        WlShmFormat = 0x3132564e // 2 plane YCbCr Cb:Cr format, 2x2 subsampled Cb:Cr plane
	WlShmFormatNv16        WlShmFormat = 0x3631564e // 2 plane YCbCr Cr:Cb format, 2x1 subsampled Cr:Cb plane
	WlShmFormatNv61        WlShmFormat = 0x3136564e // 2 plane YCbCr Cb:Cr format, 2x1 subsampled Cb:Cr plane
	WlShmFormatYuv410      WlShmFormat = 0x39565559 // 3 plane YCbCr format, 4x4 subsampled Cb (1) and Cr (2) planes
	WlShmFormatYvu410      WlShmFormat = 0x39555659 // 3 plane YCbCr format, 4x4 subsampled Cr (1) and Cb (2) planes
	WlShmFormatYuv411      WlShmFormat = 0x31315559 // 3 plane YCbCr format, 4x1 subsampled Cb (1) and Cr (2) planes
	WlShmFormatYvu411      WlShmFormat = 0x31315659 // 3 plane YCbCr format, 4x1 subsampled Cr (1) and Cb (2) planes
	WlShmFormatYuv420      WlShmFormat = 0x32315559 // 3 plane YCbCr format, 2x2 subsampled Cb (1) and Cr (2) planes
	WlShmFormatYvu420      WlShmFormat = 0x32315659 // 3 plane YCbCr format, 2x2 subsampled Cr (1) and Cb (2) planes
	WlShmFormatYuv422      WlShmFormat = 0x36315559 // 3 plane YCbCr format, 2x1 subsampled Cb (1) and Cr (2) planes
	WlShmFormatYvu422      WlShmFormat = 0x36315659 // 3 plane YCbCr format, 2x1 subsampled Cr (1) and Cb (2) planes
	WlShmFormatYuv444      WlShmFormat = 0x34325559 // 3 plane YCbCr format, non-subsampled Cb (1) and Cr (2) planes
	WlShmFormatYvu444      WlShmFormat = 0x34325659 // 3 plane YCbCr format, non-subsampled Cr (1) and Cb (2) planes
)

/*
 * EVENT HANDLER TYPES
 */

// WlShmFormatHandler is a handler for pixel format description.
//
// Informs the client about a valid pixel format that
// can be used for buffers. Known formats include
// argb8888 and xrgb8888.
type WlShmFormatHandler func(format wire.Uint) error

/*
 * TYPE
 */
// WlShm is shared memory support.
//
// A singleton global object that provides support for shared
// memory.
//
// Clients can create wl_shm_pool objects using the create_pool
// request.
//
// At connection setup time, the wl_shm object emits one or more
// format events to inform clients about the valid pixel formats
// that can be used for buffers.
type WlShm struct {
	Base

	formatHandler WlShmFormatHandler
}

// NewWlShm creates a WlShm object.
func NewWlShm(c *wire.Conn) *WlShm {
	return NewWlShmWithID(c, c.NewID())
}

// NewWlShmWithID creates a WlShm object with a given id.
func NewWlShmWithID(c *wire.Conn, id wire.ID) *WlShm {
	o := &WlShm{Base: Base{c, id}}
	c.RegisterObject(o)
	return o
}

/*
 * REQUESTS
 */

// CreatePool is for create a shm pool
//
// Create a new wl_shm_pool object.
//
// The pool can be used to create shared memory based buffer
// objects.  The server will mmap size bytes of the passed file
// descriptor, to use as backing memory for the pool.
func (o *WlShm) CreatePool(id *WlShmPool, fd wire.FD, size wire.Int) error {
	msg, err := wire.NewMessage(o.ID(), 0)
	if err != nil {
		return err
	}

	if err = msg.Write(id.ID()); err != nil {
		return err
	}

	if err = msg.Write(fd); err != nil {
		return err
	}

	if err = msg.Write(size); err != nil {
		return err
	}

	if err = o.Base.Conn.Write(msg); err != nil {
		return err
	}

	return nil
}

/*
 * EVENTS
 */

// ServeFormat is for pixel format description.
//
// Informs the client about a valid pixel format that
// can be used for buffers. Known formats include
// argb8888 and xrgb8888.
func (o *WlShm) ServeFormat(format wire.Uint) error {
	if o.formatHandler == nil {
		return nil
	}
	return o.formatHandler(format)
}

// HandleFormat registers a handler for a Format event.
func (o *WlShm) HandleFormat(h WlShmFormatHandler) {
	o.formatHandler = h
}

// ServeMessage is a multiplexer for a message.
func (o *WlShm) ServeMessage(msg *wire.Message) error {
	switch msg.Opcode {
	case 0:
		if o.formatHandler == nil {
			return nil
		}

		r := wire.NewReader(msg.Payload)

		format, err := r.ReadUint()
		if err != nil {
			return err
		}

		return o.formatHandler(format)

	default:
		return fmt.Errorf("WlShm: unhandled message(%v)", msg)
	}
}
