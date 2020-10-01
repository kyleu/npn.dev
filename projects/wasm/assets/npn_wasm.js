const go = new Go();
WebAssembly.instantiateStreaming(fetch('npn.wasm'), go.importObject).then((result) => {
  go.run(result.instance);
  npn_register(function(msgStr) {
    const msg = JSON.parse(msgStr);
    socket.recv(msg);
  });
});
