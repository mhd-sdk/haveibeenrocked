export enum Recommendations {
  LENGTH = 'Password must contain at least 12 characters.',
  UPPERCASE = 'Password must contain uppercase letters.',
  LOWERCASE = 'Password must contain lowercase letters.',
  NUMBERS = 'Password must contain digits.',
  SPECIAL_CHARACTERS = 'Password must contain special characters.',
}

export const validatePassword = (password: string): Recommendations[] => {
  const recommendations: Recommendations[] = [];

  if (password.length < 12) {
    recommendations.push(Recommendations.LENGTH);
  }
  if (!/[A-Z]/.test(password)) {
    recommendations.push(Recommendations.UPPERCASE);
  }
  if (!/[a-z]/.test(password)) {
    recommendations.push(Recommendations.LOWERCASE);
  }
  if (!/[0-9]/.test(password)) {
    recommendations.push(Recommendations.NUMBERS);
  }
  if (!/[!@#$%^&*(),.?":{}|<>]/.test(password)) {
    recommendations.push(Recommendations.SPECIAL_CHARACTERS);
  }

  return recommendations;
};
