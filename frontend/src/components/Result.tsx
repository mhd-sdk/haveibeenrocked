import { AnssiRecommendations } from './AnssiRecommendations';
import { PasswordStatus } from './PasswordStatus';
import { PasswordStrength } from './PasswordStrength';

interface Props {
  results: {
    isLeaked: boolean;
    meetsANSSI: boolean;
    score: number;
  };
}

export const Result = ({ results }: Props) => {
  const { isLeaked, meetsANSSI, score } = results;

  return (
    <section className="w-full py-8 md:py-12">
      <div className="px-4 md:px-6">
        <div className="mx-auto max-w-3xl space-y-6">
          <h2 className="text-2xl font-bold text-center mb-6">Password Analysis Results</h2>
          <PasswordStatus isLeaked={isLeaked} />
          <AnssiRecommendations meetsANSSI={meetsANSSI} />
          <PasswordStrength score={score} />
        </div>
      </div>
    </section>
  );
};
