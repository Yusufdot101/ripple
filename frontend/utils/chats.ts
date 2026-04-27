import { api, BASE_CHAT_SERVICE_API_URL } from "./api";
import { UserType } from "./users";

export interface ChatType {
    ID: number;
    Name: string;
}

export const getChatByUserIDs = async (
    userIDs?: number[],
    rolePermissions?: Map<string, string[]>,
    userRoles?: Map<number, string>,
    chatName?: string,
): Promise<ChatType | undefined> => {
    try {
        if (!userIDs && (!rolePermissions || !userRoles)) return;

        if (!rolePermissions) {
            rolePermissions = new Map<string, string[]>();
            rolePermissions.set("admin", ["send message"]);
            rolePermissions.set("member", ["send message"]);
        }

        if (!userRoles) {
            userRoles = new Map<number, string>();
            if (!userIDs) return;
            for (const user of userIDs) {
                userRoles.set(user, "member");
            }
        }

        const res = await api(`${BASE_CHAT_SERVICE_API_URL}/chats`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                rolePermissions: Object.fromEntries(rolePermissions),
                userRoles: Object.fromEntries(userRoles),
                name: chatName,
            }),
        });

        if (!res) return;
        const data = await res.json();
        if (data.error) {
            console.error(data.error);
            return;
        }
        const chat = data.chat;
        return chat;
    } catch (error) {
        console.error(error);
    }
};

export const getChatByID = async (
    chatID: number,
): Promise<ChatType | undefined> => {
    try {
        const res = await api(`${BASE_CHAT_SERVICE_API_URL}/chats/${chatID}`);
        if (!res) return;
        const data = await res.json();
        if (data.error) {
            console.error(data.error);
            return;
        }
        const chat = data.chat;
        return chat;
    } catch (error) {
        console.error(error);
    }
};

export const getChatUsers = async (
    chatID: number,
): Promise<UserType[] | undefined> => {
    try {
        const res = await api(
            `${BASE_CHAT_SERVICE_API_URL}/chats/${chatID}/users`,
        );
        if (!res) return;
        const data = await res.json();
        if (data.error) {
            console.error(data.error);
            return;
        }
        const users = data.chatUsers;
        return users;
    } catch (error) {
        console.error(error);
    }
};

export type ConversationDataType = {
    Chats: ChatType[];
    Groups: ChatType[];
    Contacts: UserType[];
};

export const getConversations = async (
    query: string,
): Promise<ConversationDataType | undefined> => {
    try {
        const res = await api(
            `${BASE_CHAT_SERVICE_API_URL}/conversations?q=${query}`,
        );
        if (!res) return;
        const data = await res.json();
        if (data.error) {
            console.error(data.error);
            return;
        }
        return data;
    } catch (error) {
        console.error(error);
    }
};
