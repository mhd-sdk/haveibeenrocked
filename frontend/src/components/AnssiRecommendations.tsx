import { Check, X } from 'lucide-react';

interface AnssiRecommendationsProps {
  meetsANSSI: boolean;
}

export const AnssiRecommendations = ({ meetsANSSI }: AnssiRecommendationsProps) => {
  return (
    <div className="bg-gray-50 p-4 rounded-lg border dark:bg-gray-800 dark:border-gray-700">
      <div className="flex items-start">
        {meetsANSSI ? <Check className="h-5 w-5 text-green-500 mr-2 mt-0.5" /> : <X className="h-5 w-5 text-red-500 mr-2 mt-0.5" />}
        <div>
          <h3 className="text-lg font-semibold">ANSSI Recommendations</h3>
          <p className="text-muted-foreground">
            {meetsANSSI
              ? 'This password meets the ANSSI security recommendations.'
              : 'This password does not meet the ANSSI security recommendations.'}
          </p>
          <div className="mt-2 text-sm">
            <p>ANSSI recommends passwords that:</p>
            <ul className="list-disc pl-5 mt-1 space-y-1">
              <li>Are at least 12 characters long</li>
              <li>Include uppercase and lowercase letters</li>
              <li>Include numbers</li>
              <li>Include special characters</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
};
