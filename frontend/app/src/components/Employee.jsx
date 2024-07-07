import React from "react";


export default function Employee(props) {
    function onClick() {
        console.log("test")
    }
    return <div>
        Emoloyees
        <button onClick={onClick}>Test</button>
    </div>
}