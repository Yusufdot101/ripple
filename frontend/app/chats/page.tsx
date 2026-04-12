import Contacts from "@/components/Contacts";

const Chats = () => {
    return (
        <div className="min-[899px]:hidden">
            {/* only show on mobile */}
            <Contacts />
        </div>
    );
};

export default Chats;
