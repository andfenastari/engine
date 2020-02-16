package main

const (
	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 fragColor;
		void main() {
			fragColor = vec4(1, 1, 1, 1);
		}
	` + "\x00"
)
