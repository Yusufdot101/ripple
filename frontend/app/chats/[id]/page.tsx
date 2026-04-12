"use client";
import { useParams } from "next/navigation";

const ChatPage = () => {
    const params = useParams();
    const id = params.id;
    return <div>{id}</div>;
};

export default ChatPage;
