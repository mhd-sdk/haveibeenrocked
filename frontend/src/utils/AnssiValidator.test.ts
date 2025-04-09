import { describe, expect, it } from 'vitest';
import { Recommendations, validatePassword } from './AnssiValidator';

describe('validatePassword', () => {
  it('devrait retourner toutes les recommandations pour un mot de passe vide', () => {
    const result = validatePassword('');
    expect(result).toEqual([
      Recommendations.LENGTH,
      Recommendations.UPPERCASE,
      Recommendations.LOWERCASE,
      Recommendations.NUMBERS,
      Recommendations.SPECIAL_CHARACTERS,
    ]);
  });

  it('devrait retourner LENGTH pour un mot de passe trop court', () => {
    const result = validatePassword('Ab1!');
    expect(result).toContain(Recommendations.LENGTH);
  });

  it('devrait retourner UPPERCASE si aucune majuscule', () => {
    const result = validatePassword('abcdefghij1!');
    expect(result).toContain(Recommendations.UPPERCASE);
  });

  it('devrait retourner LOWERCASE si aucune minuscule', () => {
    const result = validatePassword('ABCDEFGHIJ1!');
    expect(result).toContain(Recommendations.LOWERCASE);
  });

  it('devrait retourner NUMBERS si aucun chiffre', () => {
    const result = validatePassword('ABCDEFghij!!');
    expect(result).toContain(Recommendations.NUMBERS);
  });

  it('devrait retourner SPECIAL_CHARACTERS si aucun caractère spécial', () => {
    const result = validatePassword('ABCDEFghij123');
    expect(result).toContain(Recommendations.SPECIAL_CHARACTERS);
  });

  it('devrait retourner un tableau vide pour un mot de passe valide', () => {
    const result = validatePassword('ABCdef123!@#');
    expect(result).toEqual([]);
  });

  it('devrait retourner plusieurs recommandations si plusieurs critères ne sont pas satisfaits', () => {
    const result = validatePassword('abcdef');
    expect(result).toContain(Recommendations.LENGTH);
    expect(result).toContain(Recommendations.UPPERCASE);
    expect(result).toContain(Recommendations.NUMBERS);
    expect(result).toContain(Recommendations.SPECIAL_CHARACTERS);
    expect(result).toHaveLength(4);
  });

  it('devrait correctement identifier un mot de passe qui ne manque que de caractères spéciaux', () => {
    const result = validatePassword('ABCDEFghij123456');
    expect(result).toEqual([Recommendations.SPECIAL_CHARACTERS]);
  });

  it('devrait correctement identifier un mot de passe qui ne manque que de chiffres', () => {
    const result = validatePassword('ABCDEFghijkl!@#');
    expect(result).toEqual([Recommendations.NUMBERS]);
  });
});
