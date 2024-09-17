import { createRoot } from 'react-dom/client'
import './assets/css/tailwind.css'
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {BaseHome} from "./view/base_home.tsx";

createRoot(document.getElementById('root')!).render(
    <BrowserRouter>
        <Routes>
            <Route path="/home" element={<BaseHome/>} />
        </Routes>
    </BrowserRouter>
)
