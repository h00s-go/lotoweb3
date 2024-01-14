import { PUBLIC_API_URL } from '$env/static/public'

export async function load({ fetch, params }) {
  const response = await fetch(PUBLIC_API_URL + '/lotteries/pick-one?numbers=7&max=35');
  const numbers = await response.json();

  return {
    numbers: numbers.join(', '),
  };
}