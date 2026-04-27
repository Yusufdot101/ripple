import UserCard from "@/components/UserCard";
import UserCardSkeleton from "@/components/UserCardSkeleton";
import { UserType } from "@/utils/users";

interface Props {
    selectedUsers: number[];
    handleUserClick: (user: UserType) => void;
    excludeUsers?: number[];
    users: UserType[];
    isLoading: boolean;
}

const Contacts = ({
    selectedUsers,
    handleUserClick,
    excludeUsers,
    users,
    isLoading,
}: Props) => {
    const visibleUsers = (users ?? []).filter(
        (elem) => !excludeUsers?.includes(elem.id),
    );

    return (
        <div className="flex-1 flex flex-col gap-y-[8px] rounded-[4px] gap-y-[8px]">
            <div
                className={`${!isLoading ? "hidden" : ""} flex flex-col transition-all duration-300`}
            >
                {[...Array(4).keys()].map((el) => (
                    <UserCardSkeleton key={el} index={el} />
                ))}
            </div>

            {visibleUsers.length === 0 && !isLoading ? (
                <p className="w-full text-center">No users</p>
            ) : null}

            <div
                className={`${visibleUsers?.length === 0 ? "opacity-0 blur-sm" : ""} flex flex-col transition-all duration-300`}
            >
                {visibleUsers.map((user) => (
                    <UserCard
                        activeUsers={selectedUsers}
                        key={user.id}
                        user={user}
                        handleClick={handleUserClick}
                    />
                ))}
            </div>
        </div>
    );
};

export default Contacts;
