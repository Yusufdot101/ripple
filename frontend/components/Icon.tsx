"use client";
import Image from "next/image";
import { useRouter } from "next/navigation";

const Icon = ({
    src,
    href,
    alt,
    background,
    width,
    height = "50px",
}: {
    src: string;
    href: string;
    alt: string;
    background?: string;
    height?: string;
    width: string;
}) => {
    const router = useRouter();
    return (
        <div
            style={{
                background: background,
                height: height,
                width: width,
            }}
            className="rounded-[8px] py-[8px] flex items-center justify-center cursor-pointer"
            role="button"
            tabIndex={0}
            onClick={() => {
                if (href.startsWith("http://") || href.startsWith("https://")) {
                    window.location.href = href;
                } else {
                    router.push(href);
                }
            }}
            onKeyDown={(e) => {
                if (e.key === "Enter" || e.key === " ") {
                    e.preventDefault();
                    if (
                        href.startsWith("http://") ||
                        href.startsWith("https://")
                    ) {
                        window.location.href = href;
                    } else {
                        router.push(href);
                    }
                }
            }}
        >
            <Image src={src} alt={alt} className="h-full" />
        </div>
    );
};

export default Icon;
