'use client'
import { useEffect } from "react";

const WebsocketTest = () => {
    const socket = new WebSocket("ws://localhost:3001");

    // strictモードがあるのでコンポーネントがマウントされた時にのみ実行
    useEffect(() => {
        socket.addEventListener("open", () => {
            socket.send("Hello Server!");
        });
        
        socket.addEventListener("message", (event) => {
            console.log("Message from server ", event.data);
        });
    }, []);

    return (
        <div className="test">websocketテスト</div>
    );
};

export default WebsocketTest;