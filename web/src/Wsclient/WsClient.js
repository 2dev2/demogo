
var socket = new WebSocket("ws://localhost:5555/ws");

let connect = cb => {
    console.log("connecting");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = event => {
        console.log(event);
        cb(event.data);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMsg = msg => {
    console.log("sending msg: ", msg);
    socket.send(msg);
};

export { connect, sendMsg };