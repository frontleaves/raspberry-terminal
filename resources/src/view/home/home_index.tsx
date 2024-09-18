import {useEffect, useRef, useState} from 'react';
import {useNavigate} from "react-router-dom";

export const HomeIndex = () => {
    const navigate = useNavigate();

    const [hasConnect, setHasConnect] = useState<boolean>(true);
    const [nowTime, setNowTime] = useState<string>(new Date().toLocaleTimeString());
    const [refresh, setRefresh] = useState<number>();
    const [webSocket, setWebSocket] = useState<number>()
    const ws = useRef<WebSocket | null>(null);

    useEffect(() => {
        if (webSocket) {
            return
        }
        setWebSocket(setTimeout(() => {
            ws.current = new WebSocket('ws://localhost:8080/ws/ping');
            ws.current.onopen = () => {
                setHasConnect(true);
                console.log('WebSocket 连接已建立');
            };
            ws.current.onclose = () => {
                setHasConnect(false);
                setWebSocket(undefined);
                console.log('WebSocket 连接已关闭');
            };
            ws.current.onerror = (error) => {
                setHasConnect(false);
                setWebSocket(undefined);
                console.error('WebSocket 错误: ', error);
            };
        }, 1000));

        setInterval(() => {
            setNowTime(new Date().toLocaleTimeString());
        }, 100);
    }, [refresh, webSocket]);

    useEffect(() => {
        if (!hasConnect) {
            // 使用 react-router-dom 进行页面导航
            setRefresh(setInterval(() => {
                navigate(location.pathname, {replace: true});
            }, 2000));
        } else {
            clearInterval(refresh);
        }
    }, [hasConnect]);

    function HasConnectText() {
        if (hasConnect) {
            return (
                <div className={"flex gap-1"}>
                    <span className={"text-emerald-500 font-extrabold"}>•</span>
                    <span className={"text-green-700"}>已连接</span>
                </div>
            );
        } else {
            return (
                <div className={"flex gap-1"}>
                    <span className={"text-red-500 font-extrabold"}>•</span>
                    <span className={"text-red-700"}>未连接</span>
                </div>
            );
        }
    }

    function NowTime() {
        return (
            <div>{nowTime}</div>
        );
    }

    return (
        <div className={"grid gap-3 grid-cols-12"}>
            <div className={"bg-white shadow-lg p-3 rounded-lg flex justify-between items-center col-span-12"}>
                <div>嵌入式机器</div>
                <div><HasConnectText/></div>
            </div>
            <div className={"bg-white shadow-lg p-3 rounded-lg flex items-center col-span-6"}>
                <div><NowTime/></div>
            </div>
            <div className={"bg-white shadow-lg p-3 rounded-lg flex justify-between items-center col-span-6"}>
                <div>213</div>
                <div><HasConnectText/></div>
            </div>
        </div>
    );
};
