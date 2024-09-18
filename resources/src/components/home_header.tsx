import {HomeOutlined, PartitionOutlined, PieChartOutlined, SettingOutlined} from "@ant-design/icons";
import {Link} from "react-router-dom";

export function HomeHeader() {
    return (
        <div className={"bg-white shadow-lg w-[60px] h-dvh grid justify-center items-center"}>
            <div className={"space-y-8"}>
                <Link className={"grid gap-1 justify-center"} to={"/home/"}>
                    <HomeOutlined className={"justify-center text-xl"}/>
                    <span className={"text-sm"}>首页</span>
                </Link>
                <Link className={"grid gap-1 justify-center"} to={""}>
                    <PieChartOutlined className={"justify-center text-xl"}/>
                    <span className={"text-sm"}>信息</span>
                </Link>
                <Link className={"grid gap-1 justify-center"} to={""}>
                    <PartitionOutlined className={"justify-center text-xl"}/>
                    <span className={"text-sm"}>设备</span>
                </Link>
                <Link className={"grid gap-1 justify-center"} to={""}>
                    <SettingOutlined className={"justify-center text-xl"}/>
                    <span className={"text-sm"}>设置</span>
                </Link>
            </div>
        </div>
    );
}
