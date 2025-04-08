import { PasswordStatus } from './PasswordStatus';
import { PasswordStrength } from './PasswordStrength';

interface Props {
  isLeaked: boolean;

  score?: number;
}

export const Result = ({ isLeaked, score = 1 }: Props) => {
  return (
    <section className="w-full ">
      <div className="px-4 md:px-6">
        <div className="mx-auto max-w-3xl space-y-6">
          <PasswordStatus isLeaked={isLeaked} />
          <PasswordStrength score={score} />
        </div>
      </div>
    </section>
  );
};
