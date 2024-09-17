import {useEffect, useRef, useState} from 'react';

export const HomeIndex = () => {
    const [messages, setMessages] = useState<string[]>([]); // 明确指定 messages 类型为 string[]
    const ws = useRef<WebSocket | null>(null); // 使用 useRef 来保持 WebSocket 实例

    useEffect(() => {
        // 创建 WebSocket 连接
        ws.current = new WebSocket('ws://localhost:8080/ws');

        // 连接打开时
        ws.current.onopen = () => {
            console.log('WebSocket 连接已建立');
        };

        // 处理收到的消息
        ws.current.onmessage = (event) => {
            setMessages((prevMessages) => [...prevMessages, event.data]);
        };

        // 连接关闭时
        ws.current.onclose = () => {
            console.log('WebSocket 连接已关闭');
        };

        // 处理错误
        ws.current.onerror = (error) => {
            console.error('WebSocket 错误: ', error);
        };

        // 清理函数，在组件卸载时关闭 WebSocket 连接
        return () => {
            if (ws.current) {
                ws.current.close();
            }
        };
    }, []);

    const sendMessage = () => {
        if (ws.current && ws.current.readyState === WebSocket.OPEN) {
            ws.current.send('Hello from React!');
        }
    };

    return (
        <div>
            <button onClick={sendMessage}>发送消息</button>
            <div>
                <h3>收到的消息:</h3>
                <ul>
                    {messages.map((msg, index) => (
                        <li key={index}>{msg}</li>
                    ))}
                </ul>
            </div>
        </div>
    );
};
