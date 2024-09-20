import React, {JSX, useEffect, useRef, useState} from "react";
import {DeviceEntity} from "../../assets/ts/model/entity/device_entity.ts";
import {
    DeviceLightControlAPI,
    DeviceRegisterAPI,
    GetDeviceAPI,
    GetNoRegisterDeviceAPI
} from "../../assets/ts/api/device_api.ts";
import {Modal} from "antd";

export function HomeDevice() {
    const [deviceList, setDeviceList] = useState<DeviceEntity[]>([] as DeviceEntity[]);
    const [noRegDeviceList, setNoRegDeviceList] = useState<DeviceEntity[]>([] as DeviceEntity[]);
    const [update, setUpdate] = useState<boolean>(false);

    const wsDevice = useRef<WebSocket | null>(null);
    const wsNoRegDevice = useRef<WebSocket | null>(null);

    // Modal
    const [modalOpen, setModalOpen] = useState<boolean>(false);

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
            wsDevice.current?.send("pong");
        };
        return () => {
            if (wsDevice.current) {
                console.log('WebSocket 连接已销毁');
                wsDevice.current.close();
            }
        }
    }, []);

    useEffect(() => {
        wsNoRegDevice.current = new WebSocket(`ws://${window.location.host}/ws/no-reg-device`);
        wsNoRegDevice.current.onopen = () => {
            console.log('WebSocket 连接已建立');
        };
        wsNoRegDevice.current.onclose = () => {
            console.log('WebSocket 连接已关闭');
        };
        wsNoRegDevice.current.onerror = (error) => {
            console.error('WebSocket 错误: ', error);
        };
        wsNoRegDevice.current.onmessage = (event) => {
            const data = JSON.parse(event.data);
            setNoRegDeviceList(data);
            wsDevice.current?.send("pong");
        };
        return () => {
            if (wsNoRegDevice.current) {
                console.log('WebSocket 连接已销毁');
                wsNoRegDevice.current.close();
            }
        }
    }, []);

    useEffect(() => {
        setTimeout(async () => {
            const getResponse = await GetNoRegisterDeviceAPI({search: "", page: 1, limit: 1000});
            if (getResponse?.output === "Ok") {
                setNoRegDeviceList(getResponse.data!);
            } else {
                console.error(getResponse?.error_message);
            }
        })
    }, [modalOpen, update]);

    async function clickToSwitchLightDevice(deviceName: string, type: boolean) {
        const getResponse = await DeviceLightControlAPI({device: deviceName, value: type});
        if (getResponse?.output === "Ok") {
            console.log(`设备 ${deviceName} 状态已切换为 ${type}`);
            setUpdate(true);
        } else {
            console.error(getResponse?.error_message);
        }
    }

    async function clickToRegisterDevice(deviceName: string) {
        const getResponse = await DeviceRegisterAPI({device_name: deviceName, authorized: true});
        if (getResponse?.output === "Ok") {
            console.log(`设备 ${deviceName} 已注册`);
            setUpdate(true);
        } else {
            console.error(getResponse?.error_message);
        }
    }

    function OnlineDriveList(): JSX.Element[] {
        const list: JSX.Element[] = [];
        if (deviceList) {
            for (let i = 0; i < deviceList.length; i++) {
                if (deviceList[i].login) {
                    // 分析当前状态处于开灯还是关灯(解析 JSON）
                    let buttonTest: string;
                    try {
                        const parse = JSON.parse(deviceList[i].now_value);
                        if (parse.value) {
                            buttonTest = "关灯";
                        } else {
                            buttonTest = "开灯";
                        }
                    } catch {
                        buttonTest = "开灯";
                    }
                    list.push(
                        <div key={`online-${i}`}
                             className={"bg-white rounded-lg shadow-lg p-2 flex justify-between gap-1"}>
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
                                            await clickToSwitchLightDevice(deviceList[i].device_name, buttonTest === "开灯");
                                        }}>
                                        {buttonTest}
                                    </button>
                                </div>
                            </div>
                        </div>
                    );
                }
            }
        }
        return list;
    }

    function OfflineDriveList(): JSX.Element[] {
        const list: JSX.Element[] = [];
        if (deviceList) {
            for (let i = 0; i < deviceList.length; i++) {
                if (!deviceList[i].login) {
                    list.push(
                        <div key={`offline-${i}`}
                             className={"bg-white rounded-lg shadow-lg p-2 flex justify-between gap-1"}>
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
        }
        return list;
    }

    function NoRegisterDeviceList(): React.JSX.Element[] | null {
        if (noRegDeviceList != null) {
            const list: JSX.Element[] = [];
            for (let i = 0; i < noRegDeviceList.length; i++) {
                list.push(
                    <tr className="odd:bg-gray-50" key={`no-reg-${i}`}>
                        <td className="whitespace-nowrap px-4 py-2 font-medium text-gray-900">{noRegDeviceList[i].device_name}</td>
                        <td className="whitespace-nowrap px-4 py-2 text-gray-700">{noRegDeviceList[i].type}</td>
                        <td className="whitespace-nowrap px-4 py-2 text-gray-700 flex gap-3 justify-end">
                            <button onClick={async () => await clickToRegisterDevice(noRegDeviceList[i].device_name)}
                                    className={"px-3 py-1 bg-blue-500 text-white rounded-md transition hover:scale-105"}>
                                注册
                            </button>
                        </td>
                    </tr>
                );
            }
            return list;
        }
        return null;
    }

    const handleOk = () => {
        console.log('Clicked OK');
        setModalOpen(false);
    }

    return (
        <>
            <div className={"grid gap-3"}>
                <div className={"grid justify-end text-[10px]"}>
                    <button onClick={() => setModalOpen(true)}
                            className={"bg-emerald-600 active:bg-emerald-700 px-2 py-1 rounded-md shadow-lg text-white"}>
                        添加设备
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
            <Modal
                open={modalOpen}
                title="未注册设备查询"
                onOk={handleOk}
                onCancel={handleOk}
                style={{top: 5}}
                footer={[
                    <div className={"flex gap-3 justify-end"}>
                        <button className={"px-4 py-1 rounded-lg shadow-lg bg-red-400 active:bg-red-500 text-white"}
                                onClick={handleOk}>
                            关闭
                        </button>
                    </div>
                ]}
            >
                <div className={"grid gap-1"}>
                    <span className={"text-gray-500 text-[10px]"}>Tips: 30秒刷新一次</span>
                    <div className="overflow-x-auto">
                        <table className="min-w-full divide-y-2 divide-gray-200 bg-white text-sm">
                            <thead className="text-left text-[12px]">
                            <tr>
                                <th className="whitespace-nowrap py-0.5 font-medium text-gray-900">设备名字</th>
                                <th className="whitespace-nowrap py-0.5 font-medium text-gray-900">设备类型</th>
                                <th className="whitespace-nowrap py-0.5 font-medium text-gray-900 text-end">操作</th>
                            </tr>
                            </thead>

                            <tbody className="divide-y divide-gray-200 text-[12px]">
                            <NoRegisterDeviceList/>
                            </tbody>
                        </table>
                    </div>
                </div>
            </Modal>
        </>
    )
}
