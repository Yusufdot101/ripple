import Contacts from "@/components/Contacts";
import SearchBar from "@/components/SearchBar";

export default function ChatsLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <div className="flex flex-col gap-y-[8px]">
            <SearchBar />

            <div className="flex">
                {/* Sidebar */}
                <div className="w-[40%] max-[900px]:hidden">
                    <Contacts />
                </div>

                {/* Main content */}
                <div className="flex-1">{children}</div>
            </div>
        </div>
    );
}
