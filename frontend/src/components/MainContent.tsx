import { useState } from 'react';
import { PasswordForm } from './PasswordForm';
import { Result } from './Result';

interface Results {
  isLeaked: boolean;
  meetsANSSI: boolean;
  score: number;
}

export default function MainContent() {
  const [showResults, setShowResults] = useState(false);
  const [results, setResults] = useState<Results>({
    isLeaked: true,
    meetsANSSI: false,
    score: 0,
  });

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const password = formData.get('password') as string;

    if (password) {
      setShowResults(true);
      setResults({
        isLeaked: true,
        meetsANSSI: false,
        score: 0,
      });
    }
  };

  return (
    <section className="w-full py-12 md:py-24 lg:py-32">
      <div className="px-4 md:px-6">
        <div className="flex flex-col items-center space-y-4 text-center">
          <div className="space-y-2">
            <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl lg:text-6xl/none">
              Check if your password has been compromised
            </h1>
            <p className="mx-auto max-w-[700px] text-muted-foreground md:text-xl">
              Search across multiple data breaches to see if your password has been compromised.
            </p>
          </div>
          <div className="w-full max-w-md space-y-2">
            <PasswordForm onSubmit={handleSubmit} />
          </div>
        </div>
      </div>
      {showResults && <Result results={results} />}
    </section>
  );
}
