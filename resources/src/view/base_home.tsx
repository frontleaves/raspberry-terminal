import {Route, Routes} from "react-router-dom";
import {HomeIndex} from "./home/home_index.tsx";

export function BaseHome() {
    return (
        <Routes>
            <Route path="/" element={<HomeIndex/>} />
        </Routes>
    );
}
