import jsQR from 'jsqr';

self.onmessage = (event) => {
  const { data, width, height } = event.data;
  const code = jsQR(data, width, height);
  if (code) {
    self.postMessage(code.data);
  } else {
    self.postMessage(null);
  }
};