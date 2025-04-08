export enum Recommendations {
  LENGTH = 'Le mot de passe doit contenir au moins 12 caractères.',
  UPPERCASE = 'Le mot de passe doit contenir des lettres majuscules.',
  LOWERCASE = 'Le mot de passe doit contenir des lettres minuscules.',
  NUMBERS = 'Le mot de passe doit contenir des chiffres.',
  SPECIAL_CHARACTERS = 'Le mot de passe doit contenir des caractères spéciaux.',
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
