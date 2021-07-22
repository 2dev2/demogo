


const ChatHistory = (props)=>{
    const list=props.messages.map((msg, index)=>{
        return <p key={index}>{msg.id ? msg.id+': ':''} {msg.body}</p>
    })
    return(
        <div>
            <div> chat history </div>
            <div style={{border: 'solid 1px'}}>
                {list}
            </div>
        </div>
    )
}
ChatHistory.defaultProps = {
    messages:[]
}

export default ChatHistory;