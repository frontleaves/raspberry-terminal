import {useEffect, useRef, useState} from 'react';
import {useNavigate} from "react-router-dom";
import {WsSystemInfoDTO} from "../../assets/ts/model/dto/ws_system_info_dto.ts";

export const HomeIndex = () => {
    const navigate = useNavigate();

    const [hasConnect, setHasConnect] = useState<boolean>(true);
    const [nowTime, setNowTime] = useState<string>(new Date().toLocaleTimeString());
    const [refresh, setRefresh] = useState<number>();
    const [webSocket, setWebSocket] = useState<number>();
    const [systemInfo, setSystemInfo] = useState<WsSystemInfoDTO>({
        cpu: 0,
        ram: 0,
        disk: 0,
    } as WsSystemInfoDTO);
    const [demoData, setDemoData] = useState<string>("未知");

    const wsPing = useRef<WebSocket | null>(null);
    const wsSystemInfo = useRef<WebSocket | null>(null);

    useEffect(() => {
        if (webSocket) {
            return
        }
        setWebSocket(setTimeout(() => {
            wsPing.current = new WebSocket(`ws://${window.location.host}/ws/ping`);
            wsPing.current.onopen = () => {
                setHasConnect(true);
                console.log('WebSocket 连接已建立');
            };
            wsPing.current.onclose = () => {
                setHasConnect(false);
                setWebSocket(undefined);
                console.log('WebSocket 连接已关闭');
            };
            wsPing.current.onerror = (error) => {
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
        wsSystemInfo.current = new WebSocket(`ws://${window.location.host}/ws/system`);
        wsSystemInfo.current.onopen = () => {
            console.log('WebSocket 连接已建立');
        };
        wsSystemInfo.current.onclose = () => {
            console.log('WebSocket 连接已关闭');
        };
        wsSystemInfo.current.onerror = (error) => {
            console.error('WebSocket 错误: ', error);
        };
        wsSystemInfo.current.onmessage = (event) => {
            const data = JSON.parse(event.data);
            // 保留两位小数
            data.cpu = parseFloat(data.cpu).toFixed(1);
            data.ram = parseFloat(data.ram).toFixed(1);
            data.disk = parseFloat(data.disk).toFixed(1);
            setSystemInfo(data);
        };
    }, []);

    useEffect(() => {
        wsSystemInfo.current = new WebSocket(`ws://${window.location.host}/ws/mqtt`);
        wsSystemInfo.current.onopen = () => {
            console.log('WebSocket 连接已建立');
        };
        wsSystemInfo.current.onclose = () => {
            console.log('WebSocket 连接已关闭');
        };
        wsSystemInfo.current.onerror = (error) => {
            console.error('WebSocket 错误: ', error);
        };
        wsSystemInfo.current.onmessage = (event) => {
            setDemoData(event.data);
        };
    })

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
            <div className={"text-xl font-extrabold"}>{nowTime}</div>
        );
    }

    return (
        <div className={"grid gap-3 grid-cols-12"}>
            <div className={"bg-white shadow-lg p-3 rounded-lg flex justify-between items-center col-span-12"}>
                <div>嵌入式终端</div>
                <div><HasConnectText/></div>
            </div>
            <div className={"grid items-center gap-3 col-span-6"}>
                <div className={"bg-white shadow-lg p-3 rounded-lg grid items-center justify-center w-full"}>
                    <div><NowTime/></div>
                </div>
                <div className={"bg-white shadow-lg p-3 rounded-lg grid items-center justify-center w-full"}>
                    <div>{demoData}</div>
                </div>
            </div>
            <div className={"bg-white shadow-lg p-3 rounded-lg grid items-center col-span-6 space-y-1"}>
                <div className={"grid"}>
                    <div className={"text-[10px]/4 font-thin"}>处理器占用率</div>
                    <span id="ProgressLabel" className="sr-only">Loading</span>
                    <div
                        role="progressbar"
                        aria-labelledby="ProgressLabel"
                        aria-valuenow={50}
                        className="rounded-full bg-gray-200 w-full"
                    >
                        <span className="block h-4 rounded-full bg-indigo-600 text-center text-[10px]/4"
                              style={{width: `${systemInfo.cpu}%`}}>
                          <span className="font-bold text-white">{systemInfo.cpu}%</span>
                        </span>
                    </div>
                </div>
                <div className={"grid"}>
                    <div className={"text-[10px]/4 font-thin"}>内存占用率</div>
                    <span id="ProgressLabel" className="sr-only">Loading</span>
                    <div
                        role="progressbar"
                        aria-labelledby="ProgressLabel"
                        aria-valuenow={50}
                        className="rounded-full bg-gray-200 w-full"
                    >
                        <span className="block h-4 rounded-full bg-indigo-600 text-center text-[10px]/4"
                              style={{width: `${systemInfo.ram}%`}}>
                          <span className="font-bold text-white">{systemInfo.ram}%</span>
                        </span>
                    </div>
                </div>
                <div className={"grid"}>
                    <div className={"text-[10px]/4 font-thin"}>存储占用率</div>
                    <span id="ProgressLabel" className="sr-only">Loading</span>
                    <div
                        role="progressbar"
                        aria-labelledby="ProgressLabel"
                        aria-valuenow={50}
                        className="rounded-full bg-gray-200 w-full"
                    >
                        <span className="block h-4 rounded-full bg-indigo-600 text-center text-[10px]/4"
                              style={{width: `${systemInfo.disk}%`}}>
                          <span className="font-bold text-white">{systemInfo.disk}%</span>
                        </span>
                    </div>
                </div>
            </div>
        </div>
    );
};
