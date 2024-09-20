import {JSX, useEffect, useRef, useState} from "react";
import {DeviceEntity} from "../../assets/ts/model/entity/device_entity.ts";
import {DeviceLightControlAPI, GetDeviceAPI} from "../../assets/ts/api/device_api.ts";

export function HomeDevice() {
    const [deviceList, setDeviceList] = useState<DeviceEntity[]>([] as DeviceEntity[]);
    const [update, setUpdate] = useState<boolean>(false);

    const wsDevice = useRef<WebSocket | null>(null);

    useEffect(() => {
        setUpdate(false);
        setTimeout(async () => {
            const getResponse = await GetDeviceAPI({search: "", page: 1, limit: 1000});
            if (getResponse?.output === "Ok") {
                setDeviceList(getResponse.data!);
            } else {
                console.error(getResponse?.error_message);
            }
        })
    }, [update]);

    useEffect(() => {
        wsDevice.current = new WebSocket(`ws://${window.location.host}/ws/device`);
        wsDevice.current.onopen = () => {
            console.log('WebSocket 连接已建立');
        };
        wsDevice.current.onclose = () => {
            console.log('WebSocket 连接已关闭');
        };
        wsDevice.current.onerror = (error) => {
            console.error('WebSocket 错误: ', error);
        };
        wsDevice.current.onmessage = (event) => {
            const data = JSON.parse(event.data);
            setDeviceList(data);
        };
    }, []);

    async function clickToSwitchLightDevice(deviceName: string, type: boolean) {
        const getResponse = await DeviceLightControlAPI({device: deviceName, value: type});
        if (getResponse?.output === "Ok") {
            console.log(`设备 ${deviceName} 状态已切换为 ${type}`);
            setUpdate(true);
        } else {
            console.error(getResponse?.error_message);
        }
    }

    function OnlineDriveList(): JSX.Element[] {
        const list: JSX.Element[] = [];
        for (let i = 0; i < deviceList.length; i++) {
            if (deviceList[i].login) {
                // 分析当前状态处于开灯还是关灯(解析 JSON）
                let buttonTest: string;
                const parse = JSON.parse(deviceList[i].now_value);
                if (parse.value) {
                    buttonTest = "关灯";
                } else {
                    buttonTest = "开灯";
                }
                list.push(
                    <div key={`online-${i}`} className={"bg-white rounded-lg shadow-lg p-2 flex justify-between gap-1"}>
                        <div className={""}>
                        </div>
                        <div className={"grid gap-1"}>
                            <div className={"text-end"}>
                                <div className={"text-sm font-bold"}>{deviceList[i].device_host}</div>
                                <div className={"text-[10px]"}>{deviceList[i].type}</div>
                            </div>
                            <div className={"grid justify-end"}>
                                <button
                                    className={"bg-blue-400 text-[10px] text-white rounded px-4 py-0.5 active:bg-blue-500"}
                                    onClick={async () => {
                                        await clickToSwitchLightDevice(deviceList[i].device_name, !parse.value)
                                    }}>
                                    {buttonTest}
                                </button>
                            </div>
                        </div>
                    </div>
                );
            }
        }
        return list;
    }

    function OfflineDriveList(): JSX.Element[] {
        const list: JSX.Element[] = [];
        for (let i = 0; i < deviceList.length; i++) {
            if (!deviceList[i].login) {
                list.push(
                    <div key={`offline-${i}`} className={"bg-white rounded-lg shadow-lg p-2 flex justify-between gap-1"}>
                        <div className={""}>
                        </div>
                        <div className={"grid gap-1"}>
                            <div className={"text-end"}>
                                <div className={"text-sm font-bold"}>{deviceList[i].device_host}</div>
                                <div className={"text-[10px]"}>{deviceList[i].type}</div>
                            </div>
                        </div>
                    </div>
                );
            }
        }
        return list;
    }

    return (
        <div className={"grid gap-3"}>
            <div className={"grid justify-end text-[10px]"}>
                <button
                    className={"bg-emerald-600 active:bg-emerald-700 px-2 py-1 rounded-md shadow-lg text-white"}>添加设备
                </button>
            </div>
            <div className={"text-md font-bold"}>在线设备</div>
            <div className={"grid grid-cols-2 gap-3"}>
                <OnlineDriveList/>
            </div>
            <div className={"text-md font-bold"}>离线设备</div>
            <div className={"grid grid-cols-2 gap-3"}>
                <OfflineDriveList/>
            </div>
        </div>
    )
}
