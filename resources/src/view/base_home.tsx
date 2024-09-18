import {Route, Routes} from "react-router-dom";
import {HomeIndex} from "./home/home_index.tsx";
import {HomeHeader} from "../components/home_header.tsx";

export function BaseHome() {
    return (
        <div className={"p-4 min-h-dvh w-full bg-gray-200"}>
            <div className={"pe-[60px] w-full"}>
                <Routes>
                    <Route path="/" element={<HomeIndex/>}/>
                </Routes>
            </div>
            <div className={"fixed top-0 right-0"}>
                <HomeHeader/>
            </div>
        </div>
    );
}
