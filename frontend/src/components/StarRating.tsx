import { Star } from 'lucide-react';

interface Props {
  score: number;
}

export const StarRating = ({ score }: Props) => {
  return (
    <div className="flex">
      {[1, 2, 3, 4, 5].map((star) => (
        <Star key={star} className={`h-6 w-6 ${star <= score ? 'text-yellow-400 fill-yellow-400' : 'text-gray-300'}`} />
      ))}
    </div>
  );
};
