let socket: WebSocket;

if (typeof window !== 'undefined') {
  socket = new window.WebSocket('ws://localhost:8080/ws');
}

export type SocketData = {
  message: string;
};

const connect = (cb: (msg: SocketData) => void): void => {
  console.log('[socket] Attempting connection...');

  socket.onopen = () => {
    console.log('[socket] Successfully connected');
  };

  socket.onmessage = (msg: MessageEvent<string>) => {
    console.log('[socket] Incoming message: ', msg);
    const parsedMessage = JSON.parse(msg.data) as SocketData;
    console.log('[socket] Parsed data:', parsedMessage);
    cb(parsedMessage);
  };

  socket.onclose = (event) => {
    console.log('[socket] Socket Closed Connection: ', event);
  };

  socket.onerror = (error) => {
    console.log('[socket] Socket Error: ', error);
  };
};

const sendMessage = (msg: SocketData) => {
  console.log('[socket] Sending message: ', msg);
  socket.send(JSON.stringify(msg));
};

export { connect, sendMessage };
