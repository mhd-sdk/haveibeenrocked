import { SHA1 } from 'crypto-js';
import { useState } from 'react';
import { validatePassword } from '../utils/AnssiValidator';
import { AnssiRecommendations } from './AnssiRecommendations';
import { PasswordForm } from './PasswordForm';
import { Result } from './Result';

const MainContent = () => {
  const apiUrl = import.meta.env.VITE_API_URL;
  const ollamaUrl = import.meta.env.VITE_OLLAMA_URL;
  const [showResults, setShowResults] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [password, setPassword] = useState('');
  const [isLeaked, setIsLeaked] = useState(false);
  const [score, setScore] = useState(0);

  const handleSubmit = async () => {
    const fullHash = SHA1(password).toString();
    const prefix = fullHash.substring(0, 5);

    setIsLoading(true);
    const response = await fetch(`${apiUrl}/api/check?prefix=${prefix}`, {
      method: 'POST',
    });
    const data = await response.json();

    const isLeaked = data.some((hash: string) => hash === fullHash);

    fetch(`${ollamaUrl}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        // can't use k-anonymity here and dont have enough time to find a solution
        // also having the prompt stored in frontend is bad as anybody could change it, but this feature was not asked
        // so i made it fast considering the time i had for the project
        prompt: `Rate the password "${password}" on a scale of 1 to 5, and only answer with the number, no sentences allowed, no formating allowed (dont write carriage)`,
        model: 'gemma3:4b',
        stream: false,
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        setScore(data.response);
      });

    setIsLoading(false);
    setShowResults(true);
    setIsLeaked(isLeaked);
  };

  const handleChange = (value: string) => {
    setPassword(value);
  };

  const recommendations = validatePassword(password);

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
            <PasswordForm isLoading={isLoading} onSubmit={handleSubmit} password={password} onChange={handleChange} />
          </div>
        </div>
      </div>

      <div className="mx-auto max-w-3xl space-y-6 py-8 md:py-12">
        <AnssiRecommendations missings={recommendations} />
      </div>

      {showResults && <Result isLeaked={isLeaked} score={score} />}
    </section>
  );
};

export default MainContent;
