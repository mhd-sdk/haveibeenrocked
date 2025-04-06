import React, { useState } from 'react';
import { Button } from './ui/button';
import { Eye, EyeOff, Search } from 'lucide-react';
import { Input } from './ui/input';

interface Props {
  onSubmit: React.FormEventHandler<HTMLFormElement>;
}

export const PasswordForm = ({ onSubmit }: Props) => {
  const [showPassword, setShowPassword] = useState(false);

  return (
    <form onSubmit={onSubmit} className="flex w-full max-w-md items-center space-x-2">
      <div className="relative w-full">
        <Input name="password" type={showPassword ? 'text' : 'password'} placeholder="Enter your password" className="pr-10 flex-grow " />
        <button
          type="button"
          className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-700"
          onClick={() => setShowPassword(!showPassword)}
        >
          {showPassword ? <EyeOff className="h-5 w-5" /> : <Eye className="h-5 w-5" />}
        </button>
      </div>
      <Button type="submit" className="bg-slate-50 hover:bg-slate-300">
        <Search className="mr-2 h-4 w-4" />
        Search
      </Button>
    </form>
  );
};
