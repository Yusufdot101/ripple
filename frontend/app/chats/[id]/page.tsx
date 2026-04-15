"use client";
import Message from "@/components/Message";
import MessageInput from "@/components/MessageInput";
import { useAuthStore } from "@/store/useAuthStore";
import { MessageType } from "@/utils/chats";

const ChatPage = () => {
    const userID = useAuthStore((state) => state.userID);

    // export interface MessageType {
    //     ID: number;
    //     ChatID: number;
    //     SenderID: number;
    //     Content: string;
    // }
    const dummyMessages: MessageType[] = [
        {
            ID: 1,
            ChatID: 1,
            SenderID: 1,
            Content: "Yo",
        },
        {
            ID: 2,
            ChatID: 1,
            SenderID: 1,
            Content: "whats up",
        },
        {
            ID: 3,
            ChatID: 1,
            SenderID: userID!,
            Content: "Just Vibin",
        },
        {
            ID: 4,
            ChatID: 1,
            SenderID: userID!,
            Content: "what about you, did you catch the exams",
        },
    ];

    return (
        <div className="flex flex-col h-full gap-y-[8px]">
            <div className="flex justify-center">Username/email</div>

            <div className="flex flex-col gap-y-[8px]">
                {dummyMessages.map((message) => (
                    <Message key={message.ID} message={message} />
                ))}
            </div>

            <MessageInput />
        </div>
    );
};

export default ChatPage;
