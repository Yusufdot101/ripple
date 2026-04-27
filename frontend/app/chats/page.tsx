import Sidebar from "@/components/Sidebar";

const Chats = () => {
    return (
        <div className="flex flex-col gap-y-[8px] flex-1 min-h-0">
            <div className="flex-1 overflow-x-hidden flex flex-col relative">
                <Sidebar />
            </div>

            <div className="max-[900px]:hidden flex justify-center h-full">
                {/* only show on desktop */}
                Welcome to Ripple
            </div>
        </div>
    );
};

export default Chats;
