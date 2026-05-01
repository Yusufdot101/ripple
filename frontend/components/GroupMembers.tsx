"use client";
import Contacts from "@/components/Contacts";
import { UserType } from "@/utils/users";
import { useCallback, useEffect, useState } from "react";
import BackArrowButton from "./BackArrowButton";
import SearchBar from "./SearchBar";
import { getChatUsers } from "@/utils/chats";
import { useAuthStore } from "@/store/useAuthStore";
import { removeUserFromGroup } from "@/utils/groups";
import { useRouter } from "next/navigation";

interface Props {
    handleClose: () => void;
    groupMembersIsOpen: boolean;
    chatID: number;
    hasPermission: (permissionName: string) => boolean;
}

const GroupMembers = ({
    handleClose,
    groupMembersIsOpen,
    chatID,
    hasPermission,
}: Props) => {
    const [isLoading, setIsLoading] = useState(false);
    const [users, setUsers] = useState<UserType[]>([]);
    const [clickedUser, setClickedUser] = useState<UserType>();
    const [menuIsOpen, setMenuIsOpen] = useState(false);

    const searchUsers = useCallback(
        async (value: string = "") => {
            setIsLoading(true);
            try {
                let users = await getChatUsers(chatID);
                users = users?.filter(
                    (user) =>
                        user.name.includes(value) || user.email.includes(value),
                );
                setUsers(users ?? []);
            } finally {
                setIsLoading(false);
            }
        },
        [chatID],
    );

    const handleRightClick = (user: UserType) => {
        setClickedUser(user);
        setMenuIsOpen(true);
    };

    useEffect(() => {
        (() => searchUsers())();
    }, [searchUsers]);

    const handleRemove = (userID: number) => {
        removeUserFromGroup(chatID, userID);
        setMenuIsOpen(false);
        handleClose();
    };

    const loggedInUserID = useAuthStore((state) => state.userID);
    const router = useRouter();

    return (
        <div
            className={`${groupMembersIsOpen ? "translate-x-0" : "translate-x-full"} transition-transform absolute w-full bg-background top-1/2 translate-y-1/2 duration-300 flex-1 flex overflow-x-hidden z-10 flex-col gap-y-[8px]`}
            onClick={() => setMenuIsOpen(false)}
            onKeyDown={(e) => {
                if (e.key !== "Escape") return;
                setMenuIsOpen(false);
            }}
        >
            <div className="h-full transition-transform duration-300 ease-in-out flex flex-1 flex-col gap-y-[8px]">
                <div className="flex w-full h-[32px] gap-x-[8px] items-center">
                    <BackArrowButton
                        handleClick={handleClose}
                        text="Group members"
                    />
                </div>

                <div className="flex flex-col gap-y-[8px] relative flex-1">
                    <SearchBar
                        placeholder="Search group members"
                        handleEnter={searchUsers}
                    />

                    <div className="h-full max-h-[200px] overflow-y-scroll">
                        <Contacts
                            isLoading={isLoading}
                            users={users}
                            handleUserClick={() => {}}
                            handleUserRightClick={(user: UserType) => {
                                handleRightClick(user);
                            }}
                            selectedUsers={[]}
                            excludeUsers={
                                loggedInUserID ? [loggedInUserID] : []
                            }
                        />
                    </div>
                </div>
            </div>

            <div
                className={`${menuIsOpen ? "max-h-96 p-[4px]" : "max-h-0 invisible p-0"} z-1 duration-300 flex justify-center items-center absolute overflow-hidden h-full w-full bg-background/80`}
            >
                <div className="bg-background w-80 border-1 border-foreground rounded-[4px] flex flex-col justify-center">
                    {hasPermission("remove users from group") && (
                        <button
                            onClick={(e) => {
                                e.stopPropagation();
                                if (!clickedUser) return;
                                if (
                                    !confirm(
                                        `are you sure you want to remove ${clickedUser.name} from this group`,
                                    )
                                )
                                    return;
                                handleRemove(clickedUser.id);
                            }}
                            onKeyDown={(e) => {
                                e.stopPropagation();
                                if (e.key !== "Enter") return;
                                if (!clickedUser) return;
                                if (
                                    !confirm(
                                        `are you sure you want to remove ${clickedUser.name} from this group`,
                                    )
                                )
                                    return;
                                handleRemove(clickedUser.id);
                                router.push("/chats");
                            }}
                            className="cursor-pointer hover:bg-foreground/20 active:bg-background duration-300 p-[4px]"
                        >
                            Remove {clickedUser?.name}
                        </button>
                    )}
                </div>
            </div>

            <button
                onClick={() => {
                    if (!loggedInUserID) return;
                    if (!confirm("are you sure you want to leave this group"))
                        return;
                    handleRemove(loggedInUserID);
                }}
                className="w-full bg-red-500 p-[4px] rounded-[4px] cursor-pointer hover:bg-red-600 active:bg-red-500 duration-300"
            >
                Exit group
            </button>
        </div>
    );
};

export default GroupMembers;
