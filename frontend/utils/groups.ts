import { api, BASE_CHAT_SERVICE_API_URL } from "./api";

export const addUsersToGroup = async (chatID: number, userIDs: number[]) => {
    try {
        const res = await api(
            `${BASE_CHAT_SERVICE_API_URL}/chats/${chatID}/addToGroup`,
            {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ userIDs: userIDs }),
            },
        );
        console.log("res: ", res);
        if (!res) return;
        const data = await res.json();
        console.log("data: ", data);
    } catch (error) {
        console.error(error);
    }
};
