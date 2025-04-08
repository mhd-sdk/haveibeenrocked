import { Eye, EyeOff, Loader, Search } from 'lucide-react';
import React, { useState } from 'react';
import { Button } from './ui/button';
import { Input } from './ui/input';

interface Props {
  password: string;
  onChange: (password: string) => void;
  onSubmit: () => void;
  isLoading: boolean;
}

export const PasswordForm = ({ password, onChange, onSubmit, isLoading }: Props) => {
  const [showPassword, setShowPassword] = useState(false);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    onSubmit();
  };

  return (
    <form onSubmit={handleSubmit} className="flex w-full max-w-md items-center space-x-2">
      <div className="relative w-full">
        <Input
          name="password"
          type={showPassword ? 'text' : 'password'}
          placeholder="Enter your password"
          className="pr-10 flex-grow"
          value={password}
          onChange={(e) => onChange(e.target.value)}
        />
        <button
          type="button"
          className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-700"
          onClick={() => setShowPassword(!showPassword)}
        >
          {showPassword ? <EyeOff className="h-5 w-5" /> : <Eye className="h-5 w-5" />}
        </button>
      </div>
      <Button type="submit" disabled={password.length === 0}>
        {isLoading && <Loader />}
        {!isLoading && <Search className="h-4 w-4" />}
      </Button>
    </form>
  );
};
