import { api, BASE_CHAT_SERVICE_API_URL } from "./api";

export type PermissionType = {
    id: number;
    name: string;
};
export const getUserPermissions = async (
    chatID: number,
): Promise<PermissionType[] | undefined> => {
    try {
        // /chats/:chatId/permissions
        const res = await api(
            `${BASE_CHAT_SERVICE_API_URL}/chats/${chatID}/permissions`,
        );
        if (!res) return;
        const data = await res.json();
        return data.permissions;
    } catch (error) {
        console.error(error);
    }
};
