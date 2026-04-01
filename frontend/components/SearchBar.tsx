import Image from "next/image";
import searchIcon from "../assets/searchIcon.svg";

const SearchBar = () => {
    return (
        <div className="border-[1px] border-solid border-[#ffffff] rounded-[8px] py-[12px] px-[12px] flex gap-x-[4px] h-[50px] cursor-pointer opacity-80">
            <Image
                role="button"
                tabIndex={0}
                aria-label="search"
                src={searchIcon}
                alt="search icon"
                className="h-full w-[40px] rounded-[8px]"
            />

            <input
                type="text"
                placeholder="Search or start new chat"
                className="border-none outline-none h-full w-full"
            />
        </div>
    );
};

export default SearchBar;
