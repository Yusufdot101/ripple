"use client";
import { useAuthStore } from "@/store/useAuthStore";
import { logout } from "@/utils/logout";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

function Logout() {
    const router = useRouter();
    const isLoggedIn = useAuthStore((state) => state.isLoggedIn);
    useEffect(() => {
        if (!isLoggedIn) {
            router.push("/");
        }
    }, [isLoggedIn]);

    useEffect(() => {
        logout();
    }, []);
    return <div></div>;
}

export default Logout;
