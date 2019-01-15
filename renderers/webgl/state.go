package webgl

import "github.com/tokkenno/seed/renderers/webgl/buffer"

type State struct {
	colorBuffer   *buffer.Color
	depthBuffer   *buffer.Depth
	stencilBuffer *buffer.Stencil
}

// https://github.com/mrdoob/three.js/blob/dev/src/renderers/webgl/WebGLState.js#L315