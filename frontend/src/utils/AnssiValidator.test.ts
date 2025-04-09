import { describe, expect, it } from 'vitest';
import { Recommendations, validatePassword } from './AnssiValidator';

describe('validatePassword', () => {
  it('should return all recommendations for an empty password', () => {
    const result = validatePassword('');
    expect(result).toEqual([
      Recommendations.LENGTH,
      Recommendations.UPPERCASE,
      Recommendations.LOWERCASE,
      Recommendations.NUMBERS,
      Recommendations.SPECIAL_CHARACTERS,
    ]);
  });

  it('should return LENGTH for a password that is too short', () => {
    const result = validatePassword('Ab1!');
    expect(result).toContain(Recommendations.LENGTH);
  });

  it('should return UPPERCASE if no uppercase letter', () => {
    const result = validatePassword('abcdefghij1!');
    expect(result).toContain(Recommendations.UPPERCASE);
  });

  it('should return LOWERCASE if no lowercase letter', () => {
    const result = validatePassword('ABCDEFGHIJ1!');
    expect(result).toContain(Recommendations.LOWERCASE);
  });

  it('should return NUMBERS if no digits', () => {
    const result = validatePassword('ABCDEFghij!!');
    expect(result).toContain(Recommendations.NUMBERS);
  });

  it('should return SPECIAL_CHARACTERS if no special characters', () => {
    const result = validatePassword('ABCDEFghij123');
    expect(result).toContain(Recommendations.SPECIAL_CHARACTERS);
  });

  it('should return an empty array for a valid password', () => {
    const result = validatePassword('ABCdef123!@#');
    expect(result).toEqual([]);
  });

  it('should return multiple recommendations if multiple criteria are not met', () => {
    const result = validatePassword('abcdef');
    expect(result).toContain(Recommendations.LENGTH);
    expect(result).toContain(Recommendations.UPPERCASE);
    expect(result).toContain(Recommendations.NUMBERS);
    expect(result).toContain(Recommendations.SPECIAL_CHARACTERS);
    expect(result).toHaveLength(4);
  });

  it('should correctly identify a password that only lacks special characters', () => {
    const result = validatePassword('ABCDEFghij123456');
    expect(result).toEqual([Recommendations.SPECIAL_CHARACTERS]);
  });

  it('should correctly identify a password that only lacks digits', () => {
    const result = validatePassword('ABCDEFghijkl!@#');
    expect(result).toEqual([Recommendations.NUMBERS]);
  });
});
