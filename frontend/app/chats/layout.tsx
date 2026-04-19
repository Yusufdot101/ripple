import Contacts from "@/components/Contacts";
export default function ChatsLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <div className="flex flex-col gap-y-[8px] flex-1 min-h-0">
            <div className="gap-x-[8px] flex-1 flex min-h-0">
                {/* Sidebar */}
                {/* Hide in mobile */}
                <div className="w-[40%] max-[900px]:hidden nlex border-r-1 border-red">
                    <Contacts />
                </div>

                {/* Main content */}
                <div className="flex-1 min-h-0 flex flex-col">{children}</div>
            </div>
        </div>
    );
}
