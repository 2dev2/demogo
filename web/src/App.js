import logo from './logo.svg';
import './App.css';
import Chat from "./chat/Chat";
import { connect, sendMsg } from "./Wsclient/WsClient";
import {useEffect} from "react";
import ChatHistory from "./ChatHistory/ChatHistory";
import {useState} from "react";

function App() {
    //react wil not rerender if it is same array- object reference is same
   // https://stackoverflow.com/questions/56266575/why-is-usestate-not-triggering-re-render
    const [messages,setMessages]=useState([])

    useEffect(() =>{
        connect((msg)=>{
            if(msg) {
                const parsedMsg = JSON.parse(msg)
                messages.push(parsedMsg)
                setMessages([...messages]);
            }
        })
    },[]);

  return (
    <div className="App">
        <Chat></Chat>
        <ChatHistory messages={messages}/>
    </div>
  );
}

export default App;
