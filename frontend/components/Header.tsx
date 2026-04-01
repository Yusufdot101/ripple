"use client";
import Link from "next/link";
import { useAuthStore } from "@/store/useAuthStore";

const Header = () => {
    const isLoggedIn = useAuthStore((state) => state.isLoggedIn);

    return (
        <header className="flex justify-between items-center w-full min-w-[300px] py-[12px] px-[24px]">
            <Link href={"/"}>
                <div className="flex items-center gap-[12px] cursor-pointer">
                    {/* TODO: Add logo */}
                    <span
                        title="site name"
                        className="text-text font-semibold max-[619px]:text-[16px]  min-[620px]:text-[24px] hover:text-accent duration-300"
                    >
                        RIPPLE
                    </span>
                </div>
            </Link>

            <Link href={isLoggedIn ? "/logout" : "/login"}>
                {isLoggedIn ? "Logout" : "Login"}
            </Link>
        </header>
    );
};

export default Header;
