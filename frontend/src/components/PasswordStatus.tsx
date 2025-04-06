import { Shield, ShieldAlert } from 'lucide-react';

interface PasswordStatusProps {
  isLeaked: boolean;
}

const PasswordStatus = ({ isLeaked }: PasswordStatusProps) => {
  return (
    <div
      className={`flex items-start p-4 rounded-lg border ${isLeaked ? 'bg-red-50 border-red-200 dark:bg-red-900 dark:border-red-700' : 'bg-green-50 border-green-200 dark:bg-green-900 dark:border-green-700'}`}
    >
      <div className="flex items-start">
        {isLeaked ? <ShieldAlert className="h-5 w-5 mr-2 mt-0.5" /> : <Shield className="h-5 w-5 mr-2 mt-0.5" />}
        <div>
          <h3 className="text-lg font-semibold">{isLeaked ? 'Password Compromised' : 'Password Not Found in Breaches'}</h3>
          <p>
            {isLeaked
              ? 'This password has been found in data breaches.'
              : "Good news! This password doesn't appear in our database of compromised passwords."}
          </p>
        </div>
      </div>
    </div>
  );
};

export default PasswordStatus;
