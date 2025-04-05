export default function Footer() {
  return (
    <footer className="flex flex-col gap-2 sm:flex-row py-6 w-full shrink-0 items-center px-4 md:px-6 border-t">
      <p className="text-xs text-slate-500 dark:text-slate-400">
        © {new Date().getFullYear()} rockYou. Tous droits réservés.
      </p>
      <nav className="sm:ml-auto flex gap-4 sm:gap-6">
        <a className="text-xs hover:underline underline-offset-4" href="#">
          Conditions d'utilisation
        </a>
        <a className="text-xs hover:underline underline-offset-4" href="#">
          Confidentialité
        </a>
        <a className="text-xs hover:underline underline-offset-4" href="#">
          Contact
        </a>
      </nav>
    </footer>
  );
} 