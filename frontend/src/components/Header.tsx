import { Moon, ShieldAlert, Sun } from 'lucide-react';

interface Props {
  toggleDarkMode: () => void;
  darkMode: boolean;
}

export const Header = ({ toggleDarkMode, darkMode }: Props) => {
  return (
    <header className="px-4 lg:px-6 h-16 flex items-center border-b">
      <a className="flex items-center justify-center" href="#">
        <ShieldAlert className="h-6 w-6 text-red-500" />
        <span className="ml-2 text-xl font-bold">rockYou</span>
      </a>
      <nav className="ml-auto flex gap-4 sm:gap-6">
        <a className="text-sm font-medium hover:underline underline-offset-4" href="#">
          Home
        </a>
        <a className="text-sm font-medium hover:underline underline-offset-4" href="#">
          About
        </a>
        <a className="text-sm font-medium hover:underline underline-offset-4" href="#">
          API
        </a>
        <a className="text-sm font-medium hover:underline underline-offset-4" href="#">
          Contact
        </a>

        <button onClick={toggleDarkMode} className="text-sm font-medium cursor-pointer">
          {darkMode ? <Sun className="h-5 w-5" /> : <Moon className="h-5 w-5" />}
        </button>
      </nav>
    </header>
  );
};
