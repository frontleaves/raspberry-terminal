import {HomeOutlined, PartitionOutlined, PieChartOutlined, SettingOutlined} from "@ant-design/icons";
import {Link, useLocation} from "react-router-dom";
import React, {JSX} from "react";

export function HomeHeader() {
    const location = useLocation();

    interface ActiveProps {
        path: string;
        concat: string
        name: string;
        icon: React.ElementType;  // 这里使用 React.ElementType 而不是 JSX.Element
    }

    function Active({path, concat, name, icon: Icon}: ActiveProps): JSX.Element {
        if (location.pathname === concat) {
            return (
                <Link className={"grid gap-1 justify-center"} to={path}>
                    <Icon className={"justify-center text-xl text-blue-500"}/>
                    <span className={"text-sm text-blue-500"}>{name}</span>
                </Link>
            )
        } else {
            return (
                <Link className={"grid gap-1 justify-center"} to={path}>
                    <Icon className={"justify-center text-xl"}/>
                    <span className={"text-sm"}>{name}</span>
                </Link>
            )
        }
    }

    return (
        <div className={"bg-white shadow-lg w-[60px] h-dvh grid justify-center items-center"}>
            <div className={"space-y-8"}>
                <Active name={"首页"} concat={"/home"} path={"/home"} icon={HomeOutlined}/>
                <Active name={"信息"} concat={"/home/info"} path={"/home/info"} icon={PieChartOutlined}/>
                <Active name={"设备"} concat={"/home/device"} path={"/home/device"} icon={PartitionOutlined}/>
                <Active name={"设置"} concat={"/home/setting"} path={"/home/setting"} icon={SettingOutlined}/>
            </div>
        </div>
    );
}
