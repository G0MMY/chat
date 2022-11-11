import React from 'react';
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Auth from './Auth';
import Rooms from './Rooms';


export default function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Auth/>}/>
                <Route path="/rooms" element={<Rooms/>}/>
            </Routes>
        </Router>
    );
}