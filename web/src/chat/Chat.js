import React, {useEffect, useState} from 'react';
import { connect, sendMsg } from "../Wsclient/WsClient";
function Chat() {
    const [message, setMessage] = useState('');
    useEffect(() =>{
        connect((msg)=>{})
    },[]);
    const chatSubmit = (event)=>{
        console.log(message)
        sendMsg(message)
        event.preventDefault()
        event.stopPropagation()
    }
    const setMessageEvent = (event)=>{
        setMessage(event.target.value)
        event.preventDefault()
        event.stopPropagation()
    }
    return (
        <form >
            <input onChange={setMessageEvent}>
            </input>
            <button onClick={chatSubmit}>Send </button>
        </form>
    );
}

export default Chat;