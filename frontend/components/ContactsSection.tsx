import { UserType } from "@/utils/users";
import Contacts from "./Contacts";

type Props = {
    selectedUsers: number[];
    handleUserClick: (user: UserType) => void;
    excludeUsers?: number[];
    users: UserType[];
    isLoading: boolean;
};

const ContactsSection = ({
    selectedUsers,
    handleUserClick,
    users,
    isLoading,
}: Props) => {
    return (
        <div className="flex flex-col">
            <span className="text-foreground/70">Contacts</span>
            <Contacts
                selectedUsers={selectedUsers}
                handleUserClick={handleUserClick}
                isLoading={isLoading}
                users={users}
            />
        </div>
    );
};

export default ContactsSection;
