"use client";
import { initAuth } from "@/utils/initAuth";
import { useEffect } from "react";
import githubLogo from "./../assets/github.svg";
import linkedInLogo from "./../assets/linkedIn.svg";
import Icon from "./Icon";

const Footer = () => {
    useEffect(() => {
        initAuth();
    }, []);

    return (
        <div className="bg-primary border-[1px] border-solid border-[#ffffff] rounded-[8px] py-[12px] w-full min-w-[300px] flex items-center justify-center gap-x-[12px] text-text gap-[8px] px-[8px] h-fit">
            <div className="flex gap-[8px] h-[80px]">
                <Icon
                    src={githubLogo}
                    href={"https://github.com/Yusufdot101"}
                    alt={"GitHub logo"}
                    height="100%"
                    width="70px"
                />
                <Icon
                    src={linkedInLogo}
                    href={
                        "https://www.linkedin.com/in/yusuf-mohamed-0a5605366/"
                    }
                    alt={"LinkedIn logo"}
                    height="100%"
                    width="70px"
                />
            </div>
            <div className="flex flex-col text-center">
                <p className="text-[20px] max-[619px]:text-[12px]">
                    Email: yusuf.mohamed.work@gmail.com
                </p>
                <p className="text-[20px] max-[619px]:text-[12px]">
                    COPYRIGHT © 2025 by Yusuf Mohamed
                </p>
            </div>
        </div>
    );
};

export default Footer;
