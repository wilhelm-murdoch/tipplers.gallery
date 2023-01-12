import cocktails from './cocktails.json';

export const load = async () => {
  return {
    cocktails: cocktails
  }
}