import { StarRating } from './StarRating';

interface PasswordStrengthProps {
  score: number;
}

export const PasswordStrength = ({ score }: PasswordStrengthProps) => {
  return (
    <div className="bg-gray-50 p-4 rounded-lg border dark:bg-gray-800 dark:border-gray-700">
      <h3 className="text-lg font-semibold mb-2">Password Strength</h3>
      <div className="flex items-center justify-between">
        <div>
          <StarRating score={score} />
        </div>
        <div className="text-right">
          <span className="text-2xl font-bold">{score}/5</span>
          <p className="text-sm text-muted-foreground">
            {score === 0 && 'Very Weak'}
            {score === 1 && 'Weak'}
            {score === 2 && 'Fair'}
            {score === 3 && 'Good'}
            {score === 4 && 'Strong'}
            {score === 5 && 'Very Strong'}
          </p>
        </div>
      </div>
    </div>
  );
};
