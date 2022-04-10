"use strict";
if (WebAssembly) {
  if (!WebAssembly.instantiateStreaming) {
    WebAssembly.instantiateStreaming = async (res, io) => {
      const src = await (await res).arrayBuffer();
      return await WebAssembly.instantiate(src, io)
    };
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((rt) => {
      go.run(rt.instance);
    }).catch((error) => console.error(error));
  }
} else {
  console.error("WebAssembly is not supported in your browser.");
}