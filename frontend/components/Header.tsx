import Link from "next/link";
const Header = () => {
    return (
        <header className="flex justify-between items-center w-full min-w-[300px] border-[1px] border-solid border-[#ffffff] rounded-[8px] py-[12px] px-[24px]">
            <Link href={"/"}>
                <div className="flex items-center gap-[12px] cursor-pointer">
                    {/* TODO: Add logo */}
                    <span className="text-text font-semibold max-[619px]:text-[16px]  min-[620px]:text-[24px] hover:text-accent duration-300">
                        RIBBLE
                    </span>
                </div>
            </Link>

            <Link href={"/login"}>Login</Link>
        </header>
    );
};

export default Header;
